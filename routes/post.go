package routes

import (
	"simple-api/controllers"

	"github.com/gorilla/mux"
)

func PostRoute(router *mux.Router) {
	router.HandleFunc("/posts", controllers.CreatePost()).Methods("POST")
}
