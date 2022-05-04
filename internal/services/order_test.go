package services

import (
	"log"
	"testing"

	"github.com/devstackq/tg_bot_ddd/internal/aggregate"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy Beverage", 2.34)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := aggregate.NewProduct("Wine", "Healthy Snacks", 0.99)
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

	orderSrv, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}
	// add new customer db
	// cust, err := aggregate.NewCustomer("zussy")
	// if err != nil {
	// 	t.Error(err)
	// }
	// orderSrv.customers.Add(cust)

	// cId := cust.GetID()

	order := []uuid.UUID{
		products[0].GetID(),
		products[2].GetID(),
	}

	price, err := orderSrv.CreateOrder(uuid.MustParse("tezt"), order)
	if err != nil {
		t.Error(err)
	}
	log.Print(price, err)
}
