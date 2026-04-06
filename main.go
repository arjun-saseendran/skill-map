package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSONP(200, gin.H{"msg": "Welcome to Go Web Development"})
	})

	if err := r.Run(); err != nil {
		log.Fatal("Error starting server!")
	}
}
