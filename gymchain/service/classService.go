package service

import (
	"lld/gymchain/entity"
	"lld/gymchain/repo"
)

// ClassService handles operations related to classes
type ClassService struct {
	gymRepo   *repo.GymRepository
	classRepo *repo.ClassRepository
}

// NewClassService creates a new class service
func NewClassService(gymRepo *repo.GymRepository, classRepo *repo.ClassRepository) *ClassService {
	return &ClassService{
		gymRepo:   gymRepo,
		classRepo: classRepo,
	}
}

// AddClass adds a new class to a gym
func (s *ClassService) AddClass(gymID, classType string, maxLimit int, startTimeStr, endTimeStr string) (string, error) {
	if maxLimit <= 0 {
		return "", entity.ErrInvalidMaxLimit
	}

	// Parse and validate time
	timeSlot, err := entity.ParseAndValidateTime(startTimeStr, endTimeStr)
	if err != nil {
		return "", err
	}

	// Get the gym
	gym, err := s.gymRepo.GetGym(gymID)
	if err != nil {
		return "", err
	}

	// Generate a new class ID
	classID := s.classRepo.GenerateClassID()

	// Add the class to the gym
	err = gym.AddClass(classID, classType, maxLimit, timeSlot)
	if err != nil {
		return "", err
	}

	return classID, nil
}

// RemoveClass removes a class from a gym
func (s *ClassService) RemoveClass(gymID, classID string) error {
	// Get the gym
	gym, err := s.gymRepo.GetGym(gymID)
	if err != nil {
		return err
	}

	return gym.RemoveClass(classID)
}

// BookingService handles operations related to bookings
type BookingService struct {
	gymRepo      *repo.GymRepository
	customerRepo *repo.CustomerRepository
	bookingRepo  *repo.BookingRepository
}
