package order

import (
	"context"
	"log"

	"github.com/devstackq/tg_bot_ddd/internal/tavern/domain/customer"
	"github.com/devstackq/tg_bot_ddd/internal/tavern/domain/product"

	customMemory "github.com/devstackq/tg_bot_ddd/internal/tavern/domain/customer/memory"
	mongoCst "github.com/devstackq/tg_bot_ddd/internal/tavern/domain/customer/mongo"
	productMemory "github.com/devstackq/tg_bot_ddd/internal/tavern/domain/product/memory"

	"github.com/google/uuid"
)

type OrderConfigurations func(os *OrderService) error

// order -> userId
type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
	// trans transaction.TransactionRepository
	// connMongo string
}

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfigurations) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return &OrderService{}, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfigurations {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// func WithMongoCustomerRepo(){}
func WithMongoCustomerRepository(conn string) OrderConfigurations {
	return func(os *OrderService) error {
		ctx := context.Background()

		mongoCstRepo, err := mongoCst.New(ctx, conn)
		if err != nil {
			return err
		}
		log.Println(mongoCstRepo, "repo")
		os.customers = mongoCstRepo
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfigurations {
	memoryCustomRepo := customMemory.New()
	return WithCustomerRepository(memoryCustomRepo)
}

func WithMemoryProductRepository(products []product.Product) OrderConfigurations {
	return func(os *OrderService) error {
		pr := productMemory.New()
		// add product to memory
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (os *OrderService) CreateOrder(customerId uuid.UUID, productIds []uuid.UUID) (float64, error) {
	// get customer from repo
	var price float64

	customerq, err := os.customers.Get(customerId)
	if err != nil {
		return 0, err
	}
	var products []product.Product
	// get db product by product ID, add to customer/ calcualte price
	for _, id := range productIds {
		pr, err := os.products.GetByID(id)
		if err != nil {
			return 0, err
		}

		products = append(products, pr)
		price += pr.GetPrice()
	}
	// then remove - product by id -> memoryProduct

	log.Printf("Customer: %s has ordered %d products at sum %f", customerq.GetID(), len(products), price)

	return price, nil
}

func (os *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	// aggregate
	cst, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}
	// add repo
	err = os.customers.Add(cst)
	if err != nil {
		return uuid.Nil, err
	}
	return cst.GetID(), nil
}
