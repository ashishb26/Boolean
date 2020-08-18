package controller

import (
	"net/http"
	"time"

	"github.com/Boolean/dbConfig"
	"github.com/Boolean/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

// AddBool function is used to add a new boolean into the database
func AddBool(c *gin.Context) {

	if !authUser(c) {
		return
	}

	var input models.BoolTable
	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = xid.New().String()

	dbConfig.DB.Create(&input)

	c.JSON(http.StatusOK, input)
}

// GetBool is used to retrieve a boolean from the database
// provided the id of the boolean database entry
func GetBool(c *gin.Context) {

	if !authUser(c) {
		return
	}

	var reqEntry models.BoolTable

	if err := dbConfig.DB.Where("id=?", c.Param("id")).First(&reqEntry).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reqEntry)
}

// UpdateBool updates the database records given the id and info to be changed
func UpdateBool(c *gin.Context) {

	if !authUser(c) {
		return
	}

	var record models.BoolTable

	if err := dbConfig.DB.Where("id=?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbConfig.DB.Model(&record).Updates(record)

	c.JSON(http.StatusOK, record)
}

// DeleteBool deletes the record from database whose id
// matches the given id
func DeleteBool(c *gin.Context) {

	if !authUser(c) {
		return
	}

	var record models.BoolTable

	if err := dbConfig.DB.Where("id=?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbConfig.DB.Delete(&record)

	c.JSON(http.StatusOK, gin.H{"status": "Succesfully deleted"})
}

// Login authenticates the user and sets a token as a cookie
func Login(c *gin.Context) {
	var userCred models.Credentials

	if err := c.ShouldBindJSON(&userCred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var authCred models.Credentials
	if err := dbConfig.DB.Where("username=?", userCred.Username).First(&authCred).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if authCred.Password != userCred.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect login credentials"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	tokenString, err := createToken(authCred, expirationTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("authToken", tokenString, 300, "/", "localhost", false, true)
	c.JSON(http.StatusOK, "Logged in succesfully")
	//	c.Request.AddCookie(cookie)
	//c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)

}
