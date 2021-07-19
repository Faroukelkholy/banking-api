package customer

import (
	"github.com/faroukelkholy/bank/internal/service/models"
	"github.com/faroukelkholy/bank/internal/storage"
)

type Service interface {
	CreateCustomerAccount(id string, account *models.Account) error
}

type service struct {
	repo storage.Repository
}

func New(repo storage.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateCustomerAccount(id string, account *models.Account) error {
	acc := serializeCCA(account)
	return s.repo.CreateCustomerAccount(id, acc)
}

// serializeCCA translate transaction data structure from the service to the repository
func serializeCCA(acc *models.Account) *storage.Account {
	return &storage.Account{
		Name: acc.Name,
		Balance: acc.Balance,
	}
}