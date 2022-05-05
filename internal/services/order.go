package services

import (
	"log"

	"github.com/devstackq/tg_bot_ddd/internal/aggregate"
	"github.com/devstackq/tg_bot_ddd/internal/domain/customer"
	customMemory "github.com/devstackq/tg_bot_ddd/internal/domain/customer/memory"
	productMemory "github.com/devstackq/tg_bot_ddd/internal/domain/product/memory"

	"github.com/devstackq/tg_bot_ddd/internal/domain/product"
	"github.com/google/uuid"
)

type OrderConfigurations func(os *OrderService) error

// order -> userId
type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
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

func WithMemoryCustomerRepository() OrderConfigurations {
	memoryCustomRepo := customMemory.New()
	return WithCustomerRepository(memoryCustomRepo)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfigurations {
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
		log.Print(err, "her")
		return 0, err
	}
	var products []aggregate.Product
	// get db product by product ID, add to customer/ calcualte price
	for _, id := range productIds {
		// aggregate.Product.GetID()
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
