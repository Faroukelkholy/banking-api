package storage

type Repository interface {
	AccountRepository
	CustomerRepository
	TransactionRepository
}

type AccountRepository interface {
	GetAccount(id string) (*Account, error)
}

type CustomerRepository interface {
	CreateCustomerAccount(id string, account *Account) error
}

type TransactionRepository interface {
	CreateTransaction(t *Transaction) error
	GetAccountTransactions(accountID string) ([]*Transaction, error)
}
