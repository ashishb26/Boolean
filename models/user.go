package models

import (
	"errors"
)

// User struct is used to accept user credentials and validate the user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// isValidInput checks if the user has supplied valid login format
func (user *User) isValidFormat() error {
	if user.Username == "" {
		return errors.New("Username required")
	}
	if user.Password == "" {
		return errors.New("Password required")
	}
	return nil
}

// IsValidUser checks whether the user is valid or not
func (user *User) IsValidUser() error {
	if err := user.isValidFormat(); err != nil {
		return err
	}

	if user.Username == "root" && user.Password == "password" {
		return nil
	}
	return errors.New("Incorrect username or password")
}
