package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Tests the getAllUsers handler, which should return a JSON list of
// all the users in the database.
func TestGetAllUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllUsers)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"username":"foo","password":"$2a$10$0SadYd0ltE.e56N.SELC3OmeyutxO7EgHt0HClGvb9gqkgQVeKZZ.","first_name":"","last_name":"","is_tutor":false,"rating":0,"Subjects":[],"email":"","phone":"","about":"","grade":0},{"username":"bar","password":"$2a$10$wHNOiBQq5Ybjo53j9NiEDulhb73aj7CSCOfOjwwyzriof3wBkm4im","first_name":"","last_name":"","is_tutor":false,"rating":0,"Subjects":[],"email":"","phone":"","about":"","grade":0}]`
	got := strings.TrimSpace(rr.Body.String())

	if got != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Tests the getUser handler at id = 1, which should have username: foo, password: bar
func TestGetUser1(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// For URL variables, you have to create a router so the vars will be added to the context.
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}", getUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"username":"foo","password":"$2a$10$0SadYd0ltE.e56N.SELC3OmeyutxO7EgHt0HClGvb9gqkgQVeKZZ.","first_name":"","last_name":"","is_tutor":false,"rating":0,"Subjects":[],"email":"","phone":"","about":"","grade":0}`
	got := strings.TrimSpace(rr.Body.String())

	if got != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Tests getUser handler at id = 2, which should have username: bar, password: foo
func TestGetUser2(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}", getUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"username":"bar","password":"$2a$10$wHNOiBQq5Ybjo53j9NiEDulhb73aj7CSCOfOjwwyzriof3wBkm4im","first_name":"","last_name":"","is_tutor":false,"rating":0,"Subjects":[],"email":"","phone":"","about":"","grade":0}`
	got := strings.TrimSpace(rr.Body.String())

	if got != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// Tests getUser handler with an invalid id.
func TestGetUser3(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/users/{id}", getUser)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `"Error retrieving user."`
	got := strings.TrimSpace(rr.Body.String())

	if got != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// creating a test database
func TestMain(m *testing.M) {
	// Open a test database connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to test database:", err)
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Run the tests
	exitCode := m.Run()

	// Exit with the appropriate code
	os.Exit(exitCode)
}

func TestNewUser(t *testing.T) {
	// Set up a mock HTTP request and response
	reqBody := []byte(`{"username": "testuser", "password": "testpass"}`)
	req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	// Set up a mock database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	// Defer the cleanup function to ensure the database is closed and deleted
	defer func() {

		err = os.Remove("test.db")
		if err != nil {
			t.Errorf("failed to delete database: %v", err)
		}
	}()

	// Call the newUser function with the mock request and response
	newUser(db)

	// Check the response from the newUser function
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.StatusCode)
	}

	// Check that the user was added to the database
	var user User
	result := db.Where("username = ?", "testuser").First(&user)
	if result.Error != nil {
		t.Errorf("failed to retrieve user: %v", result.Error)
	}

	// Check that the user password was hashed correctly
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("testpass"))
	if err != nil {
		t.Errorf("failed to compare hashed password: %v", err)
	}
}

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

	// Set up a new router with the user creation handler
	r := mux.NewRouter()
	r.HandleFunc("/api/users", newUser(db)).Methods("POST")

	// Create a new test request with sample data
	reqBody, _ := json.Marshal(map[string]string{
		"name":     "food",
		"password": "bar",
	})
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Send the request to the handler
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Verify that a new user was created in the database
	var users []User
	db.Find(&users)
	if len(users) != 1 {
		t.Errorf("handler did not create a new user in the database")
	}
}
