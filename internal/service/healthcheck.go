package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheckHandler(c *gin.Context) {
	c.Status(http.StatusOK)
}
