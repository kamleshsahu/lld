package repo

import (
	"fmt"
	"lld/gymchain/entity"
	"sync"
)

// ==================== Repository Layer ======================

// GymRepository manages gym data storage and retrieval
type GymRepository struct {
	gyms       map[string]*entity.Gym
	gymCounter int
	mu         sync.RWMutex
}

// NewGymRepository creates a new gym repository
func NewGymRepository() *GymRepository {
	return &GymRepository{
		gyms: make(map[string]*entity.Gym),
	}
}

// CreateGym creates and stores a new gym
func (r *GymRepository) CreateGym(name, location string, maxAccommodation int) (*entity.Gym, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.gymCounter++
	gymID := fmt.Sprintf("gym_%d", r.gymCounter)

	gym := &entity.Gym{
		ID:               gymID,
		Name:             name,
		Location:         location,
		MaxAccommodation: maxAccommodation,
		Classes:          make(map[string]*entity.GymClass),
	}

	r.gyms[gymID] = gym
	return gym, nil
}

// GetGym retrieves a gym by ID
func (r *GymRepository) GetGym(gymID string) (*entity.Gym, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	gym, exists := r.gyms[gymID]
	if !exists {
		return nil, entity.ErrGymNotFound
	}

	return gym, nil
}

// DeleteGym removes a gym
func (r *GymRepository) DeleteGym(gymID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.gyms[gymID]; !exists {
		return entity.ErrGymNotFound
	}

	delete(r.gyms, gymID)
	return nil
}

// GetAllGyms returns all gyms
func (r *GymRepository) GetAllGyms() []*entity.Gym {
	r.mu.RLock()
	defer r.mu.RUnlock()

	gyms := make([]*entity.Gym, 0, len(r.gyms))
	for _, gym := range r.gyms {
		gyms = append(gyms, gym)
	}

	return gyms
}
