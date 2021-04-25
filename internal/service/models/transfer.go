package models

import "errors"

const (
	InvalidAmount = "amount must be postive"
)


type Transfer struct {
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount float64 `json:"amount"`
}


func(t *Transfer) Validate() (err error){
	return validateAmount(t.Amount)
}


func validateAmount(amount float64)error{
	if amount <= float64(0) {
		return errors.New(InvalidAmount)
	}
	return nil
}