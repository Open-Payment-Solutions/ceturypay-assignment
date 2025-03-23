package processors

import (
	"centurypay/internal/interfaces"
	"centurypay/internal/models"
	"sync"
)

type TransactionProcessor struct {
	transactionService interfaces.TransactionService
	accountService     interfaces.AccountsService

	queue chan *models.Transaction
	wg    sync.WaitGroup
}

func NewTransactionsProcessor(
	transactionService interfaces.TransactionService,
	accountService interfaces.AccountsService,
) *TransactionProcessor {
	return &TransactionProcessor{
		transactionService: transactionService,
		accountService:     accountService,

		queue: make(chan *models.Transaction, 100),
		wg:    sync.WaitGroup{},
	}
}

func (p *TransactionProcessor) Start() error {
	return nil
}

func (p *TransactionProcessor) SubmitTransaction(transaction *models.Transaction) error {
	return nil
}

func (p *TransactionProcessor) Stop() error {
	return nil
}
