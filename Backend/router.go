package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func mainGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":"200",
		"message":"Welcome in Syntil API",
	})
}