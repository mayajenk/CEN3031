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

func TestAddConnection(t *testing.T) {
	// Set up test database and router
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback() // Resets the database on function exit
	router := mux.NewRouter()
	router.HandleFunc("/api/connection", AddConnection(tx)).Methods("POST")

	// Create two users
	user1 := models.User{Username: "user1", Password: "password"}
	user2 := models.User{Username: "user2", Password: "password"}
	tx.Create(&user1)
	tx.Create(&user2)

	tx.Where("username = ?", user1.Username).First(&user1)
	tx.Where("username = ?", user2.Username).First(&user2)

	// Make a POST request to create a connection between the users
	body, _ := json.Marshal(map[string]uint{"user_1": user1.ID, "user_2": user2.ID})
	req, err := http.NewRequest("POST", "/api/connection", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check that the response status is 200 OK
	if rr.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", rr.Code, http.StatusOK)
	}
	// Check that the users are now connected in the database
	var user1Connections []models.User
	tx.Model(&user1).Association("Connections").Find(&user1Connections)
	if len(user1Connections) != 1 {
		t.Errorf("Unexpected number of connections for user1: got %v, want 1", len(user1Connections))
	}
	if user1Connections[0].ID != user2.ID {
		t.Errorf("Unexpected connection for user1: got %v, want %v", user1Connections[0].ID, user2.ID)
	}
	var user2Connections []models.User
	tx.Model(&user2).Association("Connections").Find(&user2Connections)
	if len(user2Connections) != 1 {
		t.Errorf("Unexpected number of connections for user2: got %v, want 1", len(user2Connections))
	}
	if user2Connections[0].ID != user1.ID {
		t.Errorf("Unexpected connection for user2: got %v, want %v", user2Connections[0].ID, user1.ID)
	}
}

func TestDeleteConnection(t *testing.T) {
	db := setupTestEnv()
	tx := db.Begin()
	defer tx.Rollback()

	router := mux.NewRouter()
	router.HandleFunc("/api/connection", DeleteConnection(tx)).Methods("DELETE")
	router.HandleFunc("/api/connection", AddConnection(tx)).Methods("POST")

	user1 := models.User{Username: "user1", Password: "password"}
	user2 := models.User{Username: "user2", Password: "password"}
	tx.Create(&user1)
	tx.Create(&user2)

	tx.Where("username = ?", user1.Username).First(&user1)
	tx.Where("username = ?", user2.Username).First(&user2)

	// Make a POST request to create a connection between the users
	body, _ := json.Marshal(map[string]uint{"user_1": user1.ID, "user_2": user2.ID})
	req, err := http.NewRequest("POST", "/api/connection", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check that the users are now connected in the database
	var user1Connections []models.User
	tx.Model(&user1).Association("Connections").Find(&user1Connections)
	if len(user1Connections) != 1 {
		t.Errorf("Unexpected number of connections for user1: got %v, want 1", len(user1Connections))
	}
	if user1Connections[0].ID != user2.ID {
		t.Errorf("Unexpected connection for user1: got %v, want %v", user1Connections[0].ID, user2.ID)
	}
	var user2Connections []models.User
	tx.Model(&user2).Association("Connections").Find(&user2Connections)
	if len(user2Connections) != 1 {
		t.Errorf("Unexpected number of connections for user2: got %v, want 1", len(user2Connections))
	}
	if user2Connections[0].ID != user1.ID {
		t.Errorf("Unexpected connection for user2: got %v, want %v", user2Connections[0].ID, user1.ID)
	}

	req, err = http.NewRequest("DELETE", "/api/connection", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check that connections are deleted
	tx.Model(&user1).Association("Connections").Find(&user1Connections)
	if len(user1Connections) != 0 {
		t.Errorf("Unexpected number of connections for user1: got %v, want 0", len(user1Connections))
	}

	tx.Model(&user2).Association("Connections").Find(&user2Connections)
	if len(user2Connections) != 0 {
		t.Errorf("Unexpected number of connections for user2: got %v, want 0", len(user2Connections))
	}
}
