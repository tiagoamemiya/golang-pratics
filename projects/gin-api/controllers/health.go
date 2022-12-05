package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Alive!",
	})
}

func TestParamsUrl(c *gin.Context) {
	param := c.Params.ByName("param")
	c.JSON(http.StatusOK, gin.H{
		"param value": param,
	})
}
