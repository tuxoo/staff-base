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

	if err = composeEmployee(&employee, row); err != nil {
		return employee, err
	}

	if err = tx.Commit(ctx); err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *EmployeeRepository) FindById(ctx context.Context, id int) (model.Employee, error) {
	var employee model.Employee

	query := fmt.Sprintf(`
	SELECT id, full_name, phone, gender, age, email, address, registered_at FROM %s WHERE id=$1
	`, employeeTable)

	row := r.db.QueryRow(ctx, query, id)
	if err := composeEmployee(&employee, row); err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *EmployeeRepository) FindByName(ctx context.Context, name string) ([]model.Employee, error) {
	var employees []model.Employee
	query := fmt.Sprintf(`
	SELECT id, full_name, phone, gender, age, email, address, registered_at FROM %s WHERE full_name LIKE '%%' || $1 || '%%'
	`, employeeTable)

	rows, err := r.db.Query(ctx, query, name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employee model.Employee
		if err := composeEmployee(&employee, rows); err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *EmployeeRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", employeeTable)

	_, err := r.db.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}

func composeEmployee(employee *model.Employee, row pgx.Row) error {
	if err := row.Scan(
		&employee.Id,
		&employee.FullName,
		&employee.Phone,
		&employee.Gender,
		&employee.Age,
		&employee.Email,
		&employee.Address,
		&employee.RegisteredAt); err != nil {
		return err
	}

	return nil
}
