package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mayajenk/CEN3031/models"
	"gorm.io/gorm"
)

func AddReview(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var review models.Review
		err := json.NewDecoder(r.Body).Decode(&review)
		if err != nil {
			sendError(err.Error(), http.StatusBadRequest, w)
		}

		var reviewer models.User
		var reviewee models.User

		err = db.First(&reviewer, review.ReviewerID).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = db.First(&reviewee, review.RevieweeID).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = db.Model(&reviewee).Association("Connections").Find(&reviewer)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		db.Model(&reviewee).Association("Reviews").Append(&review)
	}
}

func DeleteReview(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := mux.Vars(r)["id"]

		var review models.Review
		db.Model(&models.Review{}).First(&review, id)
		db.Delete(&review)

		json.NewEncoder(w).Encode(review)
	}
}
