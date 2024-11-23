package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/foo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"response": "bar",
		})
	})
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
