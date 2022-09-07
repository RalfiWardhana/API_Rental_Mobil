package transaction

import "rental/entity"

type TransactionRepository interface {
	CreateTransaction(transaction entity.Transaction) error
	FindAllTransaction() ([]entity.Transaction_get, error)
	FindByIDTransaction(id string) (entity.Transaction_get, error)
	UpdateTransaction(id string, transaction entity.Transaction) (error, string)
	DeleteTransaction(id string, transaction entity.Transaction) (error, string)
}
