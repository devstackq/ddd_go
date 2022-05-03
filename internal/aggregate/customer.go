package aggregate

import (
	"errors"

	"github.com/devstackq/tg_bot_ddd/internal/entity"
	"github.com/devstackq/tg_bot_ddd/internal/valueObject"

	"github.com/google/uuid"
)

var ErrInvalidPerson = errors.New("a customer has to have an valid person")

// inside have another entities; lowCwerCase letter; no access data
// no tags; for serialization; || save db

type Customer struct {
	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	person       *entity.Person // incapsulation
	products     []*entity.Item
	transactions []*valueObject.Transaction // with pointers - for mutable object by pointer
}

func NewCustomer(name string) (Customer, error) {
	// can - sanitaze, validation
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	// create a new persone and generate id
	person := &entity.Person{ID: uuid.New(), Name: name}
	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]*valueObject.Transaction, 0),
	}, nil
}

// getter
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// setter

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) string {
	return c.person.Name
}
