package repo

import (
	"fmt"
	"lld/gymchain/entity"
	"sync"
	"time"
)

// BookingRepository manages booking data storage and retrieval
type BookingRepository struct {
	bookings       map[string]*entity.Booking
	bookingCounter int
	mu             sync.RWMutex
}

// NewBookingRepository creates a new booking repository
func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		bookings: make(map[string]*entity.Booking),
	}
}

// CreateBooking creates and stores a new booking
func (r *BookingRepository) CreateBooking(customerID, gymID, classID string) *entity.Booking {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.bookingCounter++
	bookingID := fmt.Sprintf("booking_%d", r.bookingCounter)

	booking := &entity.Booking{
		ID:          bookingID,
		CustomerID:  customerID,
		GymID:       gymID,
		ClassID:     classID,
		BookingTime: time.Now(),
	}

	r.bookings[bookingID] = booking
	return booking
}

// GetBooking retrieves a booking by ID
func (r *BookingRepository) GetBooking(bookingID string) (*entity.Booking, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	booking, exists := r.bookings[bookingID]
	if !exists {
		return nil, entity.ErrBookingNotFound
	}

	return booking, nil
}

// DeleteBooking removes a booking
func (r *BookingRepository) DeleteBooking(bookingID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.bookings[bookingID]; !exists {
		return entity.ErrBookingNotFound
	}

	delete(r.bookings, bookingID)
	return nil
}
