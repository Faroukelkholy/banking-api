package account

import (
	"fmt"

	"github.com/faroukelkholy/bank/internal/service/models"
	"github.com/faroukelkholy/bank/internal/storage"
)

//Service define services related to account domain
type Service interface {
	GetAccount(id string) (*models.Account, error)
	GetTransactions(id string) ([]*models.Transaction, error)
}

//service struct implement the Service interface
type service struct {
	repo   storage.AccountRepository
	trRepo storage.TransactionRepository
}

func New(repo storage.AccountRepository, trRepo storage.TransactionRepository) Service {
	return &service{repo: repo, trRepo: trRepo}
}

func (s *service) GetAccount(id string) (*models.Account, error) {
	e, err := s.repo.GetAccount(id)
	if e != nil {
		return serializeGA(e), err
	}
	return nil, err
}

func (s *service) GetTransactions(id string) ([]*models.Transaction, error) {
	e, err := s.trRepo.GetAccountTransactions(id)
	fmt.Println("e :",e[0].Sender.Customer)
	if e != nil {
		return serializeGT(e), err
	}
	return nil, err
}

//serializeGA translate account data structure from the repository to the service
func serializeGA(e *storage.Account) *models.Account {
	return &models.Account{
		Name:    e.Name,
		Balance: e.Balance,
	}
}

func serializeGT(e []*storage.Transaction) (m []*models.Transaction) {
	for _, tr := range e {
		trans := &models.Transaction{
			Sender: &models.Customer{
				ID:   tr.Sender.Customer.ID,
				Name: tr.Sender.Customer.Name,
			},
			Receiver: &models.Customer{
				ID:   tr.Receiver.Customer.ID,
				Name: tr.Receiver.Customer.Name,
			},
			Amount: tr.Amount,
		}
		m = append(m, trans)
	}
	return m
}
