package user

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

func NewUserRepository(db *sql.DB) UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user entity.User) error {
	query := `
        INSERT INTO user (username, email, password, id_type_user) 
        VALUES (?, ?, ?, ?)`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Id_user_type)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindAllUser() ([]entity.User, error) {
	query := `
	SELECT u.id, u.username, u.email, u.password, ut.user_type
	FROM user u
	JOIN user_type ut ON u.id_type_user = ut.id`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []entity.User

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.User_type,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *repository) FindByIDUser(id string) (entity.User, error) {
	query := `
        SELECT u.id, u.username, u.email, u.password, ut.user_type
        FROM user u
		JOIN user_type ut ON u.id_type_user = ut.id
        WHERE id = $1`

	var user entity.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.User_type,
	)

	if err != nil {
		result := entity.User{}
		return result, err
	}

	return user, nil
}

func (r *repository) UpdateUser(id string, user entity.User) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE user set username=?, email=?, password=?, id_user_type=? WHERE id=?`
	result, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Id_user_type, user.Id)
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

func (r *repository) DeleteUser(id string, user entity.User) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE from user WHERE id=?`
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
