package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"detail": "Welcome to TaskManager",
	})
}
