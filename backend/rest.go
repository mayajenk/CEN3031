package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
)

func httpHandler(store *gormstore.Store, db *gorm.DB) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/api/user", getUserFromSession(store, db)).Methods("GET")
	router.HandleFunc("/api/users", getAllUsers(db)).Methods("GET")
	router.HandleFunc("/api/users", newUser(db)).Methods("POST")
	router.HandleFunc("/api/users/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/api/users/{id}", deleteUser(db)).Methods("DELETE")
	router.HandleFunc("/api/users/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/api/login", login(store, db)).Methods("POST")
	router.HandleFunc("/api/search", searchDatabase(db)).Methods("GET")
	router.HandleFunc("/api/logout", logout(store)).Methods("POST")
	router.HandleFunc("/api/connection", addConnection(db)).Methods("POST")
	router.HandleFunc("/api/connection", deleteConnection(db)).Methods("DELETE")

	router.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
				"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
				"Cache-Control", "Content-Range", "Range"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:8080"}),
			handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			handlers.MaxAge(86400),
		)(router))
}
