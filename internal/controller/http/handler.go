package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tuxoo/smart-loader/staff-base/internal/config"
	"github.com/tuxoo/smart-loader/staff-base/internal/service"
	"net/http"
	"time"
)

type Handler struct {
	employeeService service.IEmployeeService
}

func NewHandler(employeeService service.IEmployeeService) *Handler {
	return &Handler{
		employeeService: employeeService,
	}
}

func (h *Handler) Init(cfg config.HTTPConfig) *gin.Engine {
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
