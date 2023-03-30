package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.GET("/name", getName)
    _ = r.Run()
}


func getName(c *gin.Context) {
    firstname := c.DefaultQuery("firstname", "ali")
    lastName := c.DefaultQuery("lastname", "esmailpur")
    c.String(http.StatusOK, "hello %s %s", firstname, lastName)
}