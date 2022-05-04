package product

import (
	"errors"

	"github.com/devstackq/tg_bot_ddd/internal/aggregate"
	"github.com/google/uuid"
)

var (
	// ErrProductNotFound is returned when a product is not found
	ErrProductNotFound = errors.New("the product was not found")
	// ErrProductAlreadyExist is returned when trying to add a product that already exists
	ErrProductAlreadyExist = errors.New("the product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(aggregate.Product) error
	Update(aggregate.Product) error
	// Delete(id uuid.UUID) error
}
