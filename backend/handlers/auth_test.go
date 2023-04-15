package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/gorilla/mux"
	"github.com/mayajenk/CEN3031/models"
	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
)

func TestLogin(t *testing.T) {
	// Set up testing environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback()

	// Set up session store
	sessionDB, err := gorm.Open(sqlite.Open("../db/sessions.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	store := gormstore.New(sessionDB, []byte(os.Getenv("SESSION_KEY")))

	r := mux.NewRouter()
	r.HandleFunc("/api/users", NewUser(tx)).Methods("POST")
	r.HandleFunc("/api/login", Login(store, tx)).Methods("POST")

	var testUser models.User = models.User{
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

	reqBody, _ = json.Marshal(map[string]any{
		"username": "foo1",
		"password": "bar1",
	})

	req, _ = http.NewRequest("POST", "/api/login", bytes.NewBuffer(reqBody))
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}

func TestLogout(t *testing.T) {
	// Set up testing environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback()

	// Set up session store
	sessionDB, err := gorm.Open(sqlite.Open("../db/sessions.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	store := gormstore.New(sessionDB, []byte(os.Getenv("SESSION_KEY")))

	r := mux.NewRouter()
	r.HandleFunc("/api/users", NewUser(tx)).Methods("POST")
	r.HandleFunc("/api/login", Login(store, tx)).Methods("POST")
	r.HandleFunc("/api/logout", Logout(store)).Methods("POST")

	var testUser models.User = models.User{
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
