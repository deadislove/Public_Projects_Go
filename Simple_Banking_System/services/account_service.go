package services

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber string
	AccountHolder string
	Balance       float64
}

// NewAccount creates and returns a new Account instance
func NewAccount(accountNumber, accountHolder string) *Account {
	return &Account{
		AccountNumber: accountNumber,
		AccountHolder: accountHolder,
		Balance:       0.0,
	}
}

// Deposit adds a given amount to the account balance
func (a *Account) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Amount to deposit should be positive.")
		return
	}
	a.Balance += amount
	fmt.Printf("Deposited: %.2f\n", amount)
}

// Withdraw deducts a given amount from the account balance.
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount to withdraw should be positive")
	}

	if a.Balance < amount {
		return errors.New("insufficient balance")
	}

	a.Balance -= amount
	fmt.Printf("Withdrawn: %.2f\n", amount)
	return nil
}

// GetBalance returns the current balance of the account.
func (a *Account) GetBalance() float64 {
	return a.Balance
}
