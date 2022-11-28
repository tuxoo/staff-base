package service

import (
	"context"
	"github.com/tuxoo/smart-loader/staff-base/internal/model"
	"github.com/tuxoo/smart-loader/staff-base/internal/repository"
)

type IEmployeeService interface {
	AddEmployee(ctx context.Context, dto model.NewEmployeeDto) (model.Employee, error)
}

type Services struct {
	EmployeeService IEmployeeService
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		EmployeeService: NewEmployeeService(repositories.EmployeeRepository),
	}
}
