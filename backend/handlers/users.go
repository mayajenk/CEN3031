package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/mayajenk/CEN3031/models"
	"github.com/wader/gormstore/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func sendError(message string, status int, w http.ResponseWriter) {
	res := make(map[string]any)
	res["message"] = message
	res["status"] = status
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func GetAllUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var users []models.User
		db.Model(&models.User{}).Preload("Subjects").Preload("Connections").Preload("Reviews").Find(&users)

		json.NewEncoder(w).Encode(users)
	}
}

func GetUserFromSession(store *gormstore.Store, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		session, err := store.Get(r, "session")
		if err != nil {
			sendError("Error retrieving user", http.StatusUnauthorized, w)
			return
		}

		userID := session.Values["userID"]
		var user models.User

		err = db.Model(&models.User{}).Preload("Subjects").Preload("Connections").Preload("Reviews").First(&user, userID).Error
		if err != nil {
			sendError("Error retrieving user", http.StatusUnauthorized, w)
			return
		}

		if user.IsTutor {
			var tutor models.Tutor
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &tutor)

			if err == nil {
				json.NewEncoder(w).Encode(tutor)
			}
		} else {
			var student models.Student
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &student)

			if err == nil {
				json.NewEncoder(w).Encode(student)
			}
		}
	}
}

func GetUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get User Endpoint Hit")
		w.Header().Set("Content-Type", "application/json")

		userID := mux.Vars(r)["id"]

		var user models.User

		err := db.Model(&models.User{}).Preload("Subjects").Preload("Connections").Preload("Reviews").First(&user, userID).Error
		if err != nil {
			sendError("Error retrieving user", http.StatusUnauthorized, w)
		} else {
			json.NewEncoder(w).Encode(user)
		}
	}

}

func NewUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Println("New User Endpoint Hit")

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		var user models.User
		err := decoder.Decode(&user)
		if err != nil {
			sendError("Bad request format", http.StatusBadRequest, w)
			return
		}

		// Checking if a user is unique in the database
		var existingUser models.User
		result := db.Where("username = ?", user.Username).First(&existingUser)
		if result.Error == nil {
			sendError("Username already exists", http.StatusConflict, w)
			return
		} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			panic(result.Error)
		}

		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			panic("Failed to hash password")
		}
		user.Password = string(password)

		db.Create(&user)

		if user.IsTutor {
			var tutor models.Tutor
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &tutor)
			json.NewEncoder(w).Encode(tutor)
		} else {
			var student models.Student
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &student)
			json.NewEncoder(w).Encode(student)
		}
	}
}

func DeleteUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		userID := mux.Vars(r)["id"]

		var user models.User
		db.First(&user, userID)
		db.Delete(&user)

		json.NewEncoder(w).Encode(user)
	}
}

func UpdateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["id"]
		var user models.User
		db.Model(&user).Preload("Connections").Preload("Subjects").Preload("Reviews").First(&user, userID)

		var updatedUser models.User
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&updatedUser)
		if err != nil {
			panic(err)
		}

		if updatedUser.Password != "" {
			password, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
			if err != nil {
				panic("Failed to hash password")
			}
			updatedUser.Password = string(password)
		}
		db.Model(&user).Updates(updatedUser)
		w.Header().Add("Content-Type", "application/json")

		if user.IsTutor {
			var tutor models.Tutor
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &tutor)
			json.NewEncoder(w).Encode(tutor)
		} else {
			var student models.Student
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &student)
			json.NewEncoder(w).Encode(student)
		}
	}
}

func UploadProfilePicture(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["id"]
		var user models.User
		db.First(&user, userID)

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error uploading file", http.StatusBadRequest)
			return
		}
		defer file.Close()
		filename := fmt.Sprintf("%s_%d_%s", userID, time.Now().Unix(), handler.Filename)

		f, err := os.OpenFile("uploads/"+filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer f.Close()
		io.Copy(f, file)

		user.ProfilePicture = filename
		db.Save(&user)
		w.Header().Set("Cache-Control", "no-cache")
		jsonResponse := map[string]string{"filename": filename}
		json.NewEncoder(w).Encode(jsonResponse)
	}
}

func GetProfilePicture(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["id"]
		var user models.User
		db.First(&user, userID)

		filename := user.ProfilePicture
		file, err := os.Open("uploads/" + filename)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
		}

		defer file.Close()

		contentType := mime.TypeByExtension(filepath.Ext(filename))
		if contentType == "" {
			contentType = "application/octet-stream"
		}

		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "no-cache")
		io.Copy(w, file)
	}
}
