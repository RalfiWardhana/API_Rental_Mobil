package routes

import (
	"rental/connectDB"
	contCar "rental/controller/car"
	gate "rental/middleware"
	repoCar "rental/repository/car"
	repoUsr "rental/repository/user"

	"github.com/gin-gonic/gin"
)

func CarRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoCar := repoCar.NewCarRepository(db)
	repoUsr := repoUsr.NewUserRepository(db)
	contCar := contCar.NewCarController(repoCar)

	router.GET("/cars", gate.WithAuthentication(repoUsr), contCar.FindAllCar)
	router.GET("/car/:id", gate.WithAuthentication(repoUsr), contCar.FindByIDCar)
	router.POST("/car", gate.WithAuthentication(repoUsr), contCar.CreateCar)
	router.PUT("/car/:id", gate.WithAuthentication(repoUsr), contCar.UpdateCar)
	router.DELETE("/car/:id", gate.WithAuthentication(repoUsr), contCar.DeleteCar)
}
