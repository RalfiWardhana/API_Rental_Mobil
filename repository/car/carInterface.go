package car

import "rental/entity"

type CarRepository interface {
	CreateCar(car entity.Car) error
	FindAllCar() ([]entity.Car, error)
	FindByIDCar(id string) (entity.Car, error)
	UpdateCar(id string, car entity.Car) (error, string)
	DeleteCar(id string, car entity.Car) (error, string)
}
