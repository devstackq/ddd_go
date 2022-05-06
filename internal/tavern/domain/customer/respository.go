package customer

// Package Customer holds all the domain logic for the customer domain.

import (
	"errors"

	"github.com/google/uuid"
)

// ошибки определены именно в пакете customer, а не в пакете репозитория. Это позволит не зависеть от реализации

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer      = errors.New("failed to update the customer in the repository")
)

// CustomerRepository is a interface that defines the rules around what a customer
type CustomerRepository interface {
	Get(uuid.UUID) (Customer, error)
	Add(Customer) error
	Update(Customer) error
}
