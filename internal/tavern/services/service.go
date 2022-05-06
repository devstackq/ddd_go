package services

import (
	"github.com/google/uuid"
)

type Order interface {
	CreateOrder(customerId uuid.UUID, productIds []uuid.UUID) (float64, error)
}

type Tavern interface {
	Order(customerID uuid.UUID, prodictsIds []uuid.UUID) error
}
