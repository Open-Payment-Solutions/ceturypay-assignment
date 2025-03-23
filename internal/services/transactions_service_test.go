package services

import (
	"testing"
	"time"

	"centurypay/internal/enums"
	"centurypay/internal/processors"
)

func TestTransferBetweenAccounts(t *testing.T) {
	accountsService := NewAccountsService()
	transactionsService := NewTransactionsService(accountsService)
	processor := processors.NewTransactionsProcessor(transactionsService)
	transactionsService.SetProcessor(processor)

	// Create test accounts
	account1, _ := accountsService.CreateAccount("Sender", 100.0, enums.USD)
	account2, _ := accountsService.CreateAccount("Receiver", 50.0, enums.USD)

	// Test transfer
	transaction, err := transactionsService.CreateTransaction(account1, account2, 25.0, enums.USD)
	if err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	// Check initial state
	if transaction.Status != enums.TransactionStatusCreated {
		t.Errorf("Expected status Created, got %s", transaction.Status)
	}

	// Wait for processing to complete
	time.Sleep(1 * time.Second)

	// Check transaction completed
	transaction, err = transactionsService.GetTransaction(transaction.ID)
	if err != nil {
		t.Fatalf("Failed to get transaction: %v", err)
	}

	if transaction.Status != enums.TransactionStatusCompleted {
		t.Errorf("Expected status Completed, got %s", transaction.Status)
	}

	// Check account balances
	account1, _ = accountsService.GetAccount(account1.ID)
	account2, _ = accountsService.GetAccount(account2.ID)

	if account1.Balance.Amount != 75.0 {
		t.Errorf("Expected sender balance 75.0, got %.2f", account1.Balance.Amount)
	}

	if account2.Balance.Amount != 75.0 {
		t.Errorf("Expected receiver balance 75.0, got %.2f", account2.Balance.Amount)
	}
}

func TestInsufficientFunds(t *testing.T) {
	accountsService := NewAccountsService()
	transactionsService := NewTransactionsService(accountsService)

	// Create test accounts
	account1, _ := accountsService.CreateAccount("Sender", 10.0, enums.USD)
	account2, _ := accountsService.CreateAccount("Receiver", 50.0, enums.USD)

	// Test transfer with insufficient funds
	_, err := transactionsService.CreateTransaction(account1, account2, 25.0, enums.USD)
	if err == nil || err.Error() != ErrInsufficientFunds.Error() {
		t.Errorf("Expected ErrInsufficientFunds, got %v", err)
	}

	// Check account balances unchanged
	account1, _ = accountsService.GetAccount(account1.ID)
	if account1.Balance.Amount != 10.0 {
		t.Errorf("Expected sender balance unchanged at 10.0, got %.2f", account1.Balance.Amount)
	}
}

func TestSameAccountTransfer(t *testing.T) {
	accountsService := NewAccountsService()
	transactionsService := NewTransactionsService(accountsService)

	// Create test account
	account, _ := accountsService.CreateAccount("Self", 100.0, enums.USD)

	// Test transfer to same account
	_, err := transactionsService.CreateTransaction(account, account, 25.0, enums.USD)
	if err == nil || err.Error() != ErrSameAccount.Error() {
		t.Errorf("Expected ErrSameAccount, got %v", err)
	}
}
