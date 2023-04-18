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

		db.Joins("JOIN user_subjects ON users.id = user_subjects.user_id").
			Joins("JOIN subjects ON user_subjects.subject_id = subjects.id").
			Where("subjects.name LIKE ?", "%"+subject+"%").
			Select("DISTINCT users.*").
			Preload("Subjects").
			Find(&users)

		var tutors []models.Tutor

		for _, user := range users {
			var tutor models.Tutor
			temp, _ := json.Marshal(user)
			err := json.Unmarshal(temp, &tutor)

			if err == nil {
				tutors = append(tutors, tutor)
			}
		}
		json.NewEncoder(w).Encode(tutors)
	}
}
