package carType

import (
	"rental/domain"
	carType "rental/repository/car_type"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	cr carType.CarTypeRepository
}

func NewCarTypeController(cr carType.CarTypeRepository) CarTypeController {
	return &Controller{
		cr: cr,
	}
}

func (cr *Controller) CreateCarType(c *gin.Context) {

	var carType domain.Car_type

	if err := c.ShouldBind(&carType); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	if carType.Car_type == "" {
		c.JSON(400, map[string]string{
			"message": "CarType name required",
		})
		return
	}

	err := cr.cr.CreateCarType(carType)

	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, map[string]any{
		"message": "success add carType",
	})

}

func (cr *Controller) FindAllCarType(c *gin.Context) {
	carTypes, err := cr.cr.FindAllCarType()
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find all carTypes",
		"data":    carTypes,
	})
}

func (cr *Controller) FindByIDCarType(c *gin.Context) {
	id := c.Param("id")

	carType, err := cr.cr.FindByIDCarType(id)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find by id carType",
		"data":    carType,
	})
}

func (cr *Controller) UpdateCarType(c *gin.Context) {
	// domain user
	var carType domain.Car_type
	// param id
	id := c.Param("id")
	// request body
	if err := c.ShouldBind(&carType); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}
	if carType.Car_type == "" {
		c.JSON(400, map[string]string{
			"message": "CarType name required",
		})
		return
	}

	// update user
	err, message := cr.cr.UpdateCarType(id, carType)
	if err != nil {
		c.JSON(500, map[string]any{
			"message": message,
			"error":   err.Error(),
		})
		return
	}
	// return success update user
	c.JSON(200, map[string]any{
		"message": message,
	})
}

func (cr *Controller) DeleteCarType(c *gin.Context) {
	// domain user
	carType := domain.Car_type{}
	// param id
	id := c.Param("id")

	// delete user
	err, message := cr.cr.DeleteCarType(id, carType)
	if err != nil {
		c.JSON(500, map[string]any{
			"message": message,
			"error":   err.Error(),
		})
		return
	}
	// return success delete user
	c.JSON(200, map[string]any{
		"message": message,
	})
}
