package handlers

import "github.com/kryast/Crud-3.git/services"

type CustomerHandler struct {
	service services.CustomerService
}

func NewCustomerHandler(s services.CustomerService) *CustomerHandler {
	return &CustomerHandler{s}
}
