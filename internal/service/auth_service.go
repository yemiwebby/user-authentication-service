package service

import (
	"github.com/yemiwebby/user-authentication-service/internal/model"
	"github.com/yemiwebby/user-authentication-service/internal/repository"
)
type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
}

type PasswordResetRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}


func RegisterUser(req RegistrationRequest) error {
	user := &model.User{
		Email: req.Email,
		Password: hashPassword(req.Password),
		Name: req.Name,
	}
	return repository.SaveUser(user)
}

func ResetPassword(req PasswordResetRequest) error {
	user, err := repository.FindUserByEmail(req.Email)
	if err != nil {
		return err
	}
	user.Password = hashPassword(req.NewPassword)
	return repository.UpdateUser(user)
}


func hashPassword(password string) string {
	return "hashed-" + password + "!@@#$$%^&()"
}
