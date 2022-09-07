package userType

import "rental/domain"

type UserTypeRepository interface {
	CreateUserType(user_type domain.User_type) error
	FindAllUserType() ([]domain.User_type, error)
	FindByIDUserType(id string) (domain.User_type, error)
	UpdateUserType(id string, user_type domain.User_type) (error, string)
	DeleteUserType(id string, user_type domain.User_type) (error, string)
}
