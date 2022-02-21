package main

import (
	"net/http"
	"shortener/v1/controllers"
	"shortener/v1/models"

	"github.com/gin-gonic/gin"
)

func ApiVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"versions": "v1"})
}

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/version", ApiVersion)
	r.GET("/api/shorteners", controllers.FindShortener)
	r.POST("/api/shorteners", controllers.CreateShortener)

	r.Run()
}
