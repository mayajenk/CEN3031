package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestEnv() *gorm.DB {
	// Connect to a test database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Automatically create database tables
	db.AutoMigrate(&User{})

	return db
}

func TestNewUserHandler(t *testing.T) {
	// Set up a test environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback() // Resets the database on function exit

	// Set up a new router with the user creation handler
	r := mux.NewRouter()
	r.HandleFunc("/api/users", newUser(tx)).Methods("POST")

	// Create a new test request with sample data
	reqBody, _ := json.Marshal(map[string]any{
		"username": "foo",
		"password": "bar",
		"is_tutor": false,
	})
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Send the request to the handler
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Verify that the new user was created in the database
	var user User
	result := tx.Where("username = ?", "foo").First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			t.Errorf("User was not added to the database")
		}
	}

	// Verify that user fields are correct
	if user.Username != "foo" {
		t.Errorf("Username is incorrect: got %v want %v", user.Username, "foo")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("bar"))
	if err != nil {
		t.Errorf("Password stored incorrectly: %v", err.Error())
	}
	if user.IsTutor != false {
		t.Errorf("IsTutor is incorrect: got %v want %v", user.IsTutor, "false")
	}
}
