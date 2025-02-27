package service

import (
	"fmt"
	"lld/gymchain/entity"
	"lld/gymchain/repo"
)

// GymService handles operations related to gyms
type GymService struct {
	gymRepo   *repo.GymRepository
	classRepo *repo.ClassRepository
}

// NewGymService creates a new gym service
func NewGymService(gymRepo *repo.GymRepository, classRepo *repo.ClassRepository) *GymService {
	return &GymService{
		gymRepo:   gymRepo,
		classRepo: classRepo,
	}
}

// AddGym adds a new gym
func (s *GymService) AddGym(name, location string, maxAccommodation int) (string, error) {
	if maxAccommodation <= 0 {
		return "", entity.ErrInvalidAccommodation
	}

	gym, err := s.gymRepo.CreateGym(name, location, maxAccommodation)
	if err != nil {
		return "", err
	}

	return gym.ID, nil
}

// RemoveGym removes a gym and all its classes
func (s *GymService) RemoveGym(gymID string) error {
	return s.gymRepo.DeleteGym(gymID)
}

// PrintGymDetails prints detailed information about a gym
func (s *GymService) PrintGymDetails(gymID string) {
	gym, err := s.gymRepo.GetGym(gymID)
	if err != nil {
		fmt.Printf("Gym with ID %s not found\n", gymID)
		return
	}

	details := gym.GetDetails()

	fmt.Printf("Gym Details:\n")
	fmt.Printf("  ID: %s\n", details["id"])
	fmt.Printf("  Name: %s\n", details["name"])
	fmt.Printf("  Location: %s\n", details["location"])
	fmt.Printf("  Max Accommodation: %d\n", details["maxAccommodation"])
	fmt.Printf("  Classes: %d\n", details["classCount"])

	gym.Mu.RLock()
	defer gym.Mu.RUnlock()

	for classID, class := range gym.Classes {
		classDetails := class.GetDetails()
		fmt.Printf("    Class ID: %s\n", classID)
		fmt.Printf("      Type: %s\n", classDetails["type"])
		fmt.Printf("      Time: %s - %s\n",
			classDetails["startTime"],
			classDetails["endTime"])
		fmt.Printf("      Bookings: %d/%d\n",
			classDetails["currentBookings"],
			classDetails["maxLimit"])
	}
}
