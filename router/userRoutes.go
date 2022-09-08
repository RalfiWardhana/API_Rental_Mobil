package routes

import (
	"rental/connectDB"
	contUser "rental/controller/user"
	repoUser "rental/repository/user"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoUser := repoUser.NewUserRepository(db)
	contUser := contUser.NewUserController(repoUser)

	router.GET("/users", contUser.FindAllUser)
	router.GET("/user/:id", contUser.FindByIDUser)
	router.POST("/user", contUser.CreateUser)
	router.PUT("/user/:id", contUser.UpdateUser)
	router.DELETE("/user/:id", contUser.DeleteUser)
}
