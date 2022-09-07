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

	router.GET("/transactions", contTransaction.FindAllTransaction)
	router.GET("/transaction/:id", contTransaction.FindByIDTransaction)
	router.POST("/transaction", contTransaction.CreateTransaction)
	router.PUT("/transaction/:id", contTransaction.UpdateTransaction)
	router.DELETE("/transaction/:id", contTransaction.DeleteTransaction)
}
