package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type errorResponse struct {
	ErrorTime string `json:"errorTime"`
	Message   string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, err error) {
	logrus.Error(err)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		ErrorTime: time.Now().Format(timeFormat),
		Message:   err.Error(),
	})
}

func timeoutResponse(c *gin.Context) {
	logrus.Error("request timeout")
	c.JSON(http.StatusRequestTimeout, errorResponse{
		ErrorTime: time.Now().Format(timeFormat),
		Message:   "request timeout",
	})
}
