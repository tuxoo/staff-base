package service

import (
	"context"
	"fmt"
	"github.com/tuxoo/smart-loader/staff-base/internal/model"
	"github.com/tuxoo/smart-loader/staff-base/internal/repository"
	"math"
	"time"
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

func (s *EmployeeService) GetEmployeeByName(ctx context.Context, name string) ([]model.Employee, error) {
	return s.repository.FindByName(ctx, name)
}

func (s *EmployeeService) GetEmployeeVacation(ctx context.Context, id int) (string, error) {
	employee, err := s.repository.FindById(ctx, id)
	if err != nil {
		return "", err
	}

	vacationDays := computeVacationDaysByPeriod(employee.RegisteredAt, time.Now())

	vacation := fmt.Sprintf("%d days", vacationDays)

	return vacation, nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	return s.repository.DeleteById(ctx, id)
}

func computeVacationDaysByPeriod(start, stop time.Time) (days int) {
	var months int
	m := start.Month()
	for start.Before(stop) {
		start = start.Add(time.Hour * 24)
		m2 := start.Month()
		if m2 != m {
			months++
		}
		m = m2
	}

	if stop.Day() >= 15 {
		months++
	}

	days = int(math.Ceil(float64(months) * 2.33))

	return
}
