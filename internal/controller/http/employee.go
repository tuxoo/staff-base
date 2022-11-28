package http

import "github.com/gin-gonic/gin"

func (h *Handler) initEmployeeRoutes(api *gin.RouterGroup) {
	load := api.Group("/employee")
	{
		load.POST("/", h.addEmployee)
		load.DELETE("/id", h.deleteEmployee)
		load.GET("/id", h.getEmployee)
		load.GET("/id/vacation", h.getEmployeeVacation)
	}
}

func (h *Handler) addEmployee(c *gin.Context) {

}

func (h *Handler) deleteEmployee(c *gin.Context) {

}

func (h *Handler) getEmployee(c *gin.Context) {

}

func (h *Handler) getEmployeeVacation(c *gin.Context) {

}
