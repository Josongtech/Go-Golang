package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsCreate() {
	func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ongosh",
		})
	}
}
