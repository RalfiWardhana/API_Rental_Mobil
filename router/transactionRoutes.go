package routes

import (
	"rental/connectDB"
	contTransaction "rental/controller/transaction"
	repoTransaction "rental/repository/transaction"

	"github.com/gin-gonic/gin"
)

func TransactionRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoTransaction := repoTransaction.NewTransactionRepository(db)
	contTransaction := contTransaction.NewTransactionController(repoTransaction)

	router.GET("/Transactions", contTransaction.FindAllTransaction)
	router.GET("/Transaction/:id", contTransaction.FindByIDTransaction)
	router.POST("/Transaction", contTransaction.CreateTransaction)
	router.PUT("/Transaction/:id", contTransaction.UpdateTransaction)
	router.DELETE("/Transaction/:id", contTransaction.DeleteTransaction)
}
