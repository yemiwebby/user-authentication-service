package service

import (
	"errors"

	"github.com/yemiwebby/user-authentication-service/internal/model"
	"github.com/yemiwebby/user-authentication-service/internal/repository"
)

// Request payloads
type RegistrationRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type PasswordResetRequest struct {
	Email string `json:"email"`
	NewPassword string `json:"new_password"`
}

func RegisterUser(req RegistrationRequest) error {
	hashedPassword := hashPassword(req.Password)
	user := model.User{
		Email: req.Email,
		Password: hashedPassword,
	}

	return repository.SaveUser(user)
}

func LoginUser(req LoginRequest) (string, error) {
	user, err := repository.FiindUserByEmail(req.Email)
	if err != nil || user.Password != hashPassword(req.Password) {
		return "", errors.New("invalid email or password")
	}

	return "perfect-mocked-jwt-token-123456789", nil
}

func ResetPassword(req PasswordResetRequest) error {
	user, err := repository.FiindUserByEmail(req.Email)
	if err != nil {
		return errors.New("user not found")
	}

	user.Password = hashPassword(req.NewPassword)
	return repository.UpdateUser(user)
}


func hashPassword(password string) string {
	// for simplicity
	return "hashed-" + password + "!@#$%^&*()"
}