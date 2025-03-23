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
	ConfirmedAt time.Time               `json:"confirmedAt,omitempty"`
	LockedUntil time.Time               `json:"lockedUntil,omitempty"`
}

func NewTransaction(from *Account, to *Account, amount float64, currency enums.Currency) *Transaction {
	return &Transaction{
		ID:        generateTransactionId(), // You'll need to implement this function
		From:      from,
		To:        to,
		Amount:    amount,
		Currency:  currency,
		Status:    enums.TransactionStatusCreated,
		CreatedAt: time.Now(),
		// Lock the funds for 1 minute
		LockedUntil: time.Now().Add(time.Minute),
	}
}

func generateTransactionId() string {
	id := gonanoid.Must(32)
	return "tx-" + id
}
