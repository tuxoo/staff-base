package http

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	contentType = "Content-Type"
	jsonContent = "application/json"
	xmlContent  = "application/xml"
)

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(3*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}
