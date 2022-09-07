package transaction

import (
	"context"
	"database/sql"
	"fmt"
	"rental/domain"
	"time"
)

type repository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateTransaction(transaction domain.Transaction) error {
	query := `
        INSERT INTO transaction (id_car, id_user, total_price, duration, status) 
        VALUES (?, ?, ?, ?, ?)`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, transaction.Id_car, transaction.Id_user, transaction.Total_price, transaction.Duration, transaction.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindAllTransaction() ([]domain.Transaction_get, error) {
	query := `
	SELECT t.id, u.username,u.email, c.car_name, c.cc, t.total_price, t.duration, t.status  FROM transaction t
	JOIN car c ON c.id = t.id_car
	JOIN user u ON  u.id = t.id_user
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var Transactions []domain.Transaction_get

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Transaction domain.Transaction_get
		err := rows.Scan(
			&Transaction.Id,
			&Transaction.Username,
			&Transaction.Email,
			&Transaction.Car_name,
			&Transaction.Cc,
			&Transaction.Total_price,
			&Transaction.Duration,
			&Transaction.Status,
		)
		if err != nil {
			return nil, err
		}

		Transactions = append(Transactions, Transaction)
	}

	return Transactions, nil
}

func (r *repository) FindByIDTransaction(id string) (domain.Transaction_get, error) {
	query := `
	SELECT t.id, u.username,u.email, c.car_name, c.cc, t.total_price, t.duration, t.status  FROM transaction t
	JOIN car c ON c.id = t.id_car
	JOIN user u ON  u.id = t.id_user
    WHERE t.id = ?`

	var Transaction domain.Transaction_get

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&Transaction.Id,
		&Transaction.Username,
		&Transaction.Email,
		&Transaction.Car_name,
		&Transaction.Cc,
		&Transaction.Total_price,
		&Transaction.Duration,
		&Transaction.Status,
	)

	if err != nil {
		result := domain.Transaction_get{}
		return result, err
	}

	return Transaction, nil
}

func (r *repository) UpdateTransaction(id string, transaction domain.Transaction) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE transaction set id_car=?, id_user=?, total_price=?, duration=?, status =? WHERE id=?`
	result, err := r.db.ExecContext(ctx, query, transaction.Id_car, transaction.Id_user, transaction.Total_price, transaction.Duration, transaction.Status, transaction.Id)
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

func (r *repository) DeleteTransaction(id string, transaction domain.Transaction) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE from transaction WHERE id=?`
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
