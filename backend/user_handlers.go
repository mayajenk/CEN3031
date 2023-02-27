package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wader/gormstore/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	IsTutor    bool      `json:"is_tutor"`
	Rating     float64   `json:"rating"`
	Subjects   []Subject `gorm:"many2many:user_subjects;"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	About      string    `json:"about"`
	Grade      int32     `json:"grade"`
}

type Subject struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	var users []User
	db.Model(&User{}).Preload("Subjects").Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get User Endpoint Hit")

	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	userID := mux.Vars(r)["id"]

	var user User

	err = db.Model(&User{}).Preload("Subjects").First(&user, userID).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Error retrieving user.")
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User Endpoint Hit")

	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var user User
	err = decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	// Checking if a user is unique in the database
	var existingUser User
	result := db.Where("username = ?", user.Username).First(&existingUser)
	if result.Error == nil {
		fmt.Fprintf(w, "This username already exists. Please try a new one.")
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
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	userID := mux.Vars(r)["id"]

	var user User
	db.First(&user, userID)

	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	userID := mux.Vars(r)["id"]

	var user User
	db.First(&user, userID)

	newUser := user
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&newUser)
	if err != nil {
		panic(err)
	}

	if newUser.Password != user.Password {
		password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			panic("Failed to hash password")
		}
		newUser.Password = string(password)
	}

	user = newUser

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}

func login(store *gormstore.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
		if err != nil {
			panic("Failed to connect database")
		}

		var reqUser User
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err = decoder.Decode(&reqUser)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		var user User

		fmt.Println("Incoming login request")
		fmt.Println("Username: " + reqUser.Username)
		fmt.Println("Password: " + reqUser.Password)

		result := db.Where("username = ?", reqUser.Username).First(&user)
		err = result.Error
		if err != nil {
			http.Error(w, "Error 1", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		session, err := store.New(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["userID"] = user.ID
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
