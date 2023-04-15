package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mayajenk/CEN3031/models"
	"gorm.io/gorm"
)

// Adds a connection between user_1 and user_2 to the database
func AddConnection(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the user IDs from the request body
		var params struct {
			User1ID uint `json:"user_1"`
			User2ID uint `json:"user_2"`
		}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var user1 models.User
		var user2 models.User
		err = db.First(&user1, params.User1ID).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = db.First(&user2, params.User2ID).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		db.Model(&user1).Association("Connections").Append(&user2)
		db.Model(&user2).Association("Connections").Append(&user1)

		var users = []models.User{user1, user2}

		// Return the new connection object as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// Deletes the connection between user_1 and user_2 from the database.
func DeleteConnection(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			User1ID uint `json:"user_1"`
			User2ID uint `json:"user_2"`
		}

		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var user1 models.User
		var user2 models.User
		err = db.First(&user1, params.User1ID).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = db.First(&user2, params.User2ID).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		db.Model(&user1).Association("Connections").Delete(&user2)
		db.Model(&user2).Association("Connections").Delete(&user1)

		var users = []models.User{user1, user2}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
