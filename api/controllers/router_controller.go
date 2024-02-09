package controllers

import (
	"api-auth/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublicRoute(c *gin.Context) {
	writer := c.Writer
	utils.ToJson(writer, "Public route", http.StatusOK)
}
