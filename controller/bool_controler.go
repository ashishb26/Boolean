package controller

import (
	"net/http"

	"github.com/ashishb26/rzpbool/models"
	"github.com/gin-gonic/gin"
)

// AddBool controller adds a new boolean to the database
func (s *Server) AddBool(c *gin.Context) {
	var inputBool models.BoolRecord

	if err := c.ShouldBindJSON(&inputBool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	(&inputBool).AddNewRecord(s.DB)

	c.JSON(http.StatusOK, inputBool)
}

// GetBool controller extracts and returns a boolean record
func (s *Server) GetBool(c *gin.Context) {
	recordID := c.Param("id")

	record := &models.BoolRecord{}
	err := record.GetRecordByID(s.DB, recordID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, record)
}

// UpdateBool controller updates a boolean record given the id
func (s *Server) UpdateBool(c *gin.Context) {
	recordID := c.Param("id")

	record := &models.BoolRecord{}

	if err := s.Mutex.Lock(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error executing your request"})
		return
	}
	if err := record.GetRecordByID(s.DB, recordID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
	if err := c.ShouldBindJSON(record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, record)
	if err := record.UpdateRecord(s.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	if ok, err := s.Mutex.Unlock(); !ok || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Request executed but encountered unexpected issues"})
		return
	}
	c.JSON(http.StatusOK, record)
}

// DeleteBool deletes a boolean record from the database
func (s *Server) DeleteBool(c *gin.Context) {
	recordID := c.Param("id")

	if err := s.Mutex.Lock(); err != nil {
		if err := s.Mutex.Lock(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error executing your request"})
			return
		}
	}
	err := models.DeleteRecordByID(s.DB, recordID)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if ok, err := s.Mutex.Unlock(); !ok || err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Request executed but encountered unexpected error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Successfully deleted the record"})
}
