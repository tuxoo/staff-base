package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tuxoo/smart-loader/staff-base/internal/model"
	"net/http"
	"strconv"
)

func (h *Handler) initEmployeeRoutes(api *gin.RouterGroup) {
	load := api.Group("/employee")
	{
		load.POST("/", h.addEmployee)
		load.DELETE("/:id", h.deleteEmployee)
		load.GET("/:id", h.getEmployee)
		load.GET("/:id/vacation", h.getEmployeeVacation)
	}
}

func (h *Handler) addEmployee(c *gin.Context) {
	var newEmployee model.NewEmployeeDto

	if err := c.ShouldBindJSON(&newEmployee); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	employee, err := h.employeeService.AddEmployee(c.Request.Context(), newEmployee)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, employee)
}

func (h *Handler) deleteEmployee(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errors.New("incorrect ID parameter"))
	}

	if err := h.employeeService.DeleteEmployee(c.Request.Context(), idParam); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errors.New("incorrect ID parameter"))
	}

	c.Status(http.StatusOK)
}

func (h *Handler) getEmployee(c *gin.Context) {

}

func (h *Handler) getEmployeeVacation(c *gin.Context) {

}
