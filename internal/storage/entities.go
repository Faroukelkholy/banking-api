package storage

import "time"

//nolint
type Customer struct {
	tableName struct{} `pg:"customers,alias:c`
	ID        string   `pg:",pk"`
	Name      string   `pg:"name"`
}

//nolint
type Account struct {
	tableName  struct{} `pg:"accounts, alias:a"`
	ID         string   `pg:",pk"`
	Name       string   `pg:"name"`
	Balance    float64  `pg:"balance"`
	CustomerID string
	Customer   *Customer `pg:"rel:has-one,join_fk:id"`
}

//nolint
type Transaction struct {
	tableName  struct{} `pg:"transactions, alias:t"`
	ID         string   `pg:",pk"`
	SenderID   string
	ReceiverID string
	Sender     *Account  `pg:"rel:has-one,join_fk:id"`
	Receiver   *Account  `pg:"rel:has-one,join_fk:id"`
	Amount     float64   `pg:"amount"`
	CreatedAt  time.Time `pg:"created_at"`
}
