package car

import (
	"github.com/gin-gonic/gin"
)

type CarController interface {
	CreateCar(c *gin.Context)
	FindAllCar(c *gin.Context)
	FindByIDCar(c *gin.Context)
	UpdateCar(c *gin.Context)
	DeleteCar(c *gin.Context)
}
