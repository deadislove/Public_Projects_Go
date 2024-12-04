// services/account_service_test.go
package services

import (
	"testing"
)

func TestAccountOperations(t *testing.T) {
	// Create a new account
	account := NewAccount("12345", "Jane Doe")

	// Test Deposit
	account.Deposit(500)
	if account.GetBalance() != 500 {
		t.Errorf("Expected balance 500, got %.2f", account.GetBalance())
	}

	// Test Withdraw
	err := account.Withdraw(200)
	if err != nil {
		t.Errorf("Withdraw failed: %v", err)
	}
	if account.GetBalance() != 300 {
		t.Errorf("Expected balance 300, got %.2f", account.GetBalance())
	}

	// Test Insufficient Funds
	err = account.Withdraw(400)
	if err == nil {
		t.Error("Expected error for insufficient funds, got nil")
	}

	// Test Invalid Deposit
	account.Deposit(-100)
	if account.GetBalance() != 300 {
		t.Errorf("Expected balance 300 after invalid deposit, got %.2f", account.GetBalance())
	}
}
