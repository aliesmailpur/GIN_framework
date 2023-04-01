package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.GET("/name", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status" : "ok",
			"message" : "hello world!",
		})
		
	})

	_ = r.Run(":5000")
}