package account

import (
	"github.com/faroukelkholy/bank/internal/service/models"
	"github.com/faroukelkholy/bank/internal/storage"
)

//Service define services related to account domain
type Service interface {
	GetAccount(id string) (*models.Account,error)
}

//service struct implement the Service interface
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

//serializeGA translate account data structure from the repository to the service
func serializeGA(e *storage.Account) *models.Account{
	return &models.Account{
		Name:         e.Name,
		Balance:      e.Balance,
	}
}
