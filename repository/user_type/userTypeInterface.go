package userType

import "rental/entity"

type UserTypeRepository interface {
	CreateUserType(user_type entity.User_type) error
	FindAllUserType() ([]entity.User_type, error)
	FindByIDUserType(id string) (entity.User_type, error)
	UpdateUserType(id string, user_type entity.User_type) (error, string)
	DeleteUserType(id string, user_type entity.User_type) (error, string)
}
