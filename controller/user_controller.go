package controller

import (
	"net/http"

	"github.com/ashishb26/rzpbool/auth"
	"github.com/ashishb26/rzpbool/models"
	"github.com/gin-gonic/gin"
)

// UserLogin is used to log a user in
func (s *Server) UserLogin(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.IsValidUser(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		return
	}

	token, err := auth.GetToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
