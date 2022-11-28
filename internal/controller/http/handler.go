package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tuxoo/smart-loader/staff-base/internal/service"
	"github.com/tuxoo/smart-loader/staff-base/pkg/auth"
	"net/http"
	"time"
)

type Handler struct {
	employeeService service.IEmployeeService
	authenticator   auth.BasicAuth
}

func NewHandler(employeeService service.IEmployeeService, authenticator auth.BasicAuth) *Handler {
	return &Handler{
		employeeService: employeeService,
		authenticator:   authenticator,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		timeoutMiddleware(),
		cors.New(corsConfig),
	)

	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	api := router.Group("/api")
	{
		h.initEmployeeRoutes(api)
	}
}
