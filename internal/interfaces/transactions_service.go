package interfaces

import (
	"centurypay/internal/enums"
	"centurypay/internal/models"
)

type TransactionsService interface {
	CreateTransaction(fromAccount *models.Account, toAccount *models.Account, amount float64, currency enums.Currency) (*models.Transaction, error)
	GetTransaction(id string) (*models.Transaction, error)
	SetPendingTransaction(id string) (*models.Transaction, error)
	ConfirmTransaction(id string) (*models.Transaction, error)
	CompleteTransaction(id string) (*models.Transaction, error)
	RejectTransaction(id string) (*models.Transaction, error)
	Stop() error
}
