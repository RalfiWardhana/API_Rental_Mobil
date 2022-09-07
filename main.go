package main

import (
	"fmt"
	"rental/connectDB"
	routing "rental/router"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cfg := connectDB.ConnectDB()
	fmt.Println(cfg)

	routing.CarRoute(router)
	routing.CarTypeTypeRoute(router)
	routing.StatusTypeRoute(router)
	routing.TransactionRoute(router)
	routing.UserTypeRoute(router)

	router.Run("localhost:9000")

}
