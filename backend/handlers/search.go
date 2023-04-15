package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mayajenk/CEN3031/models"
	"gorm.io/gorm"
)

// search function if the user wants to look for a particular tutor or a subject
func SearchDatabase(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		query := r.URL.Query().Get("q")
		subject := r.URL.Query().Get("subject")

		var users []models.User

		if subject == "" {
			db.Where("username LIKE ?", "%"+query+"%").Find(&users)
		} else {
			db.Where("username LIKE ? AND subject LIKE ?", "%"+query+"%", "%"+subject+"%")
		}

		json.NewEncoder(w).Encode(users)
	}
}
