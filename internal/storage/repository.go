package storage

import "github.com/faroukelkholy/bank/internal/service/models"

type Repository interface {
	AccountRepository
	CustomerRepository
}


type AccountRepository interface {
	GetAccount(id string) (*Account, error)
}

type CustomerRepository interface {
	CreateCustomerAccount(id string, account *models.Account) error
}