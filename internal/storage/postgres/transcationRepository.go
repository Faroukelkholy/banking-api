package postgres

import (
	"errors"

	"github.com/go-pg/pg/v10"

	"github.com/faroukelkholy/bank/internal/storage"
)

type TransactionRepository struct {
	DB *pg.DB
}

func NewTransactionRepo(DB *pg.DB) storage.TransactionRepository {
	return &TransactionRepository{DB: DB}
}

func (repo *TransactionRepository) CreateTransaction(t *storage.Transaction) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Close()

	if _, err = tx.Model(&storage.Account{}).Set("balance = balance - ?", t.Amount).Where("id = ?", t.SenderID).Update(); err != nil {
		return err
	}

	if _, err = tx.Model(&storage.Account{}).Set("balance = balance + ?", t.Amount).Where("id = ?", t.ReceiverID).Update(); err != nil {
		return err
	}

	if _, err = tx.Model(t).Insert(); err != nil {
		if IsViolateFK(err.Error()) {
			return errors.New(NoAccountID)
		}
		return err
	}

	if trErr := tx.Commit(); trErr != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			return rollErr
		}
		return trErr
	}

	return err
}

func (repo *TransactionRepository) GetAccountTransactions(accountID string) (ts []*storage.Transaction, err error) {
	if err = repo.DB.Model(&ts).Relation("Sender.Customer").Relation("Receiver.Customer").Where("sender_id = ?", accountID).WhereOr("receiver_id = ?", accountID).Select(); err == pg.ErrNoRows {
		return nil, nil
	}
	return ts, err
}
