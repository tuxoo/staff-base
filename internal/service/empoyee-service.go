package service

import (
	"context"
	"github.com/tuxoo/smart-loader/staff-base/internal/model"
	"github.com/tuxoo/smart-loader/staff-base/internal/repository"
)

type EmployeeService struct {
	repository repository.IEmployeeRepository
}

func NewEmployeeService(repository repository.IEmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repository: repository,
	}
}

func (s *EmployeeService) AddEmployee(ctx context.Context, dto model.NewEmployeeDto) (model.Employee, error) {
	employee := model.Employee{
		FullName: dto.FullName,
		Phone:    dto.Phone,
		Gender:   dto.Gender,
		Age:      dto.Age,
		Email:    dto.Email,
		Address:  dto.Address,
	}

	return s.repository.Save(ctx, employee)
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	return s.repository.DeleteById(ctx, id)
}
