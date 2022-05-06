package main

import (
	"github.com/devstackq/tg_bot_ddd/internal/tavern/domain/product"
	"github.com/devstackq/tg_bot_ddd/internal/tavern/services/order"
	"github.com/devstackq/tg_bot_ddd/internal/tavern/services/tavern"
	"github.com/google/uuid"
)

func main() {
	// todo - company side; transaction; product minus db;

	// customer side - trans; history trans; when , what , amount

	productz := productInventory()
	// Create Order Service to use in tavern
	orderService, err := order.NewOrderService(
		order.WithMongoCustomerRepository("mongodb://localhost:27017"),
		order.WithMemoryProductRepository(productz),
	)
	if err != nil {
		panic(err)
	}
	//  tavern service init
	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(orderService),
	)
	if err != nil {
		panic(err)
	}
	// handler use here?
	uuidCst, err := orderService.AddCustomer("Jussy")
	if err != nil {
		panic(err)
	}
	// customer buy product;
	ordersUuids := []uuid.UUID{
		productz[0].GetID(),
	}
	// execute tavern
	err = tavern.Order(uuidCst, ordersUuids)
	if err != nil {
		panic(err)
	}
}

// mock, added; db - products
func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		panic(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}
