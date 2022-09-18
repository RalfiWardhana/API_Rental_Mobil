package transaction

import "rental/domain"

type TransactionRepository interface {
	CreateTransaction(transaction domain.Transaction) error
	FindAllTransaction() ([]domain.Transaction_get, error)
	FindByIDTransaction(id string) (domain.Transaction_get, error)
	UpdateTransaction(id string, transaction domain.Transaction) (error, string)
	DeleteTransaction(id string, transaction domain.Transaction) (error, string)
	UpdateTransactionPayment(id string, transaction domain.Transaction_payment) (error, string)
}
