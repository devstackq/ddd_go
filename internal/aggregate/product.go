package aggregate

import (
	"errors"

	"github.com/devstackq/tg_bot_ddd/internal/entity"
	"github.com/google/uuid"
)

var ErrMissingValues = errors.New("missing values")

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name != "" && description != "" && price > 0 {
		return Product{
			item: &entity.Item{
				ID:          uuid.New(),
				Name:        name,
				Description: description,
			},
			price:    price,
			quantity: 0,
		}, nil
	}

	return Product{}, ErrMissingValues
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}
