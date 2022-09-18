package main

import (
	"fmt"
	"mime/multipart"
	"rental/connectDB"
	routing "rental/router"

	"github.com/gin-gonic/gin"
)

// MailUpload struct for e-mail
type MailUpload struct {
	Mail        string                `json:"mail" form:"mail" binding:"required"`
	Attachments *multipart.FileHeader `json:"attachment" form:"attachments" binding:"required"`
}

func Mail(c *gin.Context) {
	var postData MailUpload
	err := c.ShouldBind(&postData)
	if err != nil {
		fmt.Println(err)
		c.String(400, err.Error())
	} else {
		fmt.Printf("%#v\n", postData.Attachments.Filename)
		c.String(200, "ok")
	}
}

func main() {
	router := gin.Default()
	cfg := connectDB.ConnectDB()
	fmt.Println(cfg)

	routing.CarRoute(router)
	routing.CarTypeRoute(router)
	routing.StatusTypeRoute(router)
	routing.TransactionRoute(router)
	routing.UserTypeRoute(router)
	routing.UserRoute(router)
	router.POST("/mail", Mail)

	router.Run("localhost:9000")

}
