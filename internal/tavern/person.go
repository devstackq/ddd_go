package tavern

import "github.com/google/uuid"

// mutable object -> use ID; fields Name - UpperCase

// Person is a entity that represents a person in all Domains
type Person struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID   uuid.UUID
	Name string
	Age  int
}
