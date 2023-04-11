package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/wader/gormstore/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	IsTutor     bool      `json:"is_tutor"`
	Rating      float64   `json:"rating"`
	Subjects    []Subject `gorm:"many2many:user_subjects" json:"subjects"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Contact     string    `json:"contact"`
	About       string    `json:"about"`
	Grade       int32     `json:"grade"`
	Connections []*User   `gorm:"many2many:user_connections" json:"connections"`
	Price       string    `json:"price"`
	Reviews     []Review  `gorm:"many2many:user_reviews" json:"reviews"`
}

type TutorView struct {
	ID          int32          `json:"id"`
	Username    string         `json:"username"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	IsTutor     bool           `json:"is_tutor"`
	Rating      float64        `json:"rating"`
	Subjects    []Subject      `json:"subjects"`
	Email       string         `json:"email"`
	Phone       string         `json:"phone"`
	Contact     string         `json:"contact"`
	About       string         `json:"about"`
	Price       string         `json:"price"`
	Connections []*StudentView `json:"connections"`
	Reviews     []Review       `json:"reviews"`
}

type StudentView struct {
	ID          int32        `json:"id"`
	Username    string       `json:"username"`
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	IsTutor     bool         `json:"is_tutor"`
	Rating      float64      `json:"rating"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
	About       string       `json:"about"`
	Grade       int32        `json:"grade"`
	Connections []*TutorView `json:"connections"`
	Reviews     []Review     `json:"reviews"`
}

type Subject struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
}

type Review struct {
	gorm.Model
	ReviewerID uint    `gorm:"index" json:"reviewer_id"`
	RevieweeID uint    `gorm:"index" json:"reviewee_id"`
	ReviewText string  `json:"review_text"`
	Rating     float64 `json:"rating"`
}

func sendError(message string, status int, w http.ResponseWriter) {
	res := make(map[string]any)
	res["message"] = message
	res["status"] = status
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func getAllUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var users []User
		db.Model(&User{}).Preload("Subjects").Preload("Connections").Preload("Reviews").Find(&users)

		json.NewEncoder(w).Encode(users)
	}
}

func getUserFromSession(store *gormstore.Store, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		session, err := store.Get(r, "session")
		if err != nil {
			sendError("Error retrieving user", http.StatusUnauthorized, w)
			return
		}

		userID := session.Values["userID"]
		var user User

		err = db.Model(&User{}).Preload("Subjects").Preload("Connections").Preload("Reviews").First(&user, userID).Error
		if err != nil {
			sendError("Error retrieving user", http.StatusUnauthorized, w)
			return
		}

		if user.IsTutor {
			var tutor TutorView
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &tutor)

			if err == nil {
				json.NewEncoder(w).Encode(tutor)
			}
		} else {
			var student StudentView
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &student)

			if err == nil {
				json.NewEncoder(w).Encode(student)
			}
		}
	}
}

func getUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get User Endpoint Hit")
		w.Header().Set("Content-Type", "application/json")

		userID := mux.Vars(r)["id"]

		var user User

		err := db.Model(&User{}).Preload("Subjects").Preload("Connections").Preload("Reviews").First(&user, userID).Error
		if err != nil {
			sendError("Error retrieving user", http.StatusUnauthorized, w)
		} else {
			json.NewEncoder(w).Encode(user)
		}
	}

}

func newUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Println("New User Endpoint Hit")

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		var user User
		err := decoder.Decode(&user)
		if err != nil {
			sendError("Bad request format", http.StatusBadRequest, w)
			return
		}

		// Checking if a user is unique in the database
		var existingUser User
		result := db.Where("username = ?", user.Username).First(&existingUser)
		if result.Error == nil {
			sendError("Username already exists", http.StatusConflict, w)
			return
		} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			panic(result.Error)
		}

		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			panic("Failed to hash password")
		}
		user.Password = string(password)

		db.Create(&user)

		json.NewEncoder(w).Encode(user)
	}
}

func deleteUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		userID := mux.Vars(r)["id"]

		var user User
		db.First(&user, userID)
		db.Delete(&user)

		json.NewEncoder(w).Encode(user)
	}
}

func updateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["id"]

		var user User
		db.First(&user, userID)

		newUser := user
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&newUser)
		if err != nil {
			panic(err)
		}

		if newUser.Password != user.Password {
			password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
			if err != nil {
				panic("Failed to hash password")
			}
			newUser.Password = string(password)
		}

		user = newUser

		db.Save(&user)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func login(store *gormstore.Store, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqUser User
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		res := make(map[string]any)

		err := decoder.Decode(&reqUser)
		if err != nil {
			sendError("Bad request format", http.StatusBadRequest, w)
			return
		}

		var user User

		fmt.Println("Incoming login request")
		fmt.Println("Username: " + reqUser.Username)
		fmt.Println("Password: " + reqUser.Password)

		result := db.Where("username = ?", reqUser.Username).First(&user)

		err = result.Error
		if err != nil {
			sendError("Username or password was incorrect", http.StatusUnauthorized, w)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password))
		if err != nil {
			sendError("Username or password was incorrect", http.StatusUnauthorized, w)
			return
		}

		session, err := store.New(r, "session")
		if err != nil {
			sendError(err.Error(), http.StatusInternalServerError, w)
			return
		}

		session.Options = &sessions.Options{
			SameSite: http.SameSiteLaxMode,
			HttpOnly: false,
			Secure:   false,
			Path:     "/",
		}

		session.Values["userID"] = user.ID
		session.Values["authenticated"] = true

		err = session.Save(r, w)
		if err != nil {
			sendError(err.Error(), http.StatusInternalServerError, w)
			return
		}

		db.Model(&User{}).Preload("Subjects").Preload("Connections").Preload("Reviews").First(&user, user.ID)
		if user.IsTutor {
			var tutor TutorView
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &tutor)

			if err == nil {
				res["user"] = tutor
				res["status"] = http.StatusOK
			}
		} else {
			var student StudentView
			temp, _ := json.Marshal(user)
			err = json.Unmarshal(temp, &student)

			if err == nil {
				res["user"] = student
				res["status"] = http.StatusOK
			}
		}
		json.NewEncoder(w).Encode(res)
	}
}

// search function if the user wants to look for a particular tutor or a subject
func searchDatabase(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		query := r.URL.Query().Get("q")
		subject := r.URL.Query().Get("subject")

		var users []User

		if subject == "" {
			db.Where("username LIKE ?", "%"+query+"%").Find(&users)
		} else {
			db.Where("username LIKE ? AND subject LIKE ?", "%"+query+"%", "%"+subject+"%")
		}

		json.NewEncoder(w).Encode(users)
	}
}

// function for the user to log out of the website
func logout(store *gormstore.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Options.MaxAge = -1
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["userID"] = nil
		session.Values["authenticated"] = false
		session.Options = &sessions.Options{
			MaxAge:   -1,
			HttpOnly: true,
		}

		w.Header().Set("Content-Type", "application/json")

		res := make(map[string]interface{})
		res["message"] = "Successfully logged out."
		res["status"] = http.StatusOK
		json.NewEncoder(w).Encode(res)

	}
}

// Adds a connection between user_1 and user_2 to the database
func addConnection(db *gorm.DB) http.HandlerFunc {
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

		var user1 User
		var user2 User
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

		var users = []User{user1, user2}

		// Return the new connection object as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// Deletes the connection between user_1 and user_2 from the database.
func deleteConnection(db *gorm.DB) http.HandlerFunc {
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

		var user1 User
		var user2 User
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

		var users = []User{user1, user2}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func addReview(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var review Review
		err := json.NewDecoder(r.Body).Decode(&review)
		if err != nil {
			sendError(err.Error(), http.StatusBadRequest, w)
		}

		var reviewer User
		var reviewee User

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

func deleteReview(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := mux.Vars(r)["id"]

		var review Review
		db.Model(&Review{}).First(&review, id)
		db.Delete(&review)

		json.NewEncoder(w).Encode(review)
	}
}
