package repositories

import "gorm.io/gorm"

type CustomerRepository interface {
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}
