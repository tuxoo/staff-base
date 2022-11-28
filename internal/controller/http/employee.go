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
		load.GET("/", h.getEmployee)
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
	id := getIdFromGinContext(c)

	if err := h.employeeService.DeleteEmployee(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
	}

	c.Status(http.StatusOK)
}

func (h *Handler) getEmployee(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		newErrorResponse(c, http.StatusBadRequest, errors.New("empty field [name]"))
	}

	employees, err := h.employeeService.GetEmployeeByName(c.Request.Context(), name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, employees)
}

func (h *Handler) getEmployeeVacation(c *gin.Context) {
	id := getIdFromGinContext(c)

	vacation, err := h.employeeService.GetEmployeeVacation(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, map[string]string{
		"vacationPeriod": vacation,
	})
}

func getIdFromGinContext(c *gin.Context) int {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errors.New("incorrect ID parameter"))
	}

	return idParam
}
