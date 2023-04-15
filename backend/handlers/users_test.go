package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/gorilla/mux"
	"github.com/mayajenk/CEN3031/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func setupTestEnv() *gorm.DB {
	// Connect to a test database
	db, err := gorm.Open(sqlite.Open("../db/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Automatically create database tables
	db.AutoMigrate(&models.User{})

	return db
}

func TestNewUserHandler(t *testing.T) {
	// Set up a test environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback() // Resets the database on function exit

	// Set up a new router with the user creation handler
	r := mux.NewRouter()
	r.HandleFunc("/api/users", NewUser(tx)).Methods("POST")

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
	var user models.User
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

func TestDeleteUserHandler(t *testing.T) {
	// Set up a test environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback() // Resets the database on function exit

	// Set up a new router with the user creation and deletion handler
	r := mux.NewRouter()
	r.HandleFunc("/api/users", NewUser(tx)).Methods("POST")
	r.HandleFunc("/api/users/{id}", DeleteUser(tx)).Methods("DELETE")

	// Insert a new user into the database
	reqBody, _ := json.Marshal(map[string]any{
		"username": "foo",
		"password": "bar",
		"is_tutor": false,
	})

	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Verify that the new user was created in the database
	var user models.User
	result := tx.Where("username = ?", "foo").First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			t.Errorf("User was not added to the database")
		}
	}

	// Create a new test request to delete the user
	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/api/users/%d", user.ID), nil)

	// Send the request to the handler
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify that the user was deleted from the database
	result = tx.Where("username = ?", "foo").First(&user)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		t.Errorf("User was not deleted from the database")
	}
}

func TestUpdateUserHandler(t *testing.T) {
	// Set up a test environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback() // Resets the database on function exit

	// Set up a new router with the user update handler
	r := mux.NewRouter()
	r.HandleFunc("/api/users", NewUser(tx)).Methods("POST")
	r.HandleFunc("/api/users/{id}", UpdateUser(tx)).Methods("PUT")

	// Insert a new user into the database
	reqBody, _ := json.Marshal(map[string]any{
		"username": "foo",
		"password": "bar",
		"is_tutor": false,
	})

	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Verify that the new user was created in the database
	var user models.User
	result := tx.Where("username = ?", "foo").First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			t.Errorf("User was not added to the database")
		}
	}

	// Create a new test request with updated user data
	newUserData := models.User{Username: "updated_foo", Password: "updated_bar", IsTutor: true}
	reqBody, _ = json.Marshal(newUserData)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/api/users/%d", user.ID), bytes.NewBuffer(reqBody))
	fmt.Printf("ID: %d\n", user.ID)
	req.Header.Set("Content-Type", "application/json")

	// Send the request to the handler
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify that the user was updated in the database
	var updatedUser models.User
	tx.First(&updatedUser, user.ID)

	if updatedUser.Username != newUserData.Username {
		t.Errorf("Username was not updated in the database: got %v want %v", updatedUser.Username, newUserData.Username)
	}
	err := bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte(newUserData.Password))
	if err != nil {
		t.Errorf("Password was not updated in the database: %v", err.Error())
	}
	if updatedUser.IsTutor != newUserData.IsTutor {
		t.Errorf("IsTutor was not updated in the database: got %v want %v", updatedUser.IsTutor, newUserData.IsTutor)
	}
}

func TestGetUserHandler(t *testing.T) {
	// Set up a test environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback() // Resets the database on function exit

	// Set up a new router with the user creation and retrieval handlers
	r := mux.NewRouter()
	r.HandleFunc("/api/users", NewUser(tx)).Methods("POST")
	r.HandleFunc("/api/users/{id}", GetUser(tx)).Methods("GET")

	// Insert a new user into the database
	reqBody, _ := json.Marshal(map[string]any{
		"username": "foo",
		"password": "bar",
		"is_tutor": false,
	})

	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Verify that the new user was created in the database
	var user models.User
	result := tx.Where("username = ?", "foo").First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			t.Errorf("User was not added to the database")
		}
	}

	// Create a new test request to retrieve the user
	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/users/%d", user.ID), nil)

	// Send the request to the handler
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify that the returned user matches the test data
	var returnedUser models.User
	json.NewDecoder(rr.Body).Decode(&returnedUser)
	if returnedUser.Username != "foo" {
		t.Errorf("Username is incorrect: got %v want %v", returnedUser.Username, "foo")
	}
	err := bcrypt.CompareHashAndPassword([]byte(returnedUser.Password), []byte("bar"))
	if err != nil {
		t.Errorf("Password stored incorrectly: %v", err.Error())
	}
	if returnedUser.IsTutor != false {
		t.Errorf("IsTutor is incorrect: got %v want %v", returnedUser.IsTutor, "false")
	}
}

func TestGetAllUsersHandler(t *testing.T) {
	// Set up a test environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback() // Resets the database on function exit

	// Set up a new router with the user creation and retrieval handlers
	r := mux.NewRouter()
	r.HandleFunc("/api/users", NewUser(tx)).Methods("POST")
	r.HandleFunc("/api/users", GetAllUsers(tx)).Methods("GET")

	var user1 models.User = models.User{
		Username: "test",
		Password: "test1",
		IsTutor:  false,
	}

	// Insert a new user into the database
	reqBody, _ := json.Marshal(user1)

	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Insert another user into the database
	var user2 models.User = models.User{
		Username: "test2",
		Password: "test3",
		IsTutor:  true,
	}
	reqBody, _ = json.Marshal(user2)

	req, _ = http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Create a new test request to get all users
	req, _ = http.NewRequest("GET", "/api/users", nil)

	// Send the request to the handler
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse the response body into a slice of users
	var users []models.User
	json.NewDecoder(rr.Body).Decode(&users)

	// Verify that the correct number of users were returned
	if len(users) != 2 {
		t.Errorf("Incorrect number of users returned: got %v want %v", len(users), 2)
	}

	// Verify that the returned users match the test data
	if users[0].Username != "test" || bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte("test1")) != nil {
		t.Errorf("Incorrect user data returned for user1")
	}

	if users[1].Username != "test2" || bcrypt.CompareHashAndPassword([]byte(users[1].Password), []byte("test3")) != nil {
		t.Errorf("Incorrect user data returned for user2")
	}
}