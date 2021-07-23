package mocks

import (
	"github.com/faroukelkholy/bank/internal/service/models"
	"time"
)

var Account = &models.Account{
	ID:      "123e4567-e89b-12d3-a456-426614174000",
	Name:    "current",
	Balance: 500,
}

var Transactions = []*models.Transaction{
	{
		Sender:    &models.Customer{},
		Receiver:  &models.Customer{},
		Amount:    500,
		CreatedAt: time.Time{},
	},
}