package car

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

func NewCarRepository(db *sql.DB) CarRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateCar(car entity.Car) error {
	query := `
        INSERT INTO car (car_name, cc, capacity, total, price, id_car_type) 
        VALUES (?, ?, ?, ?, ?)`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, car.Car_name, car.Cc, car.Capacity, car.Total, car.Price, car.Id_car_type)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindAllCar() ([]entity.Car_get, error) {
	query := `
	SELECT c.id, c.car_name, c.cc, c.capacity, c.total, c.price, ct.car_type
	FROM car c
	JOIN car_type ct ON c.id_Car_type = ct.id`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var Cars []entity.Car_get

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Car entity.Car_get
		err := rows.Scan(
			&Car.Id,
			&Car.Car_name,
			&Car.Cc,
			&Car.Capacity,
			&Car.Total,
			&Car.Price,
			&Car.Car_type,
		)
		if err != nil {
			return nil, err
		}

		Cars = append(Cars, Car)
	}

	return Cars, nil
}

func (r *repository) FindByIDCar(id string) (entity.Car_get, error) {
	query := `
	SELECT c.id, c.car_name, c.cc, c.capacity, c.total, c.price, ct.car_type
	FROM car c
	JOIN car_type ct ON c.id_Car_type = ct.id
    WHERE c.id = ?`

	var Car entity.Car_get

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&Car.Id,
		&Car.Car_name,
		&Car.Cc,
		&Car.Capacity,
		&Car.Total,
		&Car.Price,
		&Car.Car_type,
	)

	if err != nil {
		result := entity.Car_get{}
		return result, err
	}

	return Car, nil
}

func (r *repository) UpdateCar(id string, car entity.Car) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE car set car_name=?, cc=?, capacity=?, total=?, price = ?, id_car_type =? WHERE id=?`
	result, err := r.db.ExecContext(ctx, query, car.Car_name, car.Cc, car.Capacity, car.Total, car.Price, car.Id_car_type, car.Id)
	if err != nil {
		return err, "Fail to update"
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err, "Fail to update"
	}

	fmt.Printf("Affected update : %d", rows)
	return nil, "success to update id = " + string(rows)
}

func (r *repository) DeleteCar(id string, car entity.Car) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE from car WHERE id=?`
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
