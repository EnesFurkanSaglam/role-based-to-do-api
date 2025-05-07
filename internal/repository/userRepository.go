package repository

import "role-based-to-do-api/internal/model"

var users = []model.User{
	{Username: "enes", Password: "1234", Role: "basic"},
	{Username: "admin", Password: "admin", Role: "admin"},
}

func FindUser(username, password string) (*model.User, bool) {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user, true
		}
	}
	return nil, false
}
