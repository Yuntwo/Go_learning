package main

import "github.com/gin-gonic/gin"

func basic() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// This method is specified to Response
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
