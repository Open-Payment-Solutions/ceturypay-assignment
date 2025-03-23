package interfaces

import "centurypay/internal/models"

type TransactionProcessor interface {
	Start() error
	SubmitTransaction(transaction *models.Transaction) error
	Stop() error
}
