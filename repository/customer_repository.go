package repository

import (
	"errors"
	"golang_db_fundamental/model/dto"
	"golang_db_fundamental/model/entity"
	"golang_db_fundamental/utils"

	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	Insert(customer *entity.Customer) error
	Update(customer *entity.Customer) error
	Delete(id string) error
	GetAll(page int, totalRow int) ([]entity.Customer, error)
	GetByID(id string) (entity.Customer, error)
	GetByName(name string) ([]entity.Customer, error)
	GetCount() ([]dto.CustomerCount, error)
	GetSum() (int, error)
	GetCountByAddress() ([]dto.CustomerAddress, error)
}

type customerRepository struct {
	db *sqlx.DB
}

func (c *customerRepository) GetCountByAddress() ([]dto.CustomerAddress, error) {
	var customerCount []dto.CustomerAddress
	if err := c.db.Select(&customerCount, utils.SELECT_ACTIVE_PASSIVE_CUSTOMER); err != nil {
		return customerCount, err
	}
	return customerCount, nil
}

func (c *customerRepository) GetCount() ([]dto.CustomerCount, error) {
	var customerCount []dto.CustomerCount
	if err := c.db.Get(&customerCount, utils.SELECT_COUNT_CUSTOMER); err != nil {
		return customerCount, err
	}
	return customerCount, nil
}

func (c *customerRepository) GetSum() (int, error) {
	var customerBalance int
	if err := c.db.Get(&customerBalance, utils.SELECT_SUM_SALDO_CUSTOMER); err != nil {
		return customerBalance, err
	}
	return customerBalance, nil
}

func (c *customerRepository) GetAll(page int, totalRow int) ([]entity.Customer, error) {
	// rumus pagination
	limit := totalRow
	offset := limit * (page - 1)
	// penggunaan fetch all data: select ($1), named query (:column_name), queryx ()
	customer := []entity.Customer{}
	if err := c.db.Select(&customer, utils.SELECT_ALL_CUSTOMER, limit, offset); err != nil {
		utils.IsError(err)
		return nil, err
	}
	return customer, nil
	// // TODO implement me
	// panic("implement me")
}

func (c *customerRepository) GetByID(id string) (entity.Customer, error) {
	customer := entity.Customer{}
	_, err := c.db.Queryx(id).Get(&customer, utils.SELECT_CUSTOMER_BY_ID, id)
	if err != nil {
		return customer, err
	}
	return customer, nil
	// // TODO implement me
	// panic("implement me")
}

func (c *customerRepository) GetByName(name string) ([]entity.Customer, error) {
	customer := []entity.Customer{}
	if err := c.db.Select(&customer, utils.SELECT_CUSTOMER_BY_NAME, "%"+name+"%"); err != nil {
		utils.IsError(err)
		return nil, err
	}
	return customer, nil
	// // TODO implement me
	// panic("implement me")
}

func (c *customerRepository) Insert(customer *entity.Customer) error {
	_, err := c.db.NamedExec(utils.INSERT_CUSTOMER, customer)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Update(customer *entity.Customer) error {
	_, err := c.db.NamedExec(utils.UPDATE_CUSTOMER, customer)
	err = utils.IsError(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Delete(id string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("delete fail")
		}
	}()
	c.db.MustExec(utils.DELETE_CUSTOMER, id)
	return
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	cstRepo := new(customerRepository)
	cstRepo.db = db
	return cstRepo
}
