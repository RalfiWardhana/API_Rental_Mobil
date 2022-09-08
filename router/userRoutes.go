package routes

import (
	"rental/connectDB"
	contUser "rental/controller/user"
	gate "rental/middleware"
	repoUser "rental/repository/user"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoUser := repoUser.NewUserRepository(db)
	contUser := contUser.NewUserController(repoUser)

	router.GET("/users", gate.WithAuthentication(repoUser), contUser.FindAllUser)
	router.GET("/user/:id", gate.WithAuthentication(repoUser), contUser.FindByIDUser)
	router.POST("/register", contUser.CreateUser)
	router.POST("/login", contUser.Login)
	router.PUT("/user/:id", gate.WithAuthentication(repoUser), contUser.UpdateUser)
	router.DELETE("/user/:id", gate.WithAuthentication(repoUser), contUser.DeleteUser)
}
