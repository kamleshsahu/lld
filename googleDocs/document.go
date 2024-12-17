package main

import (
	"sync"
	"time"
)

// Document represents a versioned document
type Document struct {
	ID        string
	Name      string
	Version   int
	Content   string
	History   []DocumentVersion
	UpdatedAt time.Time
	UpdatedBy string
	Mutex     sync.Mutex // For fine-grained locking
}

// DocumentVersion represents a specific version of a document
type DocumentVersion struct {
	Version   int
	Content   string
	UpdatedAt time.Time
	UpdatedBy string
}
