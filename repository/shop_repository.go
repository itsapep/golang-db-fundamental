package repository

import (
	"errors"
	"golang_db_fundamental/model/entity"
	"golang_db_fundamental/utils"

	"github.com/jmoiron/sqlx"
)

type ShopRepository interface {
	Insert(shop *entity.Shop) error
	Update(shop *entity.Shop) error
	Delete(id string) error
	GetShopProduct()
}

type shopRepository struct {
	db *sqlx.DB
}

func (c *shopRepository) Insert(shop *entity.Shop) error {
	_, err := c.db.NamedExec(utils.INSERT_SHOP, shop)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *shopRepository) Update(shop *entity.Shop) error {
	_, err := c.db.NamedExec(utils.UPDATE_SHOP, shop)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *shopRepository) Delete(id string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("delete fail")
		}
	}()
	c.db.MustExec(utils.DELETE_SHOP, id)
	return
}

func NewShopRepository(db *sqlx.DB) ShopRepository {
	shopRepo := new(shopRepository)
	shopRepo.db = db
	return shopRepo
}
