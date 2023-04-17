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
		subject := r.URL.Query().Get("subject")

		var users []models.User

		if subject != "" {
			db.Joins("JOIN user_subjects ON users.id = user_subjects.user_id").
				Joins("JOIN subjects ON user_subjects.subject_id = subjects.id").
				Where("subjects.name LIKE ?", "%"+subject+"%").
				Preload("Subjects").
				Find(&users)
		}
		json.NewEncoder(w).Encode(users)
	}
}
