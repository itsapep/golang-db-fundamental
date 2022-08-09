package repository

import (
	"errors"
	"golang_db_fundamental/model/entity"
	"golang_db_fundamental/utils"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Insert(transaction *entity.Transaction) error
	Update(transaction *entity.Transaction) error
	Delete(id string) error
}

type transactionRepository struct {
	db *sqlx.DB
}

func (c *transactionRepository) Insert(transaction *entity.Transaction) error {
	_, err := c.db.NamedExec(utils.INSERT_TRANSACTION, transaction)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *transactionRepository) Update(transaction *entity.Transaction) error {
	_, err := c.db.NamedExec(utils.UPDATE_TRANSACTION, transaction)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *transactionRepository) Delete(id string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("delete fail")
		}
	}()
	c.db.MustExec(utils.DELETE_TRANSACTION, id)
	return
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	transactionRepo := new(transactionRepository)
	transactionRepo.db = db
	return transactionRepo
}
