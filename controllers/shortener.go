package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"shortener/v1/models"
	"shortener/v1/wire"
	"time"

	"github.com/gin-gonic/gin"
)

func FindShortener(c *gin.Context) {
	var shortener []models.Shortener
	models.DB.Find(&shortener)

	c.JSON(http.StatusOK, gin.H{"shortener": shortener})
}

func EncryptURL(from string) string {
	data := []byte(from)
	hash := md5.Sum(data)

	return hex.EncodeToString(hash[:])
}

func CreateShortener(c *gin.Context) {
	var input wire.CreateShortener

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var to string = EncryptURL(input.From)

	shortener := models.Shortener{
		From:      input.From,
		To:        to,
		TTL:       time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	models.DB.Create(&shortener)

	c.JSON(http.StatusCreated, gin.H{"shortener": shortener})
}
