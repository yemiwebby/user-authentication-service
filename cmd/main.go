package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yemiwebby/user-authentication-service/config"
	"github.com/yemiwebby/user-authentication-service/internal/handlers"
)


func main() {
	config.LoadConfig()
	r := mux.NewRouter()
	handlers.RegisterAuthRoutes(r)

	log.Println("User Authentication service running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}