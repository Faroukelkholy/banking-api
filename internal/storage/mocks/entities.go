package mocks

import (
	"github.com/faroukelkholy/bank/internal/storage"
	"time"
)

var Account = &storage.Account{
	ID:      "123e4567-e89b-12d3-a456-426614174000",
	Name:    "current",
	Balance: 500,
	Customer: &storage.Customer{
		ID:   "123e4567-e89b-12d3-a456-426614174000",
		Name: "Arisha Barron",
	},
}

var Transactions = []*storage.Transaction{
	{
		Sender:    Account,
		Receiver:  Account,
		Amount:    500,
		CreatedAt: time.Time{},
	},
}