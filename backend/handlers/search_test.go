package handlers

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mayajenk/CEN3031/models"
)

// unit test for the search database function
func TestSearchDatabase(t *testing.T) {
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback()

	user1 := models.User{
		Username: "foo",
		Password: "bar",
		IsTutor:  false,
		Subjects: []models.Subject{
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
	r.HandleFunc("/api/search", SearchDatabase(tx)).Methods("GET")
	r.ServeHTTP(rr, req)
	fmt.Println(rr.Result().Body)

	// Assert that the response contains the expected data
	var users []models.User
	err := json.Unmarshal(rr.Body.Bytes(), &users)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("expected 1 users, got %d", len(users))
	}
}
