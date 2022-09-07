package user

import "rental/entity"

type UserRepository interface {
	CreateUser(user entity.User) error
	FindAllUser() ([]entity.User, error)
	FindByIDUser(id string) (entity.User, error)
	UpdateUser(id string, user entity.User) (error, string)
	DeleteUser(id string, user entity.User) (error, string)
}
