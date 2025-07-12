package services

import (
	"github.com/kryast/Crud-3.git/models"
	"github.com/kryast/Crud-3.git/repositories"
)

type CustomerService interface {
	CreateCustomer(customer *models.Customer) error
}

type customerService struct {
	repo repositories.CustomerRepository
}

func NewCustomerService(r repositories.CustomerRepository) CustomerService {
	return &customerService{r}
}

func (s *customerService) CreateCustomer(customer *models.Customer) error {
	return s.repo.Create(customer)
}
