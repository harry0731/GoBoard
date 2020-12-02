// models/user.go

package models

import (
	"errors"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var UserList = []User{
	User{Username: "user1", Password: "pass1"},
	User{Username: "user2", Password: "pass2"},
	User{Username: "user3", Password: "pass3"},
}

func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !IsUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}

	u := User{Username: username, Password: password}

	UserList = append(UserList, u)

	return &u, nil
}

func IsUsernameAvailable(username string) bool {
	for _, u := range UserList {
		if u.Username == username {
			return false
		}
	}
	return true
}
