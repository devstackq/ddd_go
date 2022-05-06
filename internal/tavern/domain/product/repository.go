package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrProductNotFound is returned when a product is not found
	ErrProductNotFound = errors.New("the product was not found")
	// ErrProductAlreadyExist is returned when trying to add a product that already exists
	ErrProductAlreadyExist = errors.New("the product already exists")
)

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Add(Product) error
	Update(Product) error
	// Delete(id uuid.UUID) error
}
