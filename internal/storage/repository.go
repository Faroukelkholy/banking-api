package storage

import "github.com/faroukelkholy/bank/internal/service/models"

type Repository interface {
	AccountRepository
	CustomerRepository
	TransactionRepository
}

type AccountRepository interface {
	GetAccount(id string) (*Account, error)
}

type CustomerRepository interface {
	CreateCustomerAccount(id string, account *models.Account) error
}

type TransactionRepository interface {
	CreateTransaction(t *models.Transfer) error
	GetAccountTransactions(accountID string) ([]*Transaction, error)
}