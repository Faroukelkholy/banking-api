package customer

import (
	"github.com/faroukelkholy/bank/internal/service/models"
	"github.com/faroukelkholy/bank/internal/storage"
)

type Service interface {
	CreateCustomerAccount(id string, account *models.Account) error
}

type service struct {
	repo storage.CustomerRepository
}

func New(repo storage.CustomerRepository) Service {
	return &service{repo: repo}
}

func (s *service) CreateCustomerAccount(id string, account *models.Account) error {
	return s.repo.CreateCustomerAccount(id, account)
}

