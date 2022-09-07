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

	router.GET("/userTypes", contUserType.FindAllUserType)
	router.GET("/userType/:id", contUserType.FindByIDUserType)
	router.POST("/userType", contUserType.CreateUserType)
	router.PUT("/userType/:id", contUserType.UpdateUserType)
	router.DELETE("/userType/:id", contUserType.DeleteUserType)
}
