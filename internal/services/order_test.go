package services

import (
	"log"
	"testing"

	"github.com/devstackq/tg_bot_ddd/internal/aggregate"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer1", "Healthy Beverage1", 2.34)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts1", "Healthy Snacks1", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := aggregate.NewProduct("Wine1", "Healthy Snacks1", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []aggregate.Product{
		beer, peenuts, wine,
	}
	return products
}

func TestOrder_NewOrderService(t *testing.T) {
	// Create a few products to insert into in memory repo
	products := init_products(t)

	// init service; add products fake db
	// add new customer db
	cust, err := aggregate.NewCustomer("user1")
	if err != nil {
		t.Error(err)
	}

	orderSrv, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	orderSrv.customers.Add(cust)

	cId := cust.GetID()

	ordersIds := []uuid.UUID{
		products[0].GetID(),
		products[2].GetID(),
	}

	price, err := orderSrv.CreateOrder(cId, ordersIds)
	if err != nil {
		t.Error(err)
	}
	log.Print(price, err, ordersIds, cust.GetID())
}
