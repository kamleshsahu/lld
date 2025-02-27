package service

import (
	"fmt"
	"lld/gymchain/entity"
	"lld/gymchain/repo"
	"time"
)

// NewBookingService creates a new booking service
func NewBookingService(gymRepo *repo.GymRepository, customerRepo *repo.CustomerRepository, bookingRepo *repo.BookingRepository) *BookingService {
	return &BookingService{
		gymRepo:      gymRepo,
		customerRepo: customerRepo,
		bookingRepo:  bookingRepo,
	}
}

// BookClass books a class for a customer
func (s *BookingService) BookClass(customerID, gymID, classID string) (string, error) {
	// Get the gym
	gym, err := s.gymRepo.GetGym(gymID)
	if err != nil {
		return "", err
	}

	// Get the class
	class, err := gym.GetClass(classID)
	if err != nil {
		return "", err
	}

	// Ensure customer exists
	customer := s.customerRepo.CreateCustomer(customerID)

	// Create booking
	booking := s.bookingRepo.CreateBooking(customerID, gymID, classID)

	// Add booking to class
	err = class.AddBooking(booking.ID, booking)
	if err != nil {
		// Rollback if class booking fails
		err = s.bookingRepo.DeleteBooking(booking.ID)
		if err != nil {
			return "", err
		}
		return "", err
	}

	// Add booking to customer
	customer.AddBooking(booking)

	return booking.ID, nil
}

// CancelBooking cancels a booking
func (s *BookingService) CancelBooking(bookingID string) error {
	// Get the booking
	booking, err := s.bookingRepo.GetBooking(bookingID)
	if err != nil {
		return err
	}

	// Get the gym and class
	gym, err := s.gymRepo.GetGym(booking.GymID)
	if err != nil {
		return err
	}

	// Get the class
	class, err := gym.GetClass(booking.ClassID)
	if err != nil {
		return err
	}

	// Remove booking from class
	err = class.RemoveBooking(bookingID)
	if err != nil {
		return err
	}

	// Try to get the customer and remove booking
	customer, err := s.customerRepo.GetCustomer(booking.CustomerID)
	if err == nil {
		err = customer.RemoveBooking(bookingID)
		if err != nil {
			return err
		}
	}

	// Remove booking from repository
	return s.bookingRepo.DeleteBooking(bookingID)
}

// GetAllBookings returns all bookings for a customer
func (s *BookingService) GetAllBookings(customerID string) ([]*entity.Booking, error) {
	customer, err := s.customerRepo.GetCustomer(customerID)
	if err != nil {
		return nil, err
	}

	return customer.GetAllBookings(), nil
}

// PrintCustomerBookings prints all bookings for a customer
func (s *BookingService) PrintCustomerBookings(customerID string) {
	bookings, err := s.GetAllBookings(customerID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Bookings for Customer %s:\n", customerID)
	if len(bookings) == 0 {
		fmt.Println("  No bookings found")
		return
	}

	for _, booking := range bookings {
		gym, err := s.gymRepo.GetGym(booking.GymID)
		if err != nil {
			fmt.Printf("  Booking ID: %s (Gym no longer exists)\n", booking.ID)
			continue
		}

		class, err := gym.GetClass(booking.ClassID)
		if err != nil {
			fmt.Printf("  Booking ID: %s (Class no longer exists)\n", booking.ID)
			continue
		}

		classDetails := class.GetDetails()
		gymDetails := gym.GetDetails()

		fmt.Printf("  Booking ID: %s\n", booking.ID)
		fmt.Printf("    Gym: %s (%s)\n", gymDetails["name"], gymDetails["location"])
		fmt.Printf("    Class: %s\n", classDetails["type"])
		fmt.Printf("    Time: %s - %s\n",
			classDetails["startTime"],
			classDetails["endTime"])
		fmt.Printf("    Booked at: %s\n", booking.BookingTime.Format(time.RFC3339))
	}
}
