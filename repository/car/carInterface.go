package car

import "rental/domain"

type CarRepository interface {
	CreateCar(car domain.Car) error
	FindAllCar() ([]domain.Car_get, error)
	FindByIDCar(id string) (domain.Car_get, error)
	UpdateCar(id string, car domain.Car) (error, string)
	DeleteCar(id string, car domain.Car) (error, string)
}
