package memory

import (
	"sync"

	"github.com/devstackq/tg_bot_ddd/internal/tavern/domain/product"
	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (m *MemoryProductRepository) GetAll() ([]product.Product, error) {
	var result []product.Product
	for _, product := range m.products {
		result = append(result, product)
	}
	return result, nil
}

func (m *MemoryProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := m.products[id]; ok {
		return product, nil
	}
	return product.Product{}, product.ErrProductNotFound
}

func (m *MemoryProductRepository) Add(newProduct product.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[newProduct.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	m.products[newProduct.GetID()] = newProduct
	return nil
}

func (m *MemoryProductRepository) Update(pr product.Product) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.products[pr.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	m.products[pr.GetID()] = pr
	return nil
}
