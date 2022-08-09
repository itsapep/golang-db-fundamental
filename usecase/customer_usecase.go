package usecase

import (
	"golang_db_fundamental/model/dto"
	"golang_db_fundamental/model/entity"
	"golang_db_fundamental/repository"
)

type CustomerUseCase interface {
	RegisterNewCustomer(customer *entity.Customer) error
	UpdateCustomer(customer *entity.Customer) error
	DeleteCustomer(id string) error
	FindAllCustomer(page int, totalRow int) ([]entity.Customer, error)
	FindCustomerByID(id string) (entity.Customer, error)
	FindCustomerByName(name string) ([]entity.Customer, error)
	GetTotalActiveCustomer() ([]dto.CustomerCount, error)
	GetTotalBalanceActiveCustomer() (int, error)
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (c *customerUseCase) GetTotalActiveCustomer() ([]dto.CustomerCount, error) {
	return c.repo.GetCount()
}

func (c *customerUseCase) GetTotalBalanceActiveCustomer() (int, error) {
	return c.repo.GetSum()
}

func (c *customerUseCase) FindAllCustomer(page int, totalRow int) ([]entity.Customer, error) {
	return c.repo.GetAll(page, totalRow)
}
func (c *customerUseCase) FindCustomerByID(id string) (entity.Customer, error) {
	return c.repo.GetByID(id)
}
func (c *customerUseCase) FindCustomerByName(name string) ([]entity.Customer, error) {
	return c.repo.GetByName(name)
}

func (c *customerUseCase) RegisterNewCustomer(customer *entity.Customer) error {
	return c.repo.Insert(customer)
}

func (c *customerUseCase) UpdateCustomer(customer *entity.Customer) error {
	return c.repo.Update(customer)
}

func (c *customerUseCase) DeleteCustomer(id string) error {
	return c.repo.Delete(id)
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	usc := new(customerUseCase)
	usc.repo = repo
	return usc
}
