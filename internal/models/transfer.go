package models

import (
	"centurypay/internal/enums"
	"time"
)

type Transfer struct {
	From        *Account             `json:"from"`
	To          *Account             `json:"to"`
	Amount      float64              `json:"amount"`
	Currency    enums.Currency       `json:"currency"`
	Status      enums.TransferStatus `json:"status"`
	CreatedAt   time.Time            `json:"createdAt"`
	ConfirmedAt time.Time            `json:"confirmedAt"`
}
