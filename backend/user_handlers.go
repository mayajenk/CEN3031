package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `json:"username"`
	Password string  `json:"password"`
	IsTutor  bool    `json:"is_tutor"`
	Rating   float64 `json:"rating"`
	Subjects string  `json:"subjects"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	About    string  `json:"about"`
	Grade    string  `json:"grade"`
}

func getAllUsers(w http.ResponseWriter, req *http.Request) {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	var users []User
	db.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Get User Endpoint Hit")

	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	userID := mux.Vars(req)["id"]

	var user User

	db.First(&user, userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func newUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println("New User Endpoint Hit")

	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	var user User
	err = decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	db.Create(&user)
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, req *http.Request) {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	userID := mux.Vars(req)["id"]

	var user User
	db.First(&user, userID)

	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, req *http.Request) {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	userID := mux.Vars(req)["id"]

	var user User
	db.First(&user, userID)

	newUser := user
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&newUser)
	if err != nil {
		panic(err)
	}

	user = newUser

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}
