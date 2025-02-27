package entity

import (
	"errors"
	"sync"
	"time"
)

// Define time format for parsing
const timeFormat = "15:04"

// Define error messages
var (
	ErrGymNotFound          = errors.New("gym not found")
	ErrClassNotFound        = errors.New("class not found")
	ErrCustomerNotFound     = errors.New("customer not found")
	ErrBookingNotFound      = errors.New("booking not found")
	ErrClassFull            = errors.New("class is full")
	ErrTimeConflict         = errors.New("time conflict with existing class")
	ErrInvalidTime          = errors.New("invalid time format or range")
	ErrAlreadyBooked        = errors.New("customer already booked this class")
	ErrCapacityExceeded     = errors.New("gym capacity exceeded")
	ErrInvalidTimeRange     = errors.New("classes must be between 6:00 and 20:00")
	ErrInvalidMaxLimit      = errors.New("max limit must be positive")
	ErrInvalidAccommodation = errors.New("max accommodation must be positive")
)

// Customer represents a gym customer
type Customer struct {
	ID       string
	Bookings map[string]*Booking
	mu       sync.RWMutex
}
type Gym struct {
	ID               string
	Name             string
	Location         string
	MaxAccommodation int
	Classes          map[string]*GymClass
	Mu               sync.RWMutex
}

// AddClass adds a class to the gym
func (g *Gym) AddClass(classID, classType string, maxLimit int, timeSlot TimeSlot) error {
	g.Mu.Lock()
	defer g.Mu.Unlock()

	// Check max limit against gym capacity
	if maxLimit > g.MaxAccommodation {
		return ErrCapacityExceeded
	}

	// Check for time conflicts with existing classes
	for _, class := range g.Classes {
		if class.HasTimeConflict(timeSlot.StartTime, timeSlot.EndTime) {
			return ErrTimeConflict
		}
	}

	// Create and add the class
	g.Classes[classID] = &GymClass{
		ID:              classID,
		GymID:           g.ID,
		Type:            classType,
		MaxLimit:        maxLimit,
		CurrentBookings: 0,
		TimeSlot:        timeSlot,
		Bookings:        make(map[string]*Booking),
	}

	return nil
}

// RemoveClass removes a class from the gym
func (g *Gym) RemoveClass(classID string) error {
	g.Mu.Lock()
	defer g.Mu.Unlock()

	if _, exists := g.Classes[classID]; !exists {
		return ErrClassNotFound
	}

	delete(g.Classes, classID)
	return nil
}

// GetClass retrieves a class by ID
func (g *Gym) GetClass(classID string) (*GymClass, error) {
	g.Mu.RLock()
	defer g.Mu.RUnlock()

	class, exists := g.Classes[classID]
	if !exists {
		return nil, ErrClassNotFound
	}

	return class, nil
}

// GetDetails returns the details of the gym
func (g *Gym) GetDetails() map[string]interface{} {
	g.Mu.RLock()
	defer g.Mu.RUnlock()

	return map[string]interface{}{
		"id":               g.ID,
		"name":             g.Name,
		"location":         g.Location,
		"maxAccommodation": g.MaxAccommodation,
		"classCount":       len(g.Classes),
	}
}

// GymClass represents a workout class within a gym
type GymClass struct {
	ID              string
	GymID           string
	Type            string
	MaxLimit        int
	CurrentBookings int
	TimeSlot        TimeSlot
	Bookings        map[string]*Booking
	mu              sync.RWMutex
}

// AddBooking adds a booking to the class
func (gc *GymClass) AddBooking(bookingID string, booking *Booking) error {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	// Check if class is full
	if gc.CurrentBookings >= gc.MaxLimit {
		return ErrClassFull
	}

	// Check if customer already booked this class
	for _, existingBooking := range gc.Bookings {
		if existingBooking.CustomerID == booking.CustomerID {
			return ErrAlreadyBooked
		}
	}

	gc.Bookings[bookingID] = booking
	gc.CurrentBookings++
	return nil
}

// RemoveBooking removes a booking from the class
func (gc *GymClass) RemoveBooking(bookingID string) error {
	gc.mu.Lock()
	defer gc.mu.Unlock()

	if _, exists := gc.Bookings[bookingID]; !exists {
		return ErrBookingNotFound
	}

	delete(gc.Bookings, bookingID)
	gc.CurrentBookings--
	return nil
}

// HasTimeConflict checks if there's a time conflict with this class
func (gc *GymClass) HasTimeConflict(newStart, newEnd time.Time) bool {
	gc.mu.RLock()
	defer gc.mu.RUnlock()

	existingStart := gc.TimeSlot.StartTime
	existingEnd := gc.TimeSlot.EndTime

	// Check for overlap
	return (newStart.Before(existingEnd) || newStart.Equal(existingEnd)) &&
		(existingStart.Before(newEnd) || existingStart.Equal(newEnd))
}

// GetDetails returns the details of the class
func (gc *GymClass) GetDetails() map[string]interface{} {
	gc.mu.RLock()
	defer gc.mu.RUnlock()

	return map[string]interface{}{
		"id":              gc.ID,
		"type":            gc.Type,
		"maxLimit":        gc.MaxLimit,
		"currentBookings": gc.CurrentBookings,
		"startTime":       gc.TimeSlot.StartTime.Format(timeFormat),
		"endTime":         gc.TimeSlot.EndTime.Format(timeFormat),
	}
}

// GetBookings returns all bookings for this class
func (gc *GymClass) GetBookings() []*Booking {
	gc.mu.RLock()
	defer gc.mu.RUnlock()

	bookings := make([]*Booking, 0, len(gc.Bookings))
	for _, booking := range gc.Bookings {
		bookings = append(bookings, booking)
	}

	return bookings
}

type Booking struct {
	ID          string
	CustomerID  string
	GymID       string
	ClassID     string
	BookingTime time.Time
}

// TimeSlot represents a time slot for a class
type TimeSlot struct {
	StartTime time.Time
	EndTime   time.Time
}

// Booking represents a customer booking

// Gym represents a gym location

// ===================== Service Layer =======================

// Helper function to parse and validate time
func ParseAndValidateTime(startTimeStr, endTimeStr string) (TimeSlot, error) {
	startTime, err := time.Parse(timeFormat, startTimeStr)
	if err != nil {
		return TimeSlot{}, ErrInvalidTime
	}

	endTime, err := time.Parse(timeFormat, endTimeStr)
	if err != nil {
		return TimeSlot{}, ErrInvalidTime
	}

	// Validate time range (6:00-20:00)
	openTime, _ := time.Parse(timeFormat, "06:00")
	closeTime, _ := time.Parse(timeFormat, "20:00")

	if startTime.Before(openTime) || endTime.After(closeTime) || !startTime.Before(endTime) {
		return TimeSlot{}, ErrInvalidTimeRange
	}

	return TimeSlot{StartTime: startTime, EndTime: endTime}, nil
}

// AddBooking adds a booking to the customer
func (c *Customer) AddBooking(booking *Booking) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Bookings[booking.ID] = booking
}

// RemoveBooking removes a booking from the customer
func (c *Customer) RemoveBooking(bookingID string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.Bookings[bookingID]; !exists {
		return ErrBookingNotFound
	}

	delete(c.Bookings, bookingID)
	return nil
}

// GetAllBookings returns all bookings for the customer
func (c *Customer) GetAllBookings() []*Booking {
	c.mu.RLock()
	defer c.mu.RUnlock()

	bookings := make([]*Booking, 0, len(c.Bookings))
	for _, booking := range c.Bookings {
		bookings = append(bookings, booking)
	}

	return bookings
}
