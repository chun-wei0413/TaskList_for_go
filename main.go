package main

import (
	"github.com/gin-gonic/gin"
)

// go run .
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.Run("localhost:8080")
}
