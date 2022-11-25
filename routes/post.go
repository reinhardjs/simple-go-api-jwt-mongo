package routes

import (
	"simple-api/controllers"

	"github.com/gorilla/mux"
)

func PostRoute(router *mux.Router) {
	router.HandleFunc("/posts/{userId}", controllers.GetPost()).Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost()).Methods("POST")
	router.HandleFunc("/posts/{userId}", controllers.UpdatePost()).Methods("PUT")
	router.HandleFunc("/posts/{userId}", controllers.DeletePost()).Methods("DELETE")
}
