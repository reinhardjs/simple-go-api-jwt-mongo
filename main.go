package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// add routes
	routes.UserRoute(router)

	fmt.Println("starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
