package user

import "rental/domain"

type UserRepository interface {
	CreateUser(user domain.User) error
	FindAllUser() ([]domain.User_get, error)
	FindByIDUser(id string) (domain.User_get, error)
	UpdateUser(id string, user domain.User) (error, string)
	DeleteUser(id string, user domain.User) (error, string)
}
