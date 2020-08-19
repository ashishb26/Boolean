package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Boolean/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("idghskdjfbsjduew")

// createToken creates and returns a token using the user credentials supplied
func createToken(userCred models.Credentials, expirationTime time.Time) (string, error) {

	claims := models.Claims{
		Username: userCred.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}

// authUser authenticates the user by verifying the legality of the
// token stored in the form of cookies (if already logged in).
// Otherwise it restricts user access
func authUser(c *gin.Context) bool {
	tokenString, err := c.Cookie("authToken")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, "User not authorised")
			return false
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return false
	}

	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %s", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return false
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return false
	}

	_, ok := token.Claims.(*models.Claims)

	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, "User not authorized")
		return false
	}

	return true
}
