package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterUser(t *testing.T) {
    requestBody := `{"email": "test@example.com", "password": "securepassword"}`
    req, _ := http.NewRequest("POST", "/registration", bytes.NewBuffer([]byte(requestBody)))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()

    RegisterUser(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }

    expected := `{"message":"User registered successfully"}`
    if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
