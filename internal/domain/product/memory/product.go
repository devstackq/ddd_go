package memory

import (
	"sync"

	"github.com/devstackq/tg_bot_ddd/internal/aggregate"
	"github.com/devstackq/tg_bot_ddd/internal/domain/product"
	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (m *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var result []aggregate.Product
	for _, product := range m.products {
		result = append(result, product)
	}
	return result, nil
}

func (m *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := m.products[id]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

func (m *MemoryProductRepository) Add(pr aggregate.Product) error {
	if m.products == nil {
		m.Lock()
		m.products = make(map[uuid.UUID]aggregate.Product)
		m.Unlock()
	}

	if _, ok := m.products[pr.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	m.Lock()
	m.products[pr.GetID()] = pr
	m.Unlock()
	return nil
}

func (m *MemoryProductRepository) Update(pr aggregate.Product) error {
	if _, ok := m.products[pr.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	m.Lock()
	m.products[pr.GetID()] = pr
	m.Unlock()
	return nil
}
