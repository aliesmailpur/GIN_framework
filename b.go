package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()
	r.GET("/name", getName)

	_ = r.Run()
}


func getName(c *gin.Context) {
	message := c.DefaultPostForm("message", "check your mail")
	name := c.DefaultPostForm("name", "ali")

	c.JSON(http.StatusOK, gin.H{
        "status": 200,
        "message": message,
        "name": name,
    })
}