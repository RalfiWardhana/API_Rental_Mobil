package userType

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

func NewUserTypeRepository(db *sql.DB) UserTypeRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUserType(userType entity.User_type) error {
	query := `
        INSERT INTO user_type (user_type) 
        VALUES (?)`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, userType.User_type)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindAllUserType() ([]entity.User_type, error) {
	query := `
		SELECT id, user_type FROM user_type`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users_type []entity.User_type

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user_type entity.User_type
		err := rows.Scan(
			&user_type.Id,
			&user_type.User_type,
		)
		if err != nil {
			return nil, err
		}

		users_type = append(users_type, user_type)
	}

	return users_type, nil
}

func (r *repository) FindByIDUserType(id string) (entity.User_type, error) {
	query := `
        SELECT id , user_type
        FROM user_type
        WHERE id = ?`

	var user_type entity.User_type

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user_type.Id,
		&user_type.User_type,
	)

	if err != nil {
		result := entity.User_type{}
		return result, err
	}

	return user_type, nil
}

func (r *repository) UpdateUserType(id string, user_type entity.User_type) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE user_type set user_type=? WHERE id=?`
	result, err := r.db.ExecContext(ctx, query, user_type.User_type, user_type.Id)
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

func (r *repository) DeleteUserType(id string, user_type entity.User_type) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE from user_type WHERE id=?`
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
