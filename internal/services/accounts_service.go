package services

import (
	"centurypay/internal/enums"
	"centurypay/internal/models"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"sync"
)

type AccountsService struct {
	accounts map[string]*models.Account
	mu       sync.Mutex
}

func NewAccountsService() *AccountsService {
	service := &AccountsService{
		accounts: make(map[string]*models.Account),
	}

	mark := models.NewAccount("acc-"+gonanoid.Must(32), "Mark", 100.0, enums.USD)
	service.accounts[mark.ID] = mark

	jane := models.NewAccount("acc-"+gonanoid.Must(32), "Jane", 50.0, enums.USD)
	service.accounts[jane.ID] = jane

	adam := models.NewAccount("acc-"+gonanoid.Must(32), "Adam", 0.0, enums.USD)
	service.accounts[adam.ID] = adam

	return service
}

func (s *AccountsService) GetAccount(id string) *models.Account {
	s.mu.Lock()
	defer s.mu.Unlock()

	account, exists := s.accounts[id]
	if !exists {
		return nil
	}

	return account
}

func (s *AccountsService) GetAccounts() []*models.Account {
	s.mu.Lock()
	defer s.mu.Unlock()

	accounts := make([]*models.Account, 0, len(s.accounts))
	for _, account := range s.accounts {
		accounts = append(accounts, account)
	}

	return accounts
}

func (s *AccountsService) CreateAccount(account *models.Account) *models.Account {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.accounts[account.ID] = account
	return account
}

func (s *AccountsService) UpdateAccount(account *models.Account) *models.Account {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.accounts[account.ID] = account
	return account
}

func (s *AccountsService) Stop() error {
	// Can be replaced with some cleanup and graceful shutdown logic
	return nil
}
