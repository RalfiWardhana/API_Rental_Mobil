package userType

import (
	"github.com/gin-gonic/gin"
)

type UserTypeController interface {
	CreateUserType(c *gin.Context)
	FindAllUserType(c *gin.Context)
	FindByIDUserType(c *gin.Context)
	UpdateUserType(c *gin.Context)
	DeleteUserType(c *gin.Context)
}
