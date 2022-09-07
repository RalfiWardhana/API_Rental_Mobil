package routes

import (
	"rental/connectDB"
	contCarType "rental/controller/car_type"
	repoCarType "rental/repository/car_type"

	"github.com/gin-gonic/gin"
)

func CarTypeRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoCarType := repoCarType.NewCarTypeRepository(db)
	contCarType := contCarType.NewCarTypeController(repoCarType)

	router.GET("/carTypes", contCarType.FindAllCarType)
	router.GET("/carType/:id", contCarType.FindByIDCarType)
	router.POST("/carType", contCarType.CreateCarType)
	router.PUT("/carType/:id", contCarType.UpdateCarType)
	router.DELETE("/carType/:id", contCarType.DeleteCarType)
}
