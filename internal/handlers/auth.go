package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yemiwebby/user-authentication-service/internal/service"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/registration", RegisterUser).Methods("POST")
	r.HandleFunc("/login", LoginUser).Methods("POST")
	r.HandleFunc("/password-reset", ResetPassword).Methods("POST")
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req service.RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := service.RegisterUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message":"User registered successfully"})
}


func LoginUser(w http.ResponseWriter, r *http.Request) {
	var req service.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := service.LoginUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})

}


func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req service.PasswordResetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := service.ResetPassword(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password reset successfully"})
}