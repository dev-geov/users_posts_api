package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dev-geov/users_posts_api/models"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	models.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	models.DB.Preload("Posts").First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	models.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	models.DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	models.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	models.DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
}
