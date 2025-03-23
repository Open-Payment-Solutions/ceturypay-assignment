package services

import (
	"errors"
	"sync"
	"time"

	"centurypay/internal/enums"
	"centurypay/internal/helpers"
	"centurypay/internal/interfaces"
	"centurypay/internal/models"
)

var (
	ErrInsufficientFunds                = errors.New("insufficient funds")
	ErrSameAccount                      = errors.New("cannot transfer to the same account")
	ErrTransactionNotFound              = errors.New("transaction not found")
	ErrTransactionMustHaveCorrectStatus = errors.New("transaction must have correct status")
	ErrTransactionExpired               = errors.New("transaction has expired")
	ErrDifferentCurrencies              = errors.New("accounts have different currencies")
	ErrNegativeAmount                   = errors.New("amount must be positive")
)

type TransactionsService struct {
	accountService interfaces.AccountsService

	mu           sync.RWMutex
	transactions map[string]*models.Transaction

	processor interfaces.TransactionProcessor
}

func (s *TransactionsService) SetProcessor(
	processor interfaces.TransactionProcessor,
) {
	s.processor = processor
}

func NewTransactionsService(
	accountService interfaces.AccountsService,
) *TransactionsService {
	return &TransactionsService{
		accountService: accountService,

		mu:           sync.RWMutex{},
		transactions: make(map[string]*models.Transaction),
	}
}

func (s *TransactionsService) CreateTransaction(
	fromAccount *models.Account,
	toAccount *models.Account,
	amount float64,
	currency enums.Currency,
) (*models.Transaction, error) {
	if amount <= 0 {
		return nil, ErrNegativeAmount
	}

	if fromAccount.ID == toAccount.ID {
		return nil, ErrSameAccount
	}

	s.mu.Lock()

	if fromAccount.Balance.Currency != currency ||
		toAccount.Balance.Currency != currency {
		return nil, ErrDifferentCurrencies
	}

	if fromAccount.Balance.Amount < amount {
		return nil, ErrInsufficientFunds
	}

	fromAccount.Balance.Amount -= amount
	err := s.accountService.UpdateAccount(fromAccount)
	if err != nil {
		return nil, err
	}

	transaction := models.NewTransaction(
		fromAccount, toAccount, amount, currency,
	)
	s.transactions[transaction.ID] = transaction

	s.mu.Unlock()

	tx := helpers.CloneTransaction(transaction)
	err = s.processor.SubmitTransaction(tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *TransactionsService) GetTransaction(
	id string,
) (*models.Transaction, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	transaction, exists := s.transactions[id]
	if !exists {
		return nil, ErrTransactionNotFound
	}

	return helpers.CloneTransaction(transaction), nil
}

func (s *TransactionsService) SetPendingTransaction(
	id string,
) (*models.Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	transaction, exists := s.transactions[id]
	if !exists {
		return nil, ErrTransactionNotFound
	}

	if transaction.Status != enums.TransactionStatusCreated {
		return nil, ErrTransactionMustHaveCorrectStatus
	}

	transaction.Status = enums.TransactionStatusPending
	transaction.PendingAt = time.Now()

	return helpers.CloneTransaction(transaction), nil
}

func (s *TransactionsService) ConfirmTransaction(
	id string,
) (*models.Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	transaction, exists := s.transactions[id]
	if !exists {
		return nil, ErrTransactionNotFound
	}

	if transaction.Status != enums.TransactionStatusPending {
		return nil, ErrTransactionMustHaveCorrectStatus
	}

	if transaction.LockedUntil.Before(time.Now()) {
		transaction.Status = enums.TransactionStatusExpired

		fromAccount := transaction.From
		fromAccount.Balance.Amount += transaction.Amount

		err := s.accountService.UpdateAccount(fromAccount)
		if err != nil {
			return nil, err
		}

		return nil, ErrTransactionExpired
	}

	transaction.Status = enums.TransactionStatusConfirmed
	transaction.ConfirmedAt = time.Now()

	return helpers.CloneTransaction(transaction), nil
}

func (s *TransactionsService) CompleteTransaction(
	id string,
) (*models.Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	transaction, exists := s.transactions[id]
	if !exists {
		return nil, ErrTransactionNotFound
	}

	if transaction.Status != enums.TransactionStatusConfirmed {
		return nil, ErrTransactionMustHaveCorrectStatus
	}

	originalStatus := transaction.Status
	transaction.Status = enums.TransactionStatusCompleted
	transaction.CompletedAt = time.Now()

	toAccount := transaction.To
	toAccount.Balance.Amount += transaction.Amount

	err := s.accountService.UpdateAccount(toAccount)
	if err != nil {
		transaction.Status = originalStatus
		return nil, err
	}

	return helpers.CloneTransaction(transaction), nil
}

func (s *TransactionsService) RejectTransaction(
	id string,
) (*models.Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	transaction, exists := s.transactions[id]
	if !exists {
		return nil, ErrTransactionNotFound
	}

	if transaction.Status != enums.TransactionStatusPending {
		return nil, ErrTransactionMustHaveCorrectStatus
	}

	originalStatus := transaction.Status
	transaction.Status = enums.TransactionStatusRejected
	transaction.RejectedAt = time.Now()

	fromAccount := transaction.From
	fromAccount.Balance.Amount += transaction.Amount

	err := s.accountService.UpdateAccount(fromAccount)
	if err != nil {
		transaction.Status = originalStatus
		return nil, err
	}

	return helpers.CloneTransaction(transaction), nil
}

func (s *TransactionsService) Stop() error {
	// Can be replaced with some cleanup and graceful shutdown logic
	return nil
}
