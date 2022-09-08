package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"rental/domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user domain.User) error {
	query := `
        INSERT INTO user (username, email, password, id_user_type) 
        VALUES (?, ?, ?, ?)`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Id_user_type)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindAllUser() ([]domain.User_get, error) {
	query := `
	SELECT u.id, u.username, u.email, u.password, ut.user_type
	FROM user u
	JOIN user_type ut ON u.id_user_type = ut.id`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []domain.User_get

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User_get
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

func (r *repository) FindByIDUser(id string) (domain.User_get, error) {
	query := `
        SELECT u.id, u.username, u.email, u.password, ut.user_type
        FROM user u
		JOIN user_type ut ON u.id_user_type = ut.id
        WHERE u.id = ?`

	var user domain.User_get

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
		result := domain.User_get{}
		return result, err
	}

	return user, nil
}

func (r *repository) UpdateUser(id string, user domain.User) (error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE user set username=?, email=?, password=?, id_user_type=? WHERE id=?`
	result, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Id_user_type, id)
	if err != nil {
		return err, "Fail to update"
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err, "Fail to update"
	}

	fmt.Printf("Affected update : %d", rows)
	return nil, "Success to update id = " + id
}

func (r *repository) DeleteUser(id string, user domain.User) (error, string) {
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
	return nil, "Success to delete id = " + id
}

func (r *repository) DecryptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		}
		return domain.PrivateKey, nil
	})
	if err != nil {
		return map[string]interface{}{}, err
	}

	if !parsedToken.Valid {
		return map[string]interface{}{}, err
	}

	return parsedToken.Claims.(jwt.MapClaims), nil
}

func (r *repository) FindByEmailUser(email string) (domain.User_get, error) {
	query := `
        SELECT u.id, u.username, u.email, u.password, ut.user_type
        FROM user u
		JOIN user_type ut ON u.id_user_type = ut.id
        WHERE u.email = ?`

	var user domain.User_get

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.User_type,
	)

	if err != nil {
		result := domain.User_get{}
		return result, err
	}

	return user, nil
}

func (r *repository) FindByPasswordUser(password string) (domain.User_get, error) {
	query := `
        SELECT u.id, u.username, u.email, u.password, ut.user_type
        FROM user u
		JOIN user_type ut ON u.id_user_type = ut.id
        WHERE u.password = ?`

	var user domain.User_get

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, password).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.User_type,
	)

	if err != nil {
		result := domain.User_get{}
		return result, err
	}

	return user, nil
}
