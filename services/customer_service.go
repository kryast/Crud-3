package services

import (
	"github.com/kryast/Crud-3.git/models"
	"github.com/kryast/Crud-3.git/repositories"
)

type CustomerService interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomers() ([]models.Customer, error)
	GetCustomerByID(id uint) (models.Customer, error)
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

func (s *customerService) GetCustomers() ([]models.Customer, error) {
	return s.repo.FindAll()
}

func (s *customerService) GetCustomerByID(id uint) (models.Customer, error) {
	return s.repo.FindByID(id)
}
