package postgres

import (
	"errors"

	"github.com/go-pg/pg/v10"

	"github.com/faroukelkholy/bank/internal/storage"
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

func (repo *CustomerRepository) CreateCustomerAccount(id string, account *storage.Account) error {
	acc := &storage.Account{
		Name:       account.Name,
		Balance:    account.Balance,
		CustomerID: id,
	}

	if _, err := repo.DB.Model(acc).Insert(); err != nil {
		if IsViolateFK(err.Error()) {
			return errors.New(NoCustomerID)
		}
		return err
	}

	return nil
}
