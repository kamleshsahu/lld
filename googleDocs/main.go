package main

import (
	"fmt"
	"time"
)

func main() {
	// Initialize Document Manager
	docManager := NewDocumentManager()

	// Define users
	admin := &User{ID: "1", Name: "Alice", Role: AdminRole}
	editor := &User{ID: "2", Name: "Bob", Role: EditorRole}

	// Admin creates a document
	doc, _ := docManager.Create(admin, "Versioned Doc", "Initial Content")

	// Simulate concurrent updates
	go func() {
		err := docManager.Update(editor, doc.ID, "Concurrent Update 1", 1)
		if err != nil {
			fmt.Println("Update 1 failed:", err)
		} else {
			fmt.Println("Update 1 succeeded")
		}
	}()

	go func() {
		err := docManager.Update(editor, doc.ID, "Concurrent Update 2", 1)
		if err != nil {
			fmt.Println("Update 2 failed:", err)
		} else {
			fmt.Println("Update 2 succeeded")
		}
	}()

	// Wait for goroutines to finish
	time.Sleep(1 * time.Second)
}
