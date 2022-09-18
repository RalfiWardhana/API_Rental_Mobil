package routes

import (
	"rental/connectDB"
	contTransaction "rental/controller/transaction"
	gate "rental/middleware"
	repoTransaction "rental/repository/transaction"
	repoUsr "rental/repository/user"

	"github.com/gin-gonic/gin"
)

func TransactionRoute(router *gin.Engine) {

	db := connectDB.ConnectDB()
	repoTransaction := repoTransaction.NewTransactionRepository(db)
	repoUsr := repoUsr.NewUserRepository(db)
	contTransaction := contTransaction.NewTransactionController(repoTransaction)

	router.GET("/transactions", gate.WithAuthentication(repoUsr), contTransaction.FindAllTransaction)
	router.GET("/transaction/:id", gate.WithAuthentication(repoUsr), contTransaction.FindByIDTransaction)
	router.POST("/transaction", gate.WithAuthentication(repoUsr), contTransaction.CreateTransaction)
	router.PUT("/transaction/:id", contTransaction.UpdateTransaction)
	router.PUT("/transaction-payment/:id", contTransaction.UpdateTransactionPayment)
	router.DELETE("/transaction/:id", gate.WithAuthentication(repoUsr), contTransaction.DeleteTransaction)
}
