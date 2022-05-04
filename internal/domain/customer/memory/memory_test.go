package memory

import (
	"testing"

	"github.com/devstackq/tg_bot_ddd/internal/aggregate"
	"github.com/devstackq/tg_bot_ddd/internal/domain/customer"
	"github.com/google/uuid"
)

// Сейчас репозиторий управляет только одним агрегатом Customer, так и должно быть. сохранять слабую связность (loose couplin

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name          string
		id            uuid.UUID
		expectedError error
	}
	cust, err := aggregate.NewCustomer("Poly")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()

	repo := New()

	repo.Add(cust)

	testCases := []testCase{
		{
			name:          "no customer by id",
			id:            uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedError: customer.ErrCustomerNotFound,
		},
		{
			name:          "ok",
			id:            id,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Customer",
			cust:        "Percy",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				customers: map[uuid.UUID]aggregate.Customer{},
			}

			cust, err := aggregate.NewCustomer(tc.cust)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(cust)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}
}
