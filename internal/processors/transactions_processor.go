package processors

import (
	"time"

	"centurypay/internal/interfaces"
	"centurypay/internal/models"
)

type TransactionsProcessor struct {
	transactionService interfaces.TransactionsService
}

func NewTransactionsProcessor(transactionService interfaces.TransactionsService) *TransactionsProcessor {
	return &TransactionsProcessor{
		transactionService: transactionService,
	}
}

func (p *TransactionsProcessor) SubmitTransaction(transaction *models.Transaction) error {
	_, err := p.transactionService.SetPendingTransaction(transaction.ID)
	if err != nil {
		return err
	}

	go func() {
		time.Sleep(500 * time.Millisecond)

		confirmed, err := p.transactionService.ConfirmTransaction(transaction.ID)
		if err != nil {
			return
		}

		_, err = p.transactionService.CompleteTransaction(confirmed.ID)
		if err != nil {
			return
		}
	}()

	return nil
}

func (p *TransactionsProcessor) Start() error {
	// Could initialize background workers or connections here
	return nil
}

func (p *TransactionsProcessor) Stop() error {
	// Cleanup resources
	return nil
}
