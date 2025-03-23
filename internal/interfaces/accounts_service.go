package interfaces

import (
	"centurypay/internal/enums"
	"centurypay/internal/models"
)

type AccountsService interface {
	CreateAccount(name string, balanceAmount float64, balanceCurrency enums.Currency) (*models.Account, error)
	GetAccount(id string) (*models.Account, error)
	GetAccounts() []*models.Account
	UpdateAccount(account *models.Account) error
	Stop() error
}
