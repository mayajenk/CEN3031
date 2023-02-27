package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
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
