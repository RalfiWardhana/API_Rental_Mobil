package routes

import (
	"rental/connectDB"
	contStatusType "rental/controller/status_type"
	gate "rental/middleware"
	repoStatusType "rental/repository/status_type"
	repoUsr "rental/repository/user"

	"github.com/gin-gonic/gin"
)

func StatusTypeRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoStatusType := repoStatusType.NewStatusTypeRepository(db)
	repoUsr := repoUsr.NewUserRepository(db)
	contStatusType := contStatusType.NewStatusTypeController(repoStatusType)

	router.GET("/statusTypes", gate.WithAuthentication(repoUsr), contStatusType.FindAllStatusType)
	router.GET("/statusType/:id", gate.WithAuthentication(repoUsr), contStatusType.FindByIDStatusType)
	router.POST("/statusType", gate.WithAuthentication(repoUsr), contStatusType.CreateStatusType)
	router.PUT("/statusType/:id", gate.WithAuthentication(repoUsr), contStatusType.UpdateStatusType)
	router.DELETE("/statusType/:id", gate.WithAuthentication(repoUsr), contStatusType.DeleteStatusType)
}
