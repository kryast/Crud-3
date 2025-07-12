package repositories

import (
	"github.com/kryast/Crud-3.git/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *models.Customer) error
	FindAll() ([]models.Customer, error)
	FindByID(id uint) (models.Customer, error)
	Update(customer *models.Customer) error
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

func (r *customerRepository) FindAll() ([]models.Customer, error) {
	var customers []models.Customer
	return customers, r.db.Find(&customers).Error
}

func (r *customerRepository) FindByID(id uint) (models.Customer, error) {
	var customer models.Customer
	return customer, r.db.First(&customer, id).Error
}

func (r *customerRepository) Update(customer *models.Customer) error {
	return r.db.Save(customer).Error
}
