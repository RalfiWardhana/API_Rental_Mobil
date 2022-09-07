package routes

import (
	"rental/connectDB"
	contUserType "rental/controller/user_type"
	repoUserType "rental/repository/user_type"

	"github.com/gin-gonic/gin"
)

func UserTypeRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoUserType := repoUserType.NewUserTypeRepository(db)
	contUserType := contUserType.NewUserTypeController(repoUserType)

	router.GET("/UserTypes", contUserType.FindAllUserType)
	router.GET("/UserType/:id", contUserType.FindByIDUserType)
	router.POST("/UserType", contUserType.CreateUserType)
	router.PUT("/UserType/:id", contUserType.UpdateUserType)
	router.DELETE("/UserType/:id", contUserType.DeleteUserType)
}
