package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mayajenk/CEN3031/models"
	"gorm.io/gorm"
)

func AddSubject(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var user models.User
		err := db.Model(&models.User{}).First(&user, id).Error
		if err != nil {
			http.Error(w, "User does not exist", http.StatusNotFound)
		}

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		var subject models.Subject
		err = decoder.Decode(&subject)
		if err != nil {
			http.Error(w, "Bad request format", http.StatusBadRequest)
		}

		var existingSubject models.Subject
		err = db.Model(&models.Subject{}).Where("name = ?", subject.Name).First(&existingSubject).Error
		if err == nil {
			// Subject already exists, so don't create a new one.
			db.Model(&user).Association("Subjects").Append(&existingSubject)
			json.NewEncoder(w).Encode(existingSubject)
			return
		} else {
			db.Create(&subject)
			db.Model(&user).Association("Subjects").Append(&subject)
			json.NewEncoder(w).Encode(subject)
			return
		}
	}
}

func UpdateSubjects(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var user models.User
		err := db.Model(&models.User{}).First(&user, id).Error
		if err != nil {
			http.Error(w, "User does not exist", http.StatusNotFound)
		}

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		var subjects []models.Subject
		decoder.Decode(&subjects)

		// Check if subjects already exist or not.
		for i, subject := range subjects {
			var existingSubject models.Subject
			err = db.Model(&models.Subject{}).Where("name = ?", subject.Name).First(&existingSubject).Error
			if err != nil {
				db.Create(&subject)
				existingSubject = subject
			}
			subjects[i] = existingSubject
		}
		db.Model(&user).Association("Subjects").Replace(subjects)
		json.NewEncoder(w).Encode(subjects)
	}
}

func DeleteSubjects(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var user models.User
		err := db.Model(&models.User{}).First(&user, id).Error
		if err != nil {
			http.Error(w, "User does not exist", http.StatusNotFound)
		}

		db.Model(&user).Association("Subjects").Clear()
	}
}
