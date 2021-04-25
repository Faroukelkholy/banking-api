package models

import "time"

type Transaction struct {
	Sender *Customer `json:"sender"`
	Receiver *Customer `json:"receiver"`
	Amount float64 `json:"amount"`
	CreatedAt time.Time `json:"create_at"`
}

