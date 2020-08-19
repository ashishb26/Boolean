package controller

import (
	"net/http"
	"time"

	"github.com/Boolean/dbConfig"
	"github.com/Boolean/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

// AddBool function is used to add a new boolean into the database
// input JSON format :
// {
//   "value": // default false
//	 "label": //optional
// }
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
// provided the id of the database record
func GetBool(c *gin.Context) {

	if !authUser(c) {
		return
	}

	dbConfig.Mu.RLock()

	var reqEntry models.BoolTable

	if err := dbConfig.DB.Where("id=?", c.Param("id")).First(&reqEntry).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		dbConfig.Mu.RUnlock()
		return
	}

	dbConfig.Mu.RUnlock()

	c.JSON(http.StatusOK, reqEntry)
}

// UpdateBool updates the database records given the id and data to be changed
// input JSON format :
//{
// "value": //optional
// "label": //optional
//}
func UpdateBool(c *gin.Context) {

	if !authUser(c) {
		return
	}

	dbConfig.Mu.Lock()

	var record models.BoolTable

	if err := dbConfig.DB.Where("id=?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		dbConfig.Mu.Unlock()
		return
	}

	if err := c.BindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		dbConfig.Mu.Unlock()
		return
	}

	dbConfig.DB.Model(&record).Updates(record)

	dbConfig.Mu.Unlock()

	c.JSON(http.StatusOK, record)
}

// DeleteBool deletes the record from database whose id
// matches the given id
func DeleteBool(c *gin.Context) {

	if !authUser(c) {
		return
	}

	dbConfig.Mu.Lock()

	var record models.BoolTable

	if err := dbConfig.DB.Where("id=?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		dbConfig.Mu.Unlock()
		return
	}

	dbConfig.DB.Delete(&record)

	dbConfig.Mu.Unlock()

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

}

// AddUser allows a logged in user to add another user
func AddUser(c *gin.Context) {
	if !authUser(c) {
		return
	}

	var cred models.Credentials
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var record models.Credentials
	err := dbConfig.DB.Where("username=?", cred.Username).First(&record).Error
	if gorm.IsRecordNotFoundError(err) {
		cred.ID = xid.New().String()
		dbConfig.DB.Create(&cred)

		c.JSON(http.StatusOK, "User added successfully")
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})

}
