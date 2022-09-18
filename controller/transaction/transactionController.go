package transaction

import (
	"rental/domain"
	"rental/repository/transaction"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	cr transaction.TransactionRepository
}

func NewTransactionController(cr transaction.TransactionRepository) TransactionController {
	return &Controller{
		cr: cr,
	}
}

func (cr *Controller) CreateTransaction(c *gin.Context) {

	var Transaction domain.Transaction

	if err := c.ShouldBind(&Transaction); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	if Transaction.Id_car == 0 {
		c.JSON(400, map[string]string{
			"message": "Id car required",
		})
		return
	}

	if Transaction.Id_user == 0 {
		c.JSON(400, map[string]string{
			"message": "Id user required",
		})
		return
	}

	if Transaction.Total_price == 0 {
		c.JSON(400, map[string]string{
			"message": "Total price required",
		})
		return
	}

	if Transaction.Duration == "" {
		c.JSON(400, map[string]string{
			"message": "Duration required",
		})
		return
	}

	if Transaction.Status == 0 {
		c.JSON(400, map[string]string{
			"message": "Status required",
		})
		return
	}

	err := cr.cr.CreateTransaction(Transaction)

	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, map[string]any{
		"message": "success add Transaction",
	})

}

func (cr *Controller) FindAllTransaction(c *gin.Context) {
	Transactions, err := cr.cr.FindAllTransaction()
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find all Transactions",
		"data":    Transactions,
	})
}

func (cr *Controller) FindByIDTransaction(c *gin.Context) {
	id := c.Param("id")

	Transaction, err := cr.cr.FindByIDTransaction(id)
	if err != nil {
		c.JSON(500, map[string]any{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": "success find by id Transaction",
		"data":    Transaction,
	})
}

func (cr *Controller) UpdateTransaction(c *gin.Context) {

	var Transaction domain.Transaction
	// param id
	id := c.Param("id")
	// request body
	if err := c.ShouldBind(&Transaction); err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid input",
		})
		return
	}

	if Transaction.Id_car == 0 {
		c.JSON(400, map[string]string{
			"message": "Id car required",
		})
		return
	}

	if Transaction.Id_user == 0 {
		c.JSON(400, map[string]string{
			"message": "Id user required",
		})
		return
	}

	if Transaction.Total_price == 0 {
		c.JSON(400, map[string]string{
			"message": "Total price required",
		})
		return
	}

	if Transaction.Duration == "" {
		c.JSON(400, map[string]string{
			"message": "Duration required",
		})
		return
	}

	if Transaction.Status == 0 {
		c.JSON(400, map[string]string{
			"message": "Status required",
		})
		return
	}

	err, message := cr.cr.UpdateTransaction(id, Transaction)
	if err != nil {
		c.JSON(500, map[string]any{
			"message": message,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, map[string]any{
		"message": message,
	})
}

func (cr *Controller) DeleteTransaction(c *gin.Context) {

	Transaction := domain.Transaction{}
	// param id
	id := c.Param("id")

	err, message := cr.cr.DeleteTransaction(id, Transaction)
	if err != nil {
		c.JSON(500, map[string]any{
			"message": message,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": message,
	})
}

func (cr *Controller) UpdateTransactionPayment(c *gin.Context) {
	// param id
	id := c.Param("id")

	file, _ := c.FormFile("Payment")
	Transaction := domain.Transaction_payment{
		Payment: file.Filename,
	}

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, "car_photo/"+file.Filename)

	err, message := cr.cr.UpdateTransactionPayment(id, Transaction)
	if err != nil {
		c.JSON(500, map[string]any{
			"message": message,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, map[string]any{
		"message": message,
	})
}
