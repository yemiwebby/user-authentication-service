package repository

import (
	"errors"

	"github.com/yemiwebby/user-authentication-service/internal/model"
)


var users = map[string]*model.User{
	"gdg@example.com": {
        Email:    "gdg@example.com",
        Password: "hashed-securepassword",
		Name: "GDG Participant",
    },
}


func SaveUser(user *model.User) error {
	if _, exists := users[user.Email]; exists {
		return errors.New("user already exists")
	}
	users[user.Email] = user
	return nil
}



func FindUserByEmail(email string) (*model.User, error) {
	user, exists := users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func UpdateUser(user *model.User) error {
	if _, exists := users[user.Email]; !exists {
		return errors.New("user not found")
	}
	users[user.Email] = user
	return nil
}