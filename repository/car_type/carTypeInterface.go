package carType

import "rental/domain"

type CarTypeRepository interface {
	CreateCarType(car_type domain.Car_type) error
	FindAllCarType() ([]domain.Car_type, error)
	FindByIDCarType(id string) (domain.Car_type, error)
	UpdateCarType(id string, car_type domain.Car_type) (error, string)
	DeleteCarType(id string, car_type domain.Car_type) (error, string)
}
