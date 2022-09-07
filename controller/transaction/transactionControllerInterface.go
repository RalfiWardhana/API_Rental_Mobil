package transaction

import (
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	CreateTransaction(c *gin.Context)
	FindAllTransaction(c *gin.Context)
	FindByIDTransaction(c *gin.Context)
	UpdateTransaction(c *gin.Context)
	DeleteTransaction(c *gin.Context)
}
