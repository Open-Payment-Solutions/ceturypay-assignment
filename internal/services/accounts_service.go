package services

import (
	"errors"
	"fmt"
	"sync"

	"centurypay/internal/enums"
	"centurypay/internal/helpers"
	"centurypay/internal/models"
)

var (
	ErrAccountNotFound = errors.New("account not found")
)

type AccountsService struct {
	mu       sync.RWMutex
	accounts map[string]*models.Account
}

func NewAccountsService() *AccountsService {
	return &AccountsService{
		mu:       sync.RWMutex{},
		accounts: make(map[string]*models.Account),
	}
}

func (s *AccountsService) GetAccount(
	id string,
) (*models.Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	account, exists := s.accounts[id]
	if !exists {
		return nil, ErrAccountNotFound
	}

	return helpers.CloneAccount(account), nil
}

func (s *AccountsService) GetAccounts() []*models.Account {
	s.mu.RLock()
	defer s.mu.RUnlock()

	accounts := make([]*models.Account, 0, len(s.accounts))
	for _, account := range s.accounts {
		accounts = append(
			accounts,
			helpers.CloneAccount(account),
		)
	}

	return accounts
}

func (s *AccountsService) CreateAccount(
	name string,
	balanceAmount float64,
	balanceCurrency enums.Currency,
) (*models.Account, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	account := models.NewAccount(name, balanceAmount, balanceCurrency)
	account.ID = fmt.Sprintf("%d", len(s.accounts)+1) // uncomment this for api request testing
	s.accounts[account.ID] = account

	return helpers.CloneAccount(account), nil
}

func (s *AccountsService) UpdateAccount(account *models.Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.accounts[account.ID]; !exists {
		return ErrAccountNotFound
	}

	s.accounts[account.ID] = helpers.CloneAccount(account)

	return nil
}

func (s *AccountsService) Stop() error {
	// Can be replaced with some cleanup and graceful shutdown logic
	return nil
}
