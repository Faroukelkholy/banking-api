package account

import (
	"github.com/faroukelkholy/bank/internal/service/models"
	"github.com/faroukelkholy/bank/internal/storage"
)

type Service interface {
	GetAccount(id string) (*models.Account,error)
}

type service struct {
	repo storage.AccountRepository
}

func New(repo storage.AccountRepository) Service {
	return &service{repo: repo}
}

func (s *service) GetAccount(id string) (*models.Account, error) {
	e, err := s.repo.GetAccount(id)
	if e !=nil {
		return serializeGA(e), err
	}
	return nil, err
}

func serializeGA(e *storage.Account) *models.Account{
	return &models.Account{
		Name:         e.Name,
		Balance:      e.Balance,
	}
}
