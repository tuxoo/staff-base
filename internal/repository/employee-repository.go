package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tuxoo/smart-loader/staff-base/internal/model"
)

type EmployeeRepository struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (r *EmployeeRepository) Save(ctx context.Context, employee model.Employee) (model.Employee, error) {
	tx, err := r.db.Begin(ctx)
	defer func(tx pgx.Tx, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}(tx, ctx)

	query := fmt.Sprintf(`
	INSERT INTO %s (full_name, phone, gender, age, email, address)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, full_name, phone, gender, age, email, address, registered_at
	`, employeeTable)

	row := tx.QueryRow(ctx, query, employee.FullName, employee.Phone, employee.Gender, employee.Age, employee.Email, employee.Address)

	if err = row.Scan(
		&employee.Id,
		&employee.FullName,
		&employee.Phone,
		&employee.Gender,
		&employee.Age,
		&employee.Email,
		&employee.Address,
		&employee.RegisteredAt); err != nil {
		return employee, err
	}

	if err = tx.Commit(ctx); err != nil {
		return employee, err
	}

	return employee, nil
}
