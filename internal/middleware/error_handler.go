package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tpetrychyn/wow-hardcore-recorder/internal/service"
)

func ErrorHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Next()

	for _, err := range c.Errors {
		// just return the first error to the user
		c.JSON(-1, gin.H{service.ErrorResponseKey: err.Error()})
		return
	}
}
