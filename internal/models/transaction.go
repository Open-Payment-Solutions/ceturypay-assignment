package models

import (
	"centurypay/internal/enums"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"time"
)

type Transaction struct {
	ID          string                  `json:"id"`
	From        *Account                `json:"from"`
	To          *Account                `json:"to"`
	Amount      float64                 `json:"amount"`
	Currency    enums.Currency          `json:"currency"`
	Status      enums.TransactionStatus `json:"status"`
	CreatedAt   time.Time               `json:"createdAt"`
	PendingAt   time.Time               `json:"pendingAt,omitempty"`
	RejectedAt  time.Time               `json:"rejectedAt,omitempty"`
	ConfirmedAt time.Time               `json:"confirmedAt,omitempty"`
	CompletedAt time.Time               `json:"completedAt,omitempty"`
	LockedUntil time.Time               `json:"lockedUntil,omitempty"`
}

func NewTransaction(
	fromAccount *Account,
	toAccount *Account,
	amount float64,
	currency enums.Currency,
) *Transaction {
	return &Transaction{
		ID:          generateTransactionId(),
		From:        fromAccount,
		To:          toAccount,
		Amount:      amount,
		Currency:    currency,
		Status:      enums.TransactionStatusCreated,
		CreatedAt:   time.Now(),
		LockedUntil: time.Now().Add(time.Minute), // amount will be locked for 1 minute
	}
}

func (t *Transaction) Copy() *Transaction {
	return &Transaction{
		ID:          t.ID,
		From:        t.From,
		To:          t.To,
		Amount:      t.Amount,
		Currency:    t.Currency,
		Status:      t.Status,
		CreatedAt:   t.CreatedAt,
		PendingAt:   t.PendingAt,
		RejectedAt:  t.RejectedAt,
		ConfirmedAt: t.ConfirmedAt,
		CompletedAt: t.CompletedAt,
		LockedUntil: t.LockedUntil,
	}
}

func generateTransactionId() string {
	id := gonanoid.Must(32)
	return "tx-" + id
}
