package postgres

import (
	"github.com/go-pg/pg/v10"

	"github.com/faroukelkholy/bank/internal/storage"
)

type AccountRepository struct {
	DB *pg.DB
}

func NewAccountRepo(DB *pg.DB) storage.AccountRepository {
	return &AccountRepository{DB: DB}
}

func (repo *AccountRepository) GetAccount(id string) (as *storage.Account, err error) {
	if err = repo.DB.Model(as).Where("id= ?", id).Select(); err == pg.ErrNoRows {
		return nil, nil
	}
	return as, err
}
