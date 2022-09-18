package car

import (
	"rental/domain"
	"rental/repository/car"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	cr car.CarRepository
}

func NewCarController(cr car.CarRepository) CarController {
	return &Controller{
		cr: cr,
	}
}

func (cr *Controller) CreateCar(c *gin.Context) {

	var car domain.Car

	if err := c.ShouldBind(&car); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	if car.Car_name == "" {
		c.JSON(400, map[string]string{
			"message": "Car name required",
		})
		return
	}

	if car.Cc == 0 {
		c.JSON(400, map[string]string{
			"message": "Cc required",
		})
		return
	}

	if car.Capacity == 0 {
		c.JSON(400, map[string]string{
			"message": "Capacity required",
		})
		return
	}

	if car.Total == 0 {
		c.JSON(400, map[string]string{
			"message": "Total required",
		})
		return
	}

	if car.Price == 0 {
		c.JSON(400, map[string]string{
			"message": "Price required",
		})
		return
	}

	if car.Id_car_type == 0 {
		c.JSON(400, map[string]string{
			"message": "Id car type required",
		})
		return
	}

	err := cr.cr.CreateCar(car)

	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, map[string]any{
		"message": "success add car",
	})

}

func (cr *Controller) FindAllCar(c *gin.Context) {
	cars, err := cr.cr.FindAllCar()
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find all cars",
		"data":    cars,
	})
}

func (cr *Controller) FindByIDCar(c *gin.Context) {
	id := c.Param("id")

	car, err := cr.cr.FindByIDCar(id)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find by id car",
		"data":    car,
	})
}

func (cr *Controller) UpdateCar(c *gin.Context) {

	var car domain.Car
	// param id
	id := c.Param("id")
	// request body
	if err := c.ShouldBind(&car); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}
	if car.Car_name == "" {
		c.JSON(400, map[string]string{
			"message": "Car name required",
		})
		return
	}

	if car.Cc == 0 {
		c.JSON(400, map[string]string{
			"message": "Cc required",
		})
		return
	}

	if car.Capacity == 0 {
		c.JSON(400, map[string]string{
			"message": "Capacity required",
		})
		return
	}

	if car.Total == 0 {
		c.JSON(400, map[string]string{
			"message": "Total required",
		})
		return
	}

	if car.Price == 0 {
		c.JSON(400, map[string]string{
			"message": "Price required",
		})
		return
	}

	if car.Id_car_type == 0 {
		c.JSON(400, map[string]string{
			"message": "Id car type required",
		})
		return
	}

	err, message := cr.cr.UpdateCar(id, car)
	if err != nil {
		c.JSON(500, map[string]any{
			"message": message,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": message,
	})
}

func (cr *Controller) DeleteCar(c *gin.Context) {

	car := domain.Car{}
	// param id
	id := c.Param("id")

	err, message := cr.cr.DeleteCar(id, car)
	if err != nil {
		c.JSON(500, map[string]any{
			"message": message,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": message,
	})
}
