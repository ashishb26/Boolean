package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ashishb26/rzpbool/config"
	"github.com/ashishb26/rzpbool/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type claims struct {
	Username string
	jwt.StandardClaims
}

// GetToken takes a user struct as input, creates a token using
// the username, expiration time and returns a token string
func GetToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey, err := config.GetSecretKey()
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

// extractToken extracts and returns the token from the authorization header
func extractToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", errors.New("Authorization token not found")
}

// IsValidToken verifies the authenticity of the authorization token
// sent in the api request
func IsValidToken(c *gin.Context) error {
	tokenString, err := extractToken(c)
	if err != nil {
		return err
	}
	return ValidateToken(tokenString)
}

// ValidateToken takes a token string as input and validates the token
func ValidateToken(tokenString string) error {
	userClaims := &claims{}

	secretKey, err := config.GetSecretKey()

	if err != nil {
		return err
	}
	token, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %s", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}
	_, ok := token.Claims.(*claims)

	if !token.Valid || !ok {
		return err
	}
	return nil
}
