package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-api/app"
	"simple-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// add JWTAuth middleware
	router.Use(app.JwtAuthentication)

	// add routes
	routes.UserRoute(router)

	fmt.Println("starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
