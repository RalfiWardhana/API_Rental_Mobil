package carType

import (
	"context"
	"database/sql"
	"fmt"
	"rental/entity"
	"time"
)

type repository struct {
	db *sql.DB
}

func NewCarTypeRepository(db *sql.DB) CarTypeRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateCarType(carType entity.Car_type) error {
	query := `
        INSERT INTO car_type (car_type) 
        VALUES (?)`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, carType.Car_type)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindAllCarType() ([]entity.Car_type, error) {
	query := `
		SELECT id, car_type FROM car_type`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var Cars_type []entity.Car_type

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Car_type entity.Car_type
		err := rows.Scan(
			&Car_type.Id,
			&Car_type.Car_type,
		)
		if err != nil {
			return nil, err
		}

		Cars_type = append(Cars_type, Car_type)
	}

	return Cars_type, nil
}

func (r *repository) FindByIDCarType(id string) (entity.Car_type, error) {
	query := `
        SELECT id , car_type
        FROM car_type
        WHERE id = ?`

	var Car_type entity.Car_type

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&Car_type.Id,
		&Car_type.Car_type,
	)

	if err != nil {
		result := entity.Car_type{}
		return result, err
	}

	return Car_type, nil
}

func (r *repository) UpdateCarType(id string, car_type entity.Car_type) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE Car_type set Car_type=? WHERE id=?`
	result, err := r.db.ExecContext(ctx, query, car_type.Car_type, car_type.Id)
	if err != nil {
		return err, "Fail to update"
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err, "Fail to update"
	}

	fmt.Printf("Affected update : %d", rows)
	return nil, "Success to update id = " + string(rows)
}

func (r *repository) DeleteCarType(id string, car_type entity.Car_type) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE from Car_type WHERE id=?`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err, "Fail to delete"
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err, "Fail to delete"
	}
	fmt.Printf("Affected delete : %d", rows)
	return nil, "Success to delete id = " + string(rows)
}
