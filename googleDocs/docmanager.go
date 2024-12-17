package main

import (
	"errors"
	"fmt"
	"time"
)

// DocumentManager manages versioned documents
type DocumentManager struct {
	Docs               map[string]*Document
	PermissionEnforcer PermissionEnforcer
}

func NewDocumentManager() *DocumentManager {
	return &DocumentManager{
		Docs:               make(map[string]*Document),
		PermissionEnforcer: PermissionEnforcer{},
	}
}

// Create a new document
func (dm *DocumentManager) Create(user *User, name, content string) (*Document, error) {
	if err := dm.PermissionEnforcer.Enforce(user, CanCreate); err != nil {
		return nil, err
	}

	doc := &Document{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		Name:      name,
		Content:   content,
		Version:   1,
		History:   []DocumentVersion{},
		UpdatedAt: time.Now(),
		UpdatedBy: user.Name,
	}
	dm.Docs[doc.ID] = doc
	fmt.Printf("Document created by %s: %+v\n", user.Name, doc)
	return doc, nil
}

// Update a document with optimistic locking
func (dm *DocumentManager) Update(user *User, docID, newContent string, expectedVersion int) error {
	if err := dm.PermissionEnforcer.Enforce(user, CanUpdate); err != nil {
		return err
	}

	doc, exists := dm.Docs[docID]
	if !exists {
		return errors.New("document not found")
	}

	// Lock the document for update
	doc.Mutex.Lock()
	defer doc.Mutex.Unlock()

	// Check for version conflict
	if doc.Version != expectedVersion {
		return errors.New(fmt.Sprintf("conflict detected: document version is %d, but expected %d", doc.Version, expectedVersion))
	}

	// Save current version to history
	doc.History = append(doc.History, DocumentVersion{
		Version:   doc.Version,
		Content:   doc.Content,
		UpdatedAt: doc.UpdatedAt,
		UpdatedBy: doc.UpdatedBy,
	})

	// Update the document
	doc.Version++
	doc.Content = newContent
	doc.UpdatedAt = time.Now()
	doc.UpdatedBy = user.Name

	fmt.Printf("Document updated by %s to version %d\n", user.Name, doc.Version)
	return nil
}

// View a document
func (dm *DocumentManager) View(user *User, docID string) (*Document, error) {
	if err := dm.PermissionEnforcer.Enforce(user, CanView); err != nil {
		return nil, err
	}

	doc, exists := dm.Docs[docID]
	if !exists {
		return nil, errors.New("document not found")
	}

	fmt.Printf("Document viewed by %s: %+v\n", user.Name, doc)
	return doc, nil
}

// Restore to a specific version with locking
func (dm *DocumentManager) RestoreVersion(user *User, docID string, version int, expectedVersion int) error {
	if err := dm.PermissionEnforcer.Enforce(user, CanUpdate); err != nil {
		return err
	}

	doc, exists := dm.Docs[docID]
	if !exists {
		return errors.New("document not found")
	}

	// Lock the document for restore
	doc.Mutex.Lock()
	defer doc.Mutex.Unlock()

	// Check for version conflict
	if doc.Version != expectedVersion {
		return errors.New(fmt.Sprintf("conflict detected: document version is %d, but expected %d", doc.Version, expectedVersion))
	}

	// Find the specified version
	for _, v := range doc.History {
		if v.Version == version {
			// Save current version to history
			doc.History = append(doc.History, DocumentVersion{
				Version:   doc.Version,
				Content:   doc.Content,
				UpdatedAt: doc.UpdatedAt,
				UpdatedBy: doc.UpdatedBy,
			})

			// Restore document
			doc.Version++
			doc.Content = v.Content
			doc.UpdatedAt = time.Now()
			doc.UpdatedBy = user.Name
			fmt.Printf("Document restored by %s to version %d\n", user.Name, version)
			return nil
		}
	}

	return errors.New("specified version not found")
}
