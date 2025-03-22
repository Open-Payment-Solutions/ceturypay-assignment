package services

import (
	"centurypay/internal/models"
)

type AccountsService struct {
}

func NewAccountsService() *AccountsService {
	return &AccountsService{}
}

func (s *AccountsService) GetAccount(id string) *models.Account {
	// ToDo: Implement the logic to get the account from the storage
	return nil
}

func (s *AccountsService) GetAccounts() []*models.Account {
	// ToDo: Implement the logic to get the accounts from the storage
	return []*models.Account{}
}

func (s *AccountsService) CreateAccount(account *models.Account) *models.Account {
	// ToDo: Implement the logic to create the account in the storage
	return nil
}

func (s *AccountsService) UpdateAccount(account *models.Account) *models.Account {
	// ToDo: Implement the logic to update the account in the storage
	return nil
}

func (s *AccountsService) Stop() error {
	// ToDo: Implement the logic to stop the service
	return nil
}
