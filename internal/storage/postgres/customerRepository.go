package postgres

import (
	"github.com/faroukelkholy/bank/internal/service/models"
	"github.com/faroukelkholy/bank/internal/storage"
	"github.com/go-pg/pg/v10"
)

type CustomerRepository struct {
	DB *pg.DB
}

func NewCustomerRepo(DB *pg.DB) storage.CustomerRepository {
	return &CustomerRepository{DB: DB}
}

func (repo *CustomerRepository) GetCustomers() (cs []*storage.Customer, err error) {
	err = repo.DB.Model(&cs).Select()
	return cs, err
}

func (repo *CustomerRepository) CreateCustomerAccount(id string, account *models.Account) error {
	ca := &storage.Account{
		Name:       account.Name,
		Balance:    account.Balance,
		CustomerID: id,
	}
	_, err := repo.DB.Model(ca).Insert()
	return err
}

