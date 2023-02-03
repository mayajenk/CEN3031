package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	IsTutor  bool
}

func getAllUsers(w http.ResponseWriter, req *http.Request) {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

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

	fmt.Println(user)
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

	var newUser User
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&newUser)
	if err != nil {
		panic(err)
	}

	user.Username = newUser.Username
	user.Password = newUser.Password
	user.IsTutor = newUser.IsTutor

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/users", newUser).Methods("POST")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")

	fmt.Printf("Starting on http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initialMigration() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func main() {
	initialMigration()
	handleRequests()
}
