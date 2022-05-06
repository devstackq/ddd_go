package tavern

import (
	"log"

	"github.com/devstackq/tg_bot_ddd/internal/tavern/services/order"
	"github.com/google/uuid"
)

type TavernConfiguration func(os *Tavern) error

// service inside service
type Tavern struct {
	OrderService   *order.OrderService
	BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// WithOrderService applies a given OrderService to the Tavern
func WithOrderService(os *order.OrderService) TavernConfiguration {
	// return a function that matches the TavernConfiguration signature
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customerID uuid.UUID, prodictsIds []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customerID, prodictsIds)
	if err != nil {
		return err
	}
	log.Println("user id ", customerID, price, "order products price equal : ", price)
	return nil
}
