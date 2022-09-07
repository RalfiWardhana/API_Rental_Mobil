package routes

import (
	"rental/connectDB"
	contCar "rental/controller/car"
	repoCar "rental/repository/car"

	"github.com/gin-gonic/gin"
)

func CarRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoCar := repoCar.NewCarRepository(db)
	contCar := contCar.NewCarController(repoCar)

	router.GET("/cars", contCar.FindAllCar)
	router.GET("/car/:id", contCar.FindByIDCar)
	router.POST("/car", contCar.CreateCar)
	router.PUT("/car/:id", contCar.UpdateCar)
	router.DELETE("/car/:id", contCar.DeleteCar)
}
