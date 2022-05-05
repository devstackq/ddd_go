package services

import (
	"github.com/devstackq/tg_bot_ddd/internal/aggregate"
	"github.com/google/uuid"
)

type Order interface {
	CreateOrder(customerId uuid.UUID, productIds []uuid.UUID) (float64, error)
}
type Customer interface {
	CreateCustomer(aggregate.Customer) error
}
