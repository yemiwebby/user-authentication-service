package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/yemiwebby/user-authentication-service/internal/service"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/registration", RegisterUser).Methods("POST")
	r.HandleFunc("/password-reset", ResetPassword).Methods("POST")
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

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

	emailServiceURL := os.Getenv("EMAIL_SERVICE_URL") 
	if emailServiceURL == "" {
		emailServiceURL = "http://localhost:8081" 
	}

	go func() {
		emailReq := map[string]string{
			"recipient": req.Email,
			"subject":   "Welcome to Our Platform",
			"body":      "Thank you for registering!",
		}
		jsonData, _ := json.Marshal(emailReq)
		resp, err := http.Post(emailServiceURL+"/send", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("Failed to send email: %v\n", err)
			return
		}
		defer resp.Body.Close()
		log.Printf("Email sent with status: %d\n", resp.StatusCode)
	}()

	log.Printf("Registration processing completed in %v\n", time.Since(start))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}


// func RegisterUser(w http.ResponseWriter, r *http.Request) {
// 	start := time.Now()

// 	var req service.RegistrationRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	err := service.RegisterUser(req)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	emailServiceURL := os.Getenv("EMAIL_SERVICE_URL")
// 	if emailServiceURL == "" {
// 		emailServiceURL = "http://localhost:8081"
// 	}

// 	// Synchronous email sending (without concurrency)
// 	emailReq := map[string]string{
// 		"recipient": req.Email,
// 		"subject":   "Welcome to Our Platform",
// 		"body":      "Thank you for registering!",
// 	}
// 	jsonData, _ := json.Marshal(emailReq)
// 	resp, err := http.Post(emailServiceURL+"/send", "application/json", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		log.Printf("Failed to send email: %v\n", err)
// 		http.Error(w, "Failed to send email", http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	log.Printf("Email sent with status: %d\n", resp.StatusCode)

// 	log.Printf("Registration processing completed in %v\n", time.Since(start))
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
// }






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

	emailServiceURL := os.Getenv("EMAIL_SERVICE_URL") 
	if emailServiceURL == "" {
		emailServiceURL = "http://localhost:8081" 
	}

	// Send password reset email to the Email Notification Service
	go func() {
		emailReq := map[string]string{
			"recipient": req.Email,
			"subject":   "Your Password Has Been Reset",
			"body":      "You have successfully reset your password.",
		}
		jsonData, _ := json.Marshal(emailReq)
		resp, err := http.Post(emailServiceURL+"/send", "application/json", bytes.NewBuffer(jsonData))

		if err != nil {
			log.Printf("Failed to send email: %v\n", err)
			return
		}
		defer resp.Body.Close()
		log.Printf("Password reset email sent with status: %d\n", resp.StatusCode)
	}()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password reset successfully"})
}