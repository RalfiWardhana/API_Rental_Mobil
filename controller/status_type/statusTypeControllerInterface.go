package statusType

import (
	"github.com/gin-gonic/gin"
)

type StatusTypeController interface {
	CreateStatusType(c *gin.Context)
	FindAllStatusType(c *gin.Context)
	FindByIDStatusType(c *gin.Context)
	UpdateStatusType(c *gin.Context)
	DeleteStatusType(c *gin.Context)
}
