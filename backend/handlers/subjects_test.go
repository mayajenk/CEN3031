package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mayajenk/CEN3031/models"
)

func TestAddSubjectHandler(t *testing.T) {
	// Set up a test environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback() // Resets the database on function exit

	// Insert a new user into the database
	var user models.User
	user.Username = "testuser"
	user.Password = "password"
	tx.Create(&user)

	// Insert an existing subject into the database
	var existingSubject models.Subject
	existingSubject.Name = "math"
	tx.Create(&existingSubject)

	// Set up a new router with the subject handler
	r := mux.NewRouter()
	r.HandleFunc("/api/users/{id}/subjects", AddSubject(tx)).Methods("POST")

	// Create a new test request to add a subject to the user's list of subjects
	newSubject := models.Subject{Name: "history"}
	reqBody, _ := json.Marshal(newSubject)
	req, _ := http.NewRequest("POST", fmt.Sprintf("/api/users/%d/subjects", user.ID), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify that the response body contains the new subject
	var addedSubject models.Subject
	json.Unmarshal(rr.Body.Bytes(), &addedSubject)
	if addedSubject.Name != newSubject.Name {
		t.Errorf("Handler returned wrong subject: got %v want %v", addedSubject, newSubject)
	}

	// Create a new test request to add the existing subject to the user's list of subjects
	reqBody, _ = json.Marshal(existingSubject)
	req, _ = http.NewRequest("POST", fmt.Sprintf("/api/users/%d/subjects", user.ID), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify that the response body contains the existing subject
	var addedExistingSubject models.Subject
	json.Unmarshal(rr.Body.Bytes(), &addedExistingSubject)
	if addedExistingSubject.Name != existingSubject.Name {
		t.Errorf("Handler returned wrong subject: got %v want %v", addedExistingSubject, existingSubject)
	}

	// Check that the existing subject was added to the user's list of subjects
	tx.Model(&user).Association("Subjects").Find(&existingSubject)
	if existingSubject.ID == 0 {
		t.Errorf("Existing subject was not added to user's subjects list")
	}
}

func TestUpdateSubjectsHandler(t *testing.T) {
	// Set up a test environment
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback() // Resets the database on function exit

	// Insert a new user into the database
	user := models.User{
		Username: "foo",
		Password: "bar",
	}
	tx.Create(&user)

	// Insert some subjects into the database
	subjects := []models.Subject{
		{Name: "Math"},
		{Name: "Science"},
		{Name: "English"},
	}
	for i := range subjects {
		tx.Create(&subjects[i])
	}

	// Set up a new router with the update subjects handler
	r := mux.NewRouter()
	r.HandleFunc("/api/users/{id}/subjects", UpdateSubjects(tx)).Methods("PUT")

	// Create a new test request to update the user's subjects
	subjectsToUpdate := []models.Subject{
		{Name: "Math"},
		{Name: "History"},
	}
	reqBody, _ := json.Marshal(subjectsToUpdate)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/users/%d/subjects", user.ID), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Verify that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify that the user's subjects were updated
	var updatedUser models.User
	tx.Preload("Subjects").First(&updatedUser, user.ID)
	if len(updatedUser.Subjects) != len(subjectsToUpdate) {
		t.Errorf("User's subjects were not updated: got %v want %v", updatedUser.Subjects, subjectsToUpdate)
	}
	for i, subject := range updatedUser.Subjects {
		if subject.Name != subjectsToUpdate[i].Name {
			t.Errorf("User's subjects were not updated: got %v want %v", updatedUser.Subjects, subjectsToUpdate)
		}
	}
}
