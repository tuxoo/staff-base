package service

import (
	"context"
	"github.com/tuxoo/smart-loader/staff-base/internal/model"
	"github.com/tuxoo/smart-loader/staff-base/internal/repository"
)

type IEmployeeService interface {
	AddEmployee(ctx context.Context, dto model.NewEmployeeDto) (model.Employee, error)
	GetEmployeeByName(ctx context.Context, name string) ([]model.Employee, error)
	GetEmployeeVacation(ctx context.Context, id int) (string, error)
	DeleteEmployee(ctx context.Context, id int) error
}

type Services struct {
	EmployeeService IEmployeeService
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		EmployeeService: NewEmployeeService(repositories.EmployeeRepository),
	}
}
