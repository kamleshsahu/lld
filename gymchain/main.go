package main

import (
	"fmt"
	"lld/gymchain/repo"
	"lld/gymchain/service"
	"sync"
)

// ======================== Model Layer ========================

// ===================== Main Application =====================

func main() {
	// Create repositories
	gymRepo := repo.NewGymRepository()
	classRepo := repo.NewClassRepository()
	customerRepo := repo.NewCustomerRepository()
	bookingRepo := repo.NewBookingRepository()

	// Create services
	gymService := service.NewGymService(gymRepo, classRepo)
	classService := service.NewClassService(gymRepo, classRepo)
	bookingService := service.NewBookingService(gymRepo, customerRepo, bookingRepo)

	// Demonstrate admin flows
	fmt.Println("=== ADMIN FLOWS ===")

	// Add gyms
	gymID1, err := gymService.AddGym("Fitness First", "Indira Nagar", 100)
	if err != nil {
		fmt.Printf("Error adding gym: %v\n", err)
	} else {
		fmt.Printf("Added gym: %s\n", gymID1)
	}

	gymID2, err := gymService.AddGym("Gold's Gym", "Koramangala", 150)
	if err != nil {
		fmt.Printf("Error adding gym: %v\n", err)
	} else {
		fmt.Printf("Added gym: %s\n", gymID2)
	}

	// Add classes
	classID1, err := classService.AddClass(gymID1, "Cardio", 20, "06:00", "07:00")
	if err != nil {
		fmt.Printf("Error adding class: %v\n", err)
	} else {
		fmt.Printf("Added class: %s\n", classID1)
	}

	classID2, err := classService.AddClass(gymID1, "Weights", 15, "07:30", "08:30")
	if err != nil {
		fmt.Printf("Error adding class: %v\n", err)
	} else {
		fmt.Printf("Added class: %s\n", classID2)
	}

	classID3, err := classService.AddClass(gymID2, "Yoga", 25, "09:00", "10:30")
	if err != nil {
		fmt.Printf("Error adding class: %v\n", err)
	} else {
		fmt.Printf("Added class: %s\n", classID3)
	}

	// Demonstrate conflict
	_, err = classService.AddClass(gymID1, "Conflict", 10, "06:30", "07:30")
	fmt.Printf("Expected conflict error: %v\n", err)

	// Print gym details
	fmt.Println("\nGym Details:")
	gymService.PrintGymDetails(gymID1)

	// Demonstrate customer flows
	fmt.Println("\n=== CUSTOMER FLOWS ===")

	// Book classes
	customer1 := "customer_1"
	customer2 := "customer_2"

	booking1, err := bookingService.BookClass(customer1, gymID1, classID1)
	if err != nil {
		fmt.Printf("Error booking class: %v\n", err)
	} else {
		fmt.Printf("Booked class: %s\n", booking1)
	}

	booking2, err := bookingService.BookClass(customer1, gymID1, classID2)
	if err != nil {
		fmt.Printf("Error booking class: %v\n", err)
	} else {
		fmt.Printf("Booked class: %s\n", booking2)
	}

	booking3, err := bookingService.BookClass(customer2, gymID1, classID1)
	if err != nil {
		fmt.Printf("Error booking class: %v\n", err)
	} else {
		fmt.Printf("Booked class: %s\n", booking3)
	}

	// Try booking the same class again (should fail)
	_, err = bookingService.BookClass(customer1, gymID1, classID1)
	fmt.Printf("Expected already booked error: %v\n", err)

	// Print customer bookings
	fmt.Println("\nCustomer Bookings:")
	bookingService.PrintCustomerBookings(customer1)

	// Cancel booking
	fmt.Println("\nCancelling booking:")
	err = bookingService.CancelBooking(booking1)
	if err != nil {
		fmt.Printf("Error cancelling booking: %v\n", err)
	} else {
		fmt.Printf("Cancelled booking: %s\n", booking1)
	}

	// Check bookings after cancellation
	fmt.Println("\nCustomer Bookings after cancellation:")
	bookingService.PrintCustomerBookings(customer1)

	// Remove class
	fmt.Println("\nRemoving class:")
	err = classService.RemoveClass(gymID1, classID2)
	if err != nil {
		fmt.Printf("Error removing class: %v\n", err)
	} else {
		fmt.Printf("Removed class: %s\n", classID2)
	}

	// Check gym details after class removal
	fmt.Println("\nGym Details after class removal:")
	gymService.PrintGymDetails(gymID1)

	// Check customer bookings after class removal
	fmt.Println("\nCustomer Bookings after class removal:")
	bookingService.PrintCustomerBookings(customer1)

	// Remove gym
	fmt.Println("\nRemoving gym:")
	err = gymService.RemoveGym(gymID1)
	if err != nil {
		fmt.Printf("Error removing gym: %v\n", err)
	} else {
		fmt.Printf("Removed gym: %s\n", gymID1)
	}

	// Check customer bookings after gym removal
	fmt.Println("\nCustomer Bookings after gym removal:")
	bookingService.PrintCustomerBookings(customer1)

	// Demonstrate concurrency by booking multiple classes simultaneously
	fmt.Println("\n=== CONCURRENCY DEMONSTRATION ===")
	gymID3, _ := gymService.AddGym("Cult Fitness", "HSR Layout", 50)
	classID4, _ := classService.AddClass(gymID3, "HIIT", 5, "18:00", "19:00")

	// Create a wait group for goroutines
	var wg sync.WaitGroup

	// Launch 10 concurrent booking attempts for a class with max limit 5
	fmt.Println("Launching 10 concurrent bookings for a class with max limit 5:")
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(customerNum int) {
			defer wg.Done()
			custID := fmt.Sprintf("concurrent_customer_%d", customerNum)
			bookingID, err := bookingService.BookClass(custID, gymID3, classID4)
			if err != nil {
				fmt.Printf("Customer %d booking failed: %v\n", customerNum, err)
			} else {
				fmt.Printf("Customer %d booked successfully: %s\n", customerNum, bookingID)
			}
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Check final class status
	fmt.Println("\nFinal class status after concurrent bookings:")
	gymService.PrintGymDetails(gymID3)
}
