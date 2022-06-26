package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func test1Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test1 Gin",
		"344":     "2333",
	})
}

func test2Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test2 Gin",
	})
}
