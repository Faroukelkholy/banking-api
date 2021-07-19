package mocks

import "github.com/faroukelkholy/bank/internal/storage"

var AccEntity = &storage.Account{
	ID:      "123e4567-e89b-12d3-a456-426614174000",
	Name:    "current",
	Balance: 500,
}
