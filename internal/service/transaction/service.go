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
	repo storage.Repository
}

func New(repo storage.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateTransaction(t *models.Transfer) (err error) {
	if err = t.Validate(); err != nil {
		return
	}

	trans := serializeCT(t)
	return s.repo.CreateTransaction(trans)
}

// serializeCT translate transaction data structure from the service to the repository
func serializeCT(t *models.Transfer) *storage.Transaction {
	return &storage.Transaction{
		 SenderID: t.Sender,
		 ReceiverID: t.Receiver,
		 Amount: t.Amount,
	}
}