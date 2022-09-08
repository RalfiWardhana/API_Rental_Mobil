package user

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(c *gin.Context)
	FindAllUser(c *gin.Context)
	FindByIDUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
