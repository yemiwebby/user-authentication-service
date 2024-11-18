package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yemiwebby/user-authentication-service/internal/handlers"
)


func main() {
	r := mux.NewRouter()

	// Register routes
	handlers.RegisterAuthRoutes(r)

	// start server
	log.Println("User Authentication service running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}