package routes

import (
	"rental/connectDB"
	contUserType "rental/controller/user_type"
	gate "rental/middleware"
	repoUsr "rental/repository/user"
	repoUserType "rental/repository/user_type"

	"github.com/gin-gonic/gin"
)

func UserTypeRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoUserType := repoUserType.NewUserTypeRepository(db)
	repoUsr := repoUsr.NewUserRepository(db)
	contUserType := contUserType.NewUserTypeController(repoUserType)

	router.GET("/userTypes", gate.WithAuthentication(repoUsr), contUserType.FindAllUserType)
	router.GET("/userType/:id", gate.WithAuthentication(repoUsr), contUserType.FindByIDUserType)
	router.POST("/userType", gate.WithAuthentication(repoUsr), contUserType.CreateUserType)
	router.PUT("/userType/:id", gate.WithAuthentication(repoUsr), contUserType.UpdateUserType)
	router.DELETE("/userType/:id", gate.WithAuthentication(repoUsr), contUserType.DeleteUserType)
}
