package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tuxoo/smart-loader/staff-base/internal/config"
	"github.com/tuxoo/smart-loader/staff-base/internal/model"
	"net/http"
	"strconv"
)

func (h *Handler) initEmployeeRoutes(api *gin.RouterGroup, cfg config.AdminConfig) {
	employees := api.Group("/employee", gin.BasicAuth(gin.Accounts{
		cfg.Login: cfg.Password,
	}))
	{
		checkContentType := employees.Group("/", h.contentIdentity)
		{
			checkContentType.POST("/", h.addEmployee)
		}

		employees.DELETE("/:id", h.deleteEmployee)
		employees.GET("/", h.getEmployees)
		employees.GET("/:id/vacation", h.getEmployeeVacation)
	}
}

func (h *Handler) addEmployee(c *gin.Context) {
	var newEmployee model.NewEmployeeDto

	if err := parseNewEmployee(c, &newEmployee); err != nil {
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
	if id == 0 {
		return
	}

	if err := h.employeeService.DeleteEmployee(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) getEmployees(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		newErrorResponse(c, http.StatusBadRequest, errors.New("empty field [name]"))
		return
	}

	employees, err := h.employeeService.GetEmployeeByName(c.Request.Context(), name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, employees)
}

func (h *Handler) getEmployeeVacation(c *gin.Context) {
	id := getIdFromGinContext(c)

	vacation, err := h.employeeService.GetEmployeeVacation(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, errors.New(fmt.Sprintf("employee not found in system by id [%d]", id)))
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"vacationPeriod": vacation,
	})
}

func parseNewEmployee(c *gin.Context, newEmployee *model.NewEmployeeDto) (err error) {
	switch c.GetHeader(contentTypeHeader) {
	case jsonContent:
		err = c.ShouldBindJSON(&newEmployee)
	case xmlContent:
		err = c.ShouldBindXML(&newEmployee)
	default:
		err = errors.New("incorrect content type")
	}

	return
}

func getIdFromGinContext(c *gin.Context) (id int) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errors.New("incorrect ID parameter"))
	}

	return
}
