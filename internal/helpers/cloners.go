package helpers

import "centurypay/internal/models"

func CloneAccount(account *models.Account) *models.Account {
	return &models.Account{
		ID:      account.ID,
		Name:    account.Name,
		Balance: account.Balance,
	}
}

func CloneTransaction(transaction *models.Transaction) *models.Transaction {
	return &models.Transaction{
		ID:          transaction.ID,
		From:        CloneAccount(transaction.From),
		To:          CloneAccount(transaction.To),
		Amount:      transaction.Amount,
		Currency:    transaction.Currency,
		Status:      transaction.Status,
		CreatedAt:   transaction.CreatedAt,
		ConfirmedAt: transaction.ConfirmedAt,
	}
}
