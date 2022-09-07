package user

import "rental/domain"

type UserRepository interface {
	CreateUser(user domain.User) error
	FindAllUser() ([]domain.User, error)
	FindByIDUser(id string) (domain.User, error)
	UpdateUser(id string, user domain.User) (error, string)
	DeleteUser(id string, user domain.User) (error, string)
}
