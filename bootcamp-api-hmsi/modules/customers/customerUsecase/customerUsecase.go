package customerUsecase

import (
	"github.com/kisahtegar/bootcamp-api-hmsi/models"
	"github.com/kisahtegar/bootcamp-api-hmsi/modules/customers"
)

type customerRepository struct {
	Repo customers.CustomerRepository
}

func NewCustomerUsecase(Repo customers.CustomerRepository) customers.CustomerUsecase {
	return &customerRepository{Repo}
}

// Implement FindAll customer for usecase.
func (r *customerRepository) FindAll() (*[]models.Customers, error) {
	results, err := r.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return results, nil
}

// Implement Insert customer for usecase.
func (r *customerRepository) Insert(c *models.RequestInsertCustomer) error {
	err := r.Repo.Create(c)
	if err != nil {
		return err
	}
	return nil
}

// Implement Update customer for usecase.
func (r *customerRepository) Update(c *models.RequestUpdateCustomer) error {
	err := r.Repo.Update(c)
	if err != nil {
		return err
	}
	return nil
}

// Implement Delete customer for usecase.
func (r *customerRepository) Delete(Id uint64) error {
	err := r.Repo.Delete(Id)
	if err != nil {
		return err
	}
	return nil
}
