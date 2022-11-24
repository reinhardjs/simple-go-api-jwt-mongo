package routes

import (
	"simple-api/controllers"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/users", controllers.CreateUser()).Methods("POST")
}
