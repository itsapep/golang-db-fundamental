package repository

import (
	"errors"
	"golang_db_fundamental/model/entity"
	"golang_db_fundamental/utils"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Insert(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id string) error
}

type productRepository struct {
	db *sqlx.DB
}

func (c *productRepository) Insert(product *entity.Product) error {
	_, err := c.db.NamedExec(utils.INSERT_PRODUCT, product)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *productRepository) Update(product *entity.Product) error {
	_, err := c.db.NamedExec(utils.UPDATE_PRODUCT, product)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *productRepository) Delete(id string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("delete fail")
		}
	}()
	c.db.MustExec(utils.DELETE_PRODUCT, id)
	return
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	productRepo := new(productRepository)
	productRepo.db = db
	return productRepo
}
