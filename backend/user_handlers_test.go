package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/gorilla/mux"
	"github.com/wader/gormstore/v2"
	"golang.org/x/crypto/bcrypt"
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

func TestDeleteUserHandler(t *testing.T) {
	// Set up a test environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback() // Resets the database on function exit

	// Set up a new router with the user creation and deletion handler
	r := mux.NewRouter()
	r.HandleFunc("/api/users", newUser(tx)).Methods("POST")
	r.HandleFunc("/api/users/{id}", deleteUser(tx)).Methods("DELETE")

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
	var user User
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
	r.HandleFunc("/api/users", newUser(tx)).Methods("POST")
	r.HandleFunc("/api/users/{id}", updateUser(tx)).Methods("PUT")

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
	var user User
	result := tx.Where("username = ?", "foo").First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			t.Errorf("User was not added to the database")
		}
	}

	// Create a new test request with updated user data
	newUserData := User{Username: "updated_foo", Password: "updated_bar", IsTutor: true}
	reqBody, _ = json.Marshal(newUserData)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/api/users/%d", user.ID), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Send the request to the handler
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify that the user was updated in the database
	var updatedUser User
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
	r.HandleFunc("/api/users", newUser(tx)).Methods("POST")
	r.HandleFunc("/api/users/{id}", getUser(tx)).Methods("GET")

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
	var user User
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
	var returnedUser User
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
	r.HandleFunc("/api/users", newUser(tx)).Methods("POST")
	r.HandleFunc("/api/users", getAllUsers(tx)).Methods("GET")

	var user1 User = User{
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
	var user2 User = User{
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
	var users []User
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

func TestLogin(t *testing.T) {
	// Set up testing environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback()

	// Set up session store
	sessionDB, err := gorm.Open(sqlite.Open("sessions.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	store := gormstore.New(sessionDB, []byte(os.Getenv("SESSION_KEY")))

	r := mux.NewRouter()
	r.HandleFunc("/api/users", newUser(tx)).Methods("POST")
	r.HandleFunc("/api/login", login(store, tx)).Methods("POST")

	var testUser User = User{
		Username: "foo",
		Password: "bar",
		IsTutor:  false,
	}
	reqBody, _ := json.Marshal(testUser)

	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	reqBody, _ = json.Marshal(map[string]any{
		"username": "foo",
		"password": "bar",
	})

	req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(reqBody))
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"is_tutor":false,"status":200}`
	if got := strings.TrimSpace(rr.Body.String()); got != expected {
		t.Errorf("Handler returned wrong body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestLogout(t *testing.T) {
	// Set up testing environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback()

	// Set up session store
	sessionDB, err := gorm.Open(sqlite.Open("sessions.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	store := gormstore.New(sessionDB, []byte(os.Getenv("SESSION_KEY")))

	r := mux.NewRouter()
	r.HandleFunc("/api/users", newUser(tx)).Methods("POST")
	r.HandleFunc("/api/login", login(store, tx)).Methods("POST")
	r.HandleFunc("/api/logout", logout(store)).Methods("POST")

	var testUser User = User{
		Username: "foo",
		Password: "bar",
		IsTutor:  false,
	}
	reqBody, _ := json.Marshal(testUser)

	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	reqBody, _ = json.Marshal(map[string]any{
		"username": "foo",
		"password": "bar",
	})

	req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(reqBody))
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	reqBody, _ = json.Marshal(map[string]any{})

	req, _ = http.NewRequest("POST", "/api/logout", bytes.NewBuffer(reqBody))
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	expected := `{"message":"Successfully logged out.","status":200}`
	if got := strings.TrimSpace(rr.Body.String()); got != expected {
		t.Errorf("Handler returned wrong body: got %v want %v", rr.Body.String(), expected)
	}

	// Check session values
	session, _ := store.Get(req, "session")
	if userID := session.Values["userID"]; userID != nil {
		t.Errorf("Session userID was not cleared: got %v want %v", userID, nil)
	}
}

// unit test for the search database function
func TestSearchDatabase(t *testing.T) {
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback()

	user1 := User{
		Username: "foo",
		Password: "bar",
		IsTutor:  false,
		Subjects: []Subject{
			{Name: "math"},
			{Name: "english"},
		},
		Rating: 5,
	}

	tx.Create(&user1)

	// Define a mock request and response
	req := httptest.NewRequest("GET", "/api/search?q=f", nil)
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/api/search", searchDatabase(tx)).Methods("GET")
	r.ServeHTTP(rr, req)
	fmt.Println(rr.Result().Body)

	// Assert that the response contains the expected data
	var users []User
	err := json.Unmarshal(rr.Body.Bytes(), &users)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("expected 1 users, got %d", len(users))
	}
}
