package usecase

import (
	"golang_db_fundamental/model/entity"
	"golang_db_fundamental/repository"
)

type ShopUseCase interface {
	InsertShop(shop *entity.Shop) error
	UpdateShop(shop *entity.Shop) error
	DeleteShop(id string) error
}

type shopUseCase struct {
	repo repository.ShopRepository
}

func (c *shopUseCase) InsertShop(shop *entity.Shop) error {
	return c.repo.Insert(shop)
}

func (c *shopUseCase) UpdateShop(shop *entity.Shop) error {
	return c.repo.Update(shop)
}

func (c *shopUseCase) DeleteShop(id string) error {
	return c.repo.Delete(id)
}

func NewShopUseCase(repo repository.ShopRepository) ShopUseCase {
	ssc := new(shopUseCase)
	ssc.repo = repo
	return ssc
}
