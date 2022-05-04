package entity

import "github.com/google/uuid"

// valid, sanitaze, etc
// mutable object -> use ID

// Item represents a Item for all sub domains
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
	// Cost float32
}
