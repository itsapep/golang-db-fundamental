package usecase

import (
	"golang_db_fundamental/model/entity"
	"golang_db_fundamental/repository"
)

type TransactionUseCase interface {
	InsertTransaction(transaction *entity.Transaction) error
	UpdateTransaction(transaction *entity.Transaction) error
	DeleteTransaction(id string) error
}

type transactionUseCase struct {
	repo repository.TransactionRepository
}

func (c *transactionUseCase) InsertTransaction(transaction *entity.Transaction) error {
	return c.repo.Insert(transaction)
}

func (c *transactionUseCase) UpdateTransaction(transaction *entity.Transaction) error {
	return c.repo.Update(transaction)
}

func (c *transactionUseCase) DeleteTransaction(id string) error {
	return c.repo.Delete(id)
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	ssc := new(transactionUseCase)
	ssc.repo = repo
	return ssc
}
