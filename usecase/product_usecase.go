package usecase

import (
	"golang_db_fundamental/model/entity"
	"golang_db_fundamental/repository"
)

type ProductUseCase interface {
	InsertProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id string) error
}

type productUseCase struct {
	repo repository.ProductRepository
}

func (c *productUseCase) InsertProduct(product *entity.Product) error {
	return c.repo.Insert(product)
}

func (c *productUseCase) UpdateProduct(product *entity.Product) error {
	return c.repo.Update(product)
}

func (c *productUseCase) DeleteProduct(id string) error {
	return c.repo.Delete(id)
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	usc := new(productUseCase)
	usc.repo = repo
	return usc
}
