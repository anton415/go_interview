package main

import (
	"fmt"
	"math"
)

type Account struct {
    Balance int64
}

func (a *Account) Deposit(amount int64) error {
	if a == nil {
		return fmt.Errorf("Deposit: аккаунт равен nil")
	}
	if amount <= 0 {
		return fmt.Errorf("Deposit: сумма должна быть положительной")
	}
	if a.Balance > math.MaxInt64-amount {
		return fmt.Errorf("Deposit: превышение лимита баланса")
	}
    a.Balance += amount
    return nil
}

func (a *Account) Withdraw(amount int64) error {
	if a == nil {
		return fmt.Errorf("Withdraw: аккаунт равен nil")
	}
	if amount <= 0 {
		return fmt.Errorf("Withdraw: сумма должна быть положительной")
	}
	if amount > a.Balance {
		return fmt.Errorf("Withdraw: недостаточный баланс")
	}
	a.Balance -= amount
	return nil
}

