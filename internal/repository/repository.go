package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tuxoo/smart-loader/staff-base/internal/model"
)

const (
	employeeTable = "employee"
)

type IEmployeeRepository interface {
	Save(ctx context.Context, employee model.Employee) (model.Employee, error)
}

type Repositories struct {
	EmployeeRepository IEmployeeRepository
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		EmployeeRepository: NewEmployeeRepository(db),
	}
}