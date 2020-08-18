package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// BoolTable is a struct that binds to the db table
type BoolTable struct {
	ID        string `json:"id" gorm:"primary key"`
	Value     bool   `json:"value"`
	Label     string `json:"label"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Credentials struct is used to bind the JSON input
// of a user's username and password
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
