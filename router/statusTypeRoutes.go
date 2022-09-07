package routes

import (
	"rental/connectDB"
	contStatusType "rental/controller/status_type"
	repoStatusType "rental/repository/status_type"

	"github.com/gin-gonic/gin"
)

func StatusTypeRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoStatusType := repoStatusType.NewStatusTypeRepository(db)
	contStatusType := contStatusType.NewStatusTypeController(repoStatusType)

	router.GET("/StatusTypes", contStatusType.FindAllStatusType)
	router.GET("/StatusType/:id", contStatusType.FindByIDStatusType)
	router.POST("/StatusType", contStatusType.CreateStatusType)
	router.PUT("/StatusType/:id", contStatusType.UpdateStatusType)
	router.DELETE("/StatusType/:id", contStatusType.DeleteStatusType)
}
