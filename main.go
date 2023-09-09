package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.GET("/", homeHandler)
	r.GET("/user", getUser)
	r.Run()
}

func main() {
	RunServer()
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Home",
	})
}

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"User": "Ian",
	})
}
