package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/mayajenk/CEN3031/models"
	"github.com/wader/gormstore/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(store *gormstore.Store, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqUser models.User
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		res := make(map[string]any)

		err := decoder.Decode(&reqUser)
		if err != nil {
			sendError("Bad request format", http.StatusBadRequest, w)
			return
		}

		var user models.User

		fmt.Println("Incoming login request")
		fmt.Println("Username: " + reqUser.Username)
		fmt.Println("Password: " + reqUser.Password)

		result := db.Where("username = ?", reqUser.Username).First(&user)

		err = result.Error
		if err != nil {
			sendError("Username or password was incorrect", http.StatusUnauthorized, w)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password))
		if err != nil {
			sendError("Username or password was incorrect", http.StatusUnauthorized, w)
			return
		}

		session, err := store.New(r, "session")
		if err != nil {
			sendError(err.Error(), http.StatusInternalServerError, w)
			return
		}

		session.Options = &sessions.Options{
			SameSite: http.SameSiteLaxMode,
			HttpOnly: false,
			Secure:   false,
			Path:     "/",
		}

		session.Values["userID"] = user.ID
		session.Values["authenticated"] = true

		err = session.Save(r, w)
		if err != nil {
			sendError(err.Error(), http.StatusInternalServerError, w)
			return
		}

		db.Model(&models.User{}).Preload("Subjects").Preload("Connections").Preload("Reviews").First(&user, user.ID)
		if user.IsTutor {
			var tutor models.Tutor
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &tutor)

			if err == nil {
				res["user"] = tutor
				res["status"] = http.StatusOK
			}
		} else {
			var student models.Student
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &student)

			if err == nil {
				res["user"] = student
				res["status"] = http.StatusOK
			}
		}
		json.NewEncoder(w).Encode(res)
	}
}

// function for the user to log out of the website
func Logout(store *gormstore.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Options.MaxAge = -1
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["userID"] = nil
		session.Values["authenticated"] = false
		session.Options = &sessions.Options{
			MaxAge:   -1,
			HttpOnly: true,
		}

		w.Header().Set("Content-Type", "application/json")

		res := make(map[string]interface{})
		res["message"] = "Successfully logged out."
		res["status"] = http.StatusOK
		json.NewEncoder(w).Encode(res)

	}
}
