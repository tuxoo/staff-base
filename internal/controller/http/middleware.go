package http

import (
	"errors"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	authorizationHeader = "Authorization"
	contentTypeHeader   = "Content-Type"
	jsonContent         = "application/json"
	xmlContent          = "application/xml"
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

	if contentType != jsonContent && contentType != xmlContent {
		newErrorResponse(c, http.StatusUnsupportedMediaType, errors.New("wrong content type"))
	}
}

func (h *Handler) userIdentity(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()

	if err := h.authenticator.Authentication(user, password, hasAuth); err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err)
	}
}
