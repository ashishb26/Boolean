package controller

import (
	"net/http"

	"github.com/Boolean/dbConfig"
	"github.com/Boolean/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

// AddBool function is used to add a new boolean into the database
func AddBool(c *gin.Context) {
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
	var reqEntry models.BoolTable

	if err := dbConfig.DB.Where("id=?", c.Param("id")).First(&reqEntry).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, reqEntry)
}

// UpdateBool updates the database records given the id and info to be changed
func UpdateBool(c *gin.Context) {
	var record models.BoolTable

	if err := dbConfig.DB.Where("id=?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := c.BindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//dbConfig.DB.Save(&oldRecord)
	//var input models.InputBool
	// if err := c.BindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	// 	return
	// }
	//
	dbConfig.DB.Model(&record).Updates(record)

	c.JSON(http.StatusOK, record)
}

// DeleteBool deletes the record from database whose id
// matches the given id
func DeleteBool(c *gin.Context) {
	var record models.BoolTable

	if err := dbConfig.DB.Where("id=?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	dbConfig.DB.Delete(&record)

	c.JSON(http.StatusOK, gin.H{"status": "Succesfully deleted"})
}
