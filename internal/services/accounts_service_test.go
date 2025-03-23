package services

import (
	"centurypay/internal/enums"
	"testing"
)

func TestAccountCreation(t *testing.T) {
	service := NewAccountsService()

	account, err := service.CreateAccount("Test", 100.0, enums.USD)
	if err != nil {
		t.Fatalf("Failed to create account: %v", err)
	}

	if account.Name != "Test" {
		t.Errorf("Expected name 'Test', got '%s'", account.Name)
	}

	if account.Balance.Amount != 100.0 {
		t.Errorf("Expected balance 100.0, got %.2f", account.Balance.Amount)
	}

	if account.Balance.Currency != enums.USD {
		t.Errorf("Expected currency USD, got %s", account.Balance.Currency)
	}
}

func TestGetAccount(t *testing.T) {
	service := NewAccountsService()

	account, _ := service.CreateAccount("Test", 100.0, enums.USD)
	retrieved, err := service.GetAccount(account.ID)

	if err != nil {
		t.Fatalf("Failed to get account: %v", err)
	}

	if retrieved.ID != account.ID {
		t.Errorf("Expected ID %s, got %s", account.ID, retrieved.ID)
	}
}

func TestGetNonExistentAccount(t *testing.T) {
	service := NewAccountsService()

	_, err := service.GetAccount("non-existent-id")
	if err == nil || err.Error() != ErrAccountNotFound.Error() {
		t.Errorf("Expected ErrAccountNotFound, got %v", err)
	}
}
