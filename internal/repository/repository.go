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
	FindById(ctx context.Context, id int) (model.Employee, error)
	FindByName(ctx context.Context, name string) ([]model.Employee, error)
	DeleteById(ctx context.Context, id int) error
}

type Repositories struct {
	EmployeeRepository IEmployeeRepository
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		EmployeeRepository: NewEmployeeRepository(db),
	}
}
