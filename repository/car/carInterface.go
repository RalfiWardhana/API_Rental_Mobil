package car

import "rental/entity"

type CarRepository interface {
	CreateCar(car entity.Car) error
	FindAllCar() ([]entity.Car_get, error)
	FindByIDCar(id string) (entity.Car_get, error)
	UpdateCar(id string, car entity.Car) (error, string)
	DeleteCar(id string, car entity.Car) (error, string)
}
