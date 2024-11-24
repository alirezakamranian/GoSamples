package main 

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/getdetails", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Ip":c.ClientIP(),"Name":"kami","Url":c.Request.URL.Path})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run("192.168.1.20:8080")
}
