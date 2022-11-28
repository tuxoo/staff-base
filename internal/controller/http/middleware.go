package http

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	contentTypeHeader = "Content-Type"
	jsonContent       = "application/json"
	xmlContent        = "application/xml"
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

func (h *Handler) contentIdentity(c *gin.Context) {
	contentType := c.GetHeader(contentTypeHeader)

	fmt.Println(contentType != jsonContent)
	fmt.Println(contentType != xmlContent)

	if contentType != jsonContent && contentType != xmlContent {
		newErrorResponse(c, http.StatusUnsupportedMediaType, errors.New("wrong content type"))
	}
}
