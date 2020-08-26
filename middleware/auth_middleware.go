package middleware

import (
	"net/http"

	"github.com/ashishb26/rzpbool/auth"
	"github.com/gin-gonic/gin"
)

// AuthUser middleware authnticates the user first before making any api calls
func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.IsValidToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
	}
}
