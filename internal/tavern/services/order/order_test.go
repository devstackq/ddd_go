package order

import (
	"log"
	"testing"

	"github.com/devstackq/tg_bot_ddd/internal/tavern/domain/customer"
	"github.com/devstackq/tg_bot_ddd/internal/tavern/domain/product"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer1", "Healthy Beverage1", 2.34)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := product.NewProduct("Peenuts1", "Healthy Snacks1", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine1", "Healthy Snacks1", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}

func TestOrder_NewOrderService(t *testing.T) {
	// Create a few products to insert into in memory repo
	products := init_products(t)

	// init service; add products fake db
	// add new customer db
	cust, err := customer.NewCustomer("user1")
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
