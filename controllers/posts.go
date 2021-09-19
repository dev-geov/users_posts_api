package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dev-geov/users_posts_api/models"
	"github.com/gorilla/mux"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	var posts []models.Post
	models.DB.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var post models.Post
	models.DB.Preload("Author").First(&post, params["id"])
	json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	models.DB.Omit("author").Create(&post)
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var post models.Post
	models.DB.First(&post, params["id"])
	json.NewDecoder(r.Body).Decode(&post)
	models.DB.Save(&post)
	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var post models.Post
	models.DB.Delete(&post, params["id"])
	json.NewEncoder(w).Encode("The Post is Deleted Successfully!")
}
