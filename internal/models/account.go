package models

import "centurypay/internal/enums"

type Account struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Balance AccountBalance `json:"balance"`
}

type AccountBalance struct {
	Amount   float64        `json:"amount"`
	Currency enums.Currency `json:"currency"`
}

func NewAccount(id string, name string, balanceAmount float64, balanceCurrency enums.Currency) *Account {
	return &Account{
		ID:   id,
		Name: name,
		Balance: AccountBalance{
			Amount:   balanceAmount,
			Currency: balanceCurrency,
		},
	}
}
