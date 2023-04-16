package main

import (
	"net/http"
	"os"

	"github.com/mayajenk/CEN3031/handlers"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
)

func httpHandler(store *gormstore.Store, db *gorm.DB) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/api/user", handlers.GetUserFromSession(store, db)).Methods("GET")
	router.HandleFunc("/api/users", handlers.GetAllUsers(db)).Methods("GET")
	router.HandleFunc("/api/users", handlers.NewUser(db)).Methods("POST")
	router.HandleFunc("/api/users/{id}", handlers.GetUser(db)).Methods("GET")
	router.HandleFunc("/api/users/{id}", handlers.DeleteUser(db)).Methods("DELETE")
	router.HandleFunc("/api/users/{id}", handlers.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/api/users/{id}/profile-picture", handlers.UploadProfilePicture(db)).Methods("POST")
	router.HandleFunc("/api/users/{id}/profile-picture", handlers.GetProfilePicture(db)).Methods("GET")
	router.HandleFunc("/api/users/{id}/subjects", handlers.AddSubject(db)).Methods("POST")
	router.HandleFunc("/api/users/{id}/subjects", handlers.UpdateSubjects(db)).Methods("PUT")
	router.HandleFunc("/api/users/{id}/subjects", handlers.DeleteSubjects(db)).Methods("DELETE")

	router.HandleFunc("/api/login", handlers.Login(store, db)).Methods("POST")
	router.HandleFunc("/api/search", handlers.SearchDatabase(db)).Methods("GET")
	router.HandleFunc("/api/logout", handlers.Logout(store)).Methods("POST")
	router.HandleFunc("/api/connection", handlers.AddConnection(db)).Methods("POST")
	router.HandleFunc("/api/connection", handlers.DeleteConnection(db)).Methods("DELETE")
	router.HandleFunc("/api/review", handlers.AddReview(db)).Methods("POST")
	router.HandleFunc("/api/review/{id}", handlers.DeleteReview(db)).Methods("DELETE")

	router.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	return ghandlers.LoggingHandler(os.Stdout,
		ghandlers.CORS(
			ghandlers.AllowCredentials(),
			ghandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
				"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
				"Cache-Control", "Content-Range", "Range"}),
			ghandlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			ghandlers.AllowedOrigins([]string{"http://localhost:8080"}),
			ghandlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			ghandlers.MaxAge(86400),
		)(router))
}
