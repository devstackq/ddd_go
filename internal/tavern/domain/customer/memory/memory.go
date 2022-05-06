package memory

import (
	"sync"

	"github.com/devstackq/tg_bot_ddd/internal/tavern/domain/customer"
	"github.com/google/uuid"
)

// Package memory is a in-memory implementation of the customer repository
// лучше держать каждую реализацию внутри собственной директории

// implement CustomerRepo interface
type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer // 1 customer - have own inof, producs, transactions
	sync.Mutex
}

// factory generate a new Repo of customers
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c customer.Customer) error {
	// check !nil, else init
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
		mr.Unlock()
	}
	// check exist cusmoer ?
	if _, ok := mr.customers[c.GetID()]; ok {
		return customer.ErrFailedToAddCustomer
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) Update(c customer.Customer) error {
	// if !find - error
	if _, ok := mr.customers[c.GetID()]; !ok {
		return customer.ErrUpdateCustomer
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}
