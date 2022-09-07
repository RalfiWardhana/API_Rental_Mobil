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

	router.GET("/statusTypes", contStatusType.FindAllStatusType)
	router.GET("/statusType/:id", contStatusType.FindByIDStatusType)
	router.POST("/statusType", contStatusType.CreateStatusType)
	router.PUT("/statusType/:id", contStatusType.UpdateStatusType)
	router.DELETE("/statusType/:id", contStatusType.DeleteStatusType)
}
