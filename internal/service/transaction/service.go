package transaction

import (
	"github.com/faroukelkholy/bank/internal/service/models"
	"github.com/faroukelkholy/bank/internal/storage"
)

type Service interface {
	CreateTransaction(t *models.Transfer) error
}

//service struct implement the Service interface
type service struct {
	repo storage.TransactionRepository
}

func New(repo storage.TransactionRepository) Service {
	return &service{repo: repo}
}


func (s *service) CreateTransaction(t *models.Transfer) (err error) {
	if err = t.Validate(); err != nil {
		return err
	}
	return s.repo.CreateTransaction(t)
}

