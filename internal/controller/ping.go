package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary Проверка работоспособности сервера
// @Tags health
// @Success 200 {object} map[string]string
// @Router / [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is up and running",
	})
}
