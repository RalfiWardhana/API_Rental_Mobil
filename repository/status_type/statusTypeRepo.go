package statusType

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

func NewStatusTypeRepository(db *sql.DB) StatusTypeRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateStatusType(statusType domain.Status_type) error {
	query := `
        INSERT INTO status_type (status) 
        VALUES (?)`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, statusType.Status_type)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindAllStatusType() ([]domain.Status_type, error) {
	query := `
		SELECT id, status_type FROM status_type`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var Statuss_type []domain.Status_type

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Status_type domain.Status_type
		err := rows.Scan(
			&Status_type.Id,
			&Status_type.Status_type,
		)
		if err != nil {
			return nil, err
		}

		Statuss_type = append(Statuss_type, Status_type)
	}

	return Statuss_type, nil
}

func (r *repository) FindByIDStatusType(id string) (domain.Status_type, error) {
	query := `
        SELECT id , status_type
        FROM status_type
        WHERE id = ?`

	var Status_type domain.Status_type

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&Status_type.Id,
		&Status_type.Status_type,
	)

	if err != nil {
		result := domain.Status_type{}
		return result, err
	}

	return Status_type, nil
}

func (r *repository) UpdateStatusType(id string, status_type domain.Status_type) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE status_type set status_type=? WHERE id=?`
	result, err := r.db.ExecContext(ctx, query, status_type.Status_type, status_type.Id)
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

func (r *repository) DeleteStatusType(id string, status_type domain.Status_type) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE from status_type WHERE id=?`
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
