package tavern

// import (
// 	"log"
// 	"testing"

// 	"github.com/devstackq/tg_bot_ddd/internal/tavern/domain/customer"
// 	"github.com/devstackq/tg_bot_ddd/internal/tavern/services/order"

// 	"github.com/google/uuid"
// )

// func Test_Tavern(t *testing.T) {
// 	products := init_products(t)

// 	orderService, err := order.NewOrderService(
// 		WithMemoryProductRepository(products),
// 		WithMemoryCustomerRepository(),
// 	)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	cust, err := customer.NewCustomer("Ushq")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	err = orderService.customers.Add(cust)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	productIds := []uuid.UUID{
// 		products[2].GetID(), // berr
// 		products[1].GetID(), // peanuts
// 	}
// 	tavern, err := NewTavern(
// 		WithOrderService(orderService),
// 	)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	err = tavern.Order(cust.GetID(), productIds)
// 	log.Println(cust.GetID(), "uid")

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	log.Println(err, " result")
// }

// func Test_MongoTavern(t *testing.T) {
// 	products := init_products(t)
// 	// can use, two other db
// 	orderService, err := NewOrderService(
// 		WithMemoryProductRepository(products),
// 		WithMongoCustomerRepository("mongodb://localhost:27017"),
// 	)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	cust, err := customer.NewCustomer("Ushq")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	err = orderService.customers.Add(cust)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	productIds := []uuid.UUID{
// 		products[2].GetID(), // berr
// 		products[1].GetID(), // peanuts
// 	}
// 	tavern, err := NewTavern(
// 		WithOrderService(orderService),
// 	)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	err = tavern.Order(cust.GetID(), productIds)
// 	// log.Println(cust.GetID(), "uid")

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	log.Println(err, " result")
// }
