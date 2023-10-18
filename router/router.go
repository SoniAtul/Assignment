package router

import (
	"ProjectAssignment/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/api/posts", handlers.ListPosts).Methods("GET")
	r.HandleFunc("/api/posts/{id:[0-9]+}", handlers.GetPost).Methods("GET")
	r.HandleFunc("/api/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/api/posts/{id:[0-9]+}", handlers.UpdatePost).Methods("PUT")
	r.HandleFunc("/api/posts/{id:[0-9]+}", handlers.RemovePost).Methods("DELETE")
}
