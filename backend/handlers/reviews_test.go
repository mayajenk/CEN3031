package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mayajenk/CEN3031/models"
	"gorm.io/gorm"
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

func TestDeleteReview(t *testing.T) {
	// Create a test database
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback()

	// Create a test review and user
	review := models.Review{Rating: 3.5, ReviewText: "Test review"}
	user := models.User{Username: "Test user", Email: "test@test.com", Reviews: []models.Review{review}}

	// Save the review and user to the database
	err := tx.Create(&user).Error
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Create a request with the review ID
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/reviews/%d", review.ID), nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the DeleteReview function with the test database and request
	handler := DeleteReview(tx)
	handler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check that the review was deleted from the database
	var deletedReview models.Review
	err = tx.First(&deletedReview, review.ID).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("Review was not deleted from the database")
	}

	// Check that the user's rating was updated
	var updatedUser models.User
	err = tx.First(&updatedUser, user.ID).Error
	if err != nil {
		t.Fatalf("Failed to retrieve updated user: %v", err)
	}
	expectedRating := 0.0
	if len(updatedUser.Reviews) > 0 {
		for _, r := range updatedUser.Reviews {
			expectedRating += r.Rating
		}
		expectedRating /= float64(len(updatedUser.Reviews))
	}
	if updatedUser.Rating != expectedRating {
		t.Errorf("User's rating was not updated: got %v want %v", updatedUser.Rating, expectedRating)
	}
}
