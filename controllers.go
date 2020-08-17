package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func addBool(c *gin.Context) {
	var inp InputBool
	err := c.BindJSON(&inp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newEntry := BoolTable{
		ID:        xid.New().String(),
		Value:     inp.Value,
		Label:     inp.Label,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.Create(&newEntry)

	c.JSON(http.StatusOK, newEntry)
}
