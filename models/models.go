package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// BoolTable is a struct that is used to create and
// bind to the database table that stores the boolean
type BoolTable struct {
	ID        string `json:"id" gorm:"primary key"`
	Value     bool   `json:"value"`
	Label     string `json:"label"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Credentials struct stores the user's username and password.
// It is used to bind to the user login info and also to the
// user credentials stored in the database
type Credentials struct {
	ID       string
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims struct represents the claims used to build the tokens
type Claims struct {
	Username string
	jwt.StandardClaims
}
