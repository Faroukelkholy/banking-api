package models


type Account struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name"`
	Balance float64 `json:"balance"`
	Transactions []*Transaction `json:"transactions,omitempty"`
}
