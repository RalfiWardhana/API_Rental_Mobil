package carType

import "rental/entity"

type CarTypeRepository interface {
	CreateCarType(car_type entity.Car_type) error
	FindAllCarType() ([]entity.Car_type, error)
	FindByIDCarType(id string) (entity.Car_type, error)
	UpdateCarType(id string, car_type entity.Car_type) (error, string)
	DeleteCarType(id string, car_type entity.Car_type) (error, string)
}
