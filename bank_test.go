package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {
	tests := []struct {
		amount   int64
		expected int64
		wantErr  bool
	}{
		{amount: 25, expected: 35, wantErr: false},
		{amount: 9223372036854775807, expected: 10, wantErr: true},
		{amount: -9223372036854775808, expected: 10, wantErr: true},
		{amount: 0, expected: 10, wantErr: true},
		{amount: -32768, expected: 10, wantErr: true},
		{amount: 32767, expected: 32777, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatInt(tt.amount, 10), func(t *testing.T) {
			account := &Account{Balance: 10}
			err := account.Deposit(tt.amount)

			assert.Equal(t, tt.expected, account.Balance, "balance after amount %d", tt.amount)
			assert.Equal(t, tt.wantErr, err != nil, "Deposit(%d) error = %v; wantErr = %t", tt.amount, err, tt.wantErr)
		})
	}
}

func TestWithdraw(t *testing.T) {
	tests := []struct {
		amount   int64
		expected int64
		wantErr  bool
	}{
		{amount: 1, expected: 9, wantErr: false},
		{amount: 9223372036854775807, expected: 10, wantErr: true},
		{amount: -9223372036854775808, expected: 10, wantErr: true},
		{amount: 0, expected: 10, wantErr: true},
		{amount: -32768, expected: 10, wantErr: true},
		{amount: 32767, expected: 10, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatInt(tt.amount, 10), func(t *testing.T) {
			account := &Account{Balance: 10}
			err := account.Withdraw(tt.amount)
			assert.Equal(t, tt.expected, account.Balance, "balance after amount %d", tt.amount)
			assert.Equal(t, tt.wantErr, err != nil, "Withdraw(%d) error = %v; wantErr = %t", tt.amount, err, tt.wantErr)
		})
	}
}
