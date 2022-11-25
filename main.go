package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-api/middlewares"
	"simple-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// add JWTAuth middleware
	router.Use(middlewares.JwtAuthentication)

	// add RolePermissonChecker middleware
	router.Use(middlewares.RolePermissionCheck)

	// add routes
	routes.UserRoute(router)
	routes.PostRoute(router)

	fmt.Println("starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
