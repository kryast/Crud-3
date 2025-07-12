package repositories

import (
	"github.com/kryast/Crud-3.git/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *models.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) Create(customer *models.Customer) error {
	return r.db.Create(customer).Error
}
