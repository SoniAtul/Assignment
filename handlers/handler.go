package handlers

import (
	"ProjectAssignment/database"
	"ProjectAssignment/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func ListPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	database.DB.Find(&posts)
	jsonResponse(w, posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	database.DB.First(&post, params["id"])
	jsonResponse(w, post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	jsonDecode(r.Body, &post)
	database.DB.Create(&post)
	jsonResponse(w, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	database.DB.First(&post, params["id"])
	jsonDecode(r.Body, &post)
	database.DB.Save(&post)
	jsonResponse(w, post)
}

func RemovePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	database.DB.Delete(&post, params["id"])
	w.WriteHeader(http.StatusNoContent)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func jsonDecode(body io.ReadCloser, v interface{}) {
	decoder := json.NewDecoder(body)
	decoder.Decode(v)
}
