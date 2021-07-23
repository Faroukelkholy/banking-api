package mocks

import "github.com/faroukelkholy/bank/internal/service/models"

var Transfer = &models.Transfer{
	Sender:   "123e4567-e89b-12d3-a456-426614174000",
	Receiver: "123e4567-e89b-12d3-a456-426614174001",
	Amount:   400,
}
