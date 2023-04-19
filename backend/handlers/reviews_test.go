package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mayajenk/CEN3031/models"
)

func TestAddReview(t *testing.T) {
	// Set up test database and router
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback() // Resets the database on function exit
	router := mux.NewRouter()
	router.HandleFunc("/api/review", AddReview(tx)).Methods("POST")

	// Create two users
	reviewer := models.User{Username: "reviewer", Password: "password"}
	reviewee := models.User{Username: "reviewee", Password: "password"}
	tx.Create(&reviewer)
	tx.Create(&reviewee)

	// Create a review from reviewer to reviewee
	review := models.Review{
		ReviewerID: reviewer.ID,
		RevieweeID: reviewee.ID,
		Rating:     4.5,
		ReviewText: "This is a great review!",
	}

	// Encode the review as JSON
	body, _ := json.Marshal(&review)

	// Make a POST request to create the review
	req, err := http.NewRequest("POST", "/api/review", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check that the response status is 200 OK
	if rr.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	// Check that the reviewee now has a rating based on the review
	var updatedReviewee models.User
	tx.Preload("Reviews").Find(&updatedReviewee, reviewee.ID)
	if updatedReviewee.Rating != 4.5 {
		t.Errorf("Unexpected rating for reviewee: got %v, want 4.5", updatedReviewee.Rating)
	}
}
