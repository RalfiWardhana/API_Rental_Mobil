package carType

import (
	"github.com/gin-gonic/gin"
)

type CarTypeController interface {
	CreateCarType(c *gin.Context)
	FindAllCarType(c *gin.Context)
	FindByIDCarType(c *gin.Context)
	UpdateCarType(c *gin.Context)
	DeleteCarType(c *gin.Context)
}
