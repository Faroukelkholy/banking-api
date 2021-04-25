package models

type Transfer struct {
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount float64 `json:"amount"`
}

