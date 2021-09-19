package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dev-geov/users_posts_api/controllers"
	"github.com/dev-geov/users_posts_api/models"
)

func Routers() {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")

	log.Println("App running on port 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func main() {
	models.InitialConfig()
	Routers()
}
