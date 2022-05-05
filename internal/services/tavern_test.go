package services

import (
	"log"
	"testing"

	"github.com/devstackq/tg_bot_ddd/internal/aggregate"
	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	orderService, err := NewOrderService(
		WithMemoryProductRepository(products),
		WithMemoryCustomerRepository(),
	)
	if err != nil {
		t.Error(err)
	}
	cust, err := aggregate.NewCustomer("Ushq")
	if err != nil {
		t.Error(err)
	}

	err = orderService.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}
	productIds := []uuid.UUID{
		products[2].GetID(), // berr
		products[1].GetID(), // peanuts
	}
	tavern, err := NewTavern(
		WithOrderService(orderService),
	)
	if err != nil {
		t.Error(err)
	}

	err = tavern.Order(cust.GetID(), productIds)
	log.Println(cust.GetID(), "uid")

	if err != nil {
		t.Error(err)
	}

	log.Println(err, " result")
}
