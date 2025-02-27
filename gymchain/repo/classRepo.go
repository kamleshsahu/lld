package repo

import (
	"fmt"
	"sync"
)

// ClassRepository manages class data storage and retrieval
type ClassRepository struct {
	classCounter int
	mu           sync.RWMutex
}

// NewClassRepository creates a new class repository
func NewClassRepository() *ClassRepository {
	return &ClassRepository{}
}

// GenerateClassID generates a unique class ID
func (r *ClassRepository) GenerateClassID() string {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.classCounter++
	return fmt.Sprintf("class_%d", r.classCounter)
}
