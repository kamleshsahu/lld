package repo

import (
	"lld/gymchain/entity"
	"sync"
)

// CustomerRepository manages customer data storage and retrieval
type CustomerRepository struct {
	customers map[string]*entity.Customer
	mu        sync.RWMutex
}

// NewCustomerRepository creates a new customer repository
func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		customers: make(map[string]*entity.Customer),
	}
}

// GetCustomer retrieves a customer by ID
func (r *CustomerRepository) GetCustomer(customerID string) (*entity.Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	customer, exists := r.customers[customerID]
	if !exists {
		return nil, entity.ErrCustomerNotFound
	}

	return customer, nil
}

// CreateCustomer creates a new customer if not exists
func (r *CustomerRepository) CreateCustomer(customerID string) *entity.Customer {
	r.mu.Lock()
	defer r.mu.Unlock()

	customer, exists := r.customers[customerID]
	if !exists {
		customer = &entity.Customer{
			ID:       customerID,
			Bookings: make(map[string]*entity.Booking),
		}
		r.customers[customerID] = customer
	}

	return customer
}
