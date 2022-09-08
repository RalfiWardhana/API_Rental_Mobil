package routes

import (
	"rental/connectDB"
	contCarType "rental/controller/car_type"
	gate "rental/middleware"
	repoCarType "rental/repository/car_type"
	repoUsr "rental/repository/user"

	"github.com/gin-gonic/gin"
)

func CarTypeRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoCarType := repoCarType.NewCarTypeRepository(db)
	repoUsr := repoUsr.NewUserRepository(db)
	contCarType := contCarType.NewCarTypeController(repoCarType)

	router.GET("/carTypes", gate.WithAuthentication(repoUsr), contCarType.FindAllCarType)
	router.GET("/carType/:id", gate.WithAuthentication(repoUsr), contCarType.FindByIDCarType)
	router.POST("/carType", gate.WithAuthentication(repoUsr), contCarType.CreateCarType)
	router.PUT("/carType/:id", gate.WithAuthentication(repoUsr), contCarType.UpdateCarType)
	router.DELETE("/carType/:id", gate.WithAuthentication(repoUsr), contCarType.DeleteCarType)
}
