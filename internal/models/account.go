package models

import (
	"centurypay/internal/enums"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Account struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Balance AccountBalance `json:"balance"`
}

type AccountBalance struct {
	Amount   float64        `json:"amount"`
	Currency enums.Currency `json:"currency"`
}

func NewAccount(
	name string,
	balanceAmount float64,
	balanceCurrency enums.Currency,
) *Account {
	return &Account{
		ID:   generateAccountId(),
		Name: name,
		Balance: AccountBalance{
			Amount:   balanceAmount,
			Currency: balanceCurrency,
		},
	}
}

func generateAccountId() string {
	id := gonanoid.Must(32)
	return "acc-" + id
}
