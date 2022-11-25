package routes

import (
	"simple-api/controllers"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/users", controllers.CreateUser()).Methods("POST")
	router.HandleFunc("/token", controllers.GetToken()).Methods("GET")
}
