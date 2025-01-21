package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type RandomWithExclusion struct {
	exclusionMap map[int]bool
	filePath     string
}

func NewRandomWithExclusion(filePath string) *RandomWithExclusion {
	r := &RandomWithExclusion{
		exclusionMap: make(map[int]bool),
		filePath:     filePath,
	}
	// Load exclusions from file if it exists
	r.loadExclusions()
	return r
}

func (r *RandomWithExclusion) loadExclusions() {
	file, err := os.Open(r.filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error reading exclusions file:", err)
		}
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&r.exclusionMap)
	if err != nil {
		fmt.Println("Error decoding exclusions file:", err)
	}
}

func (r *RandomWithExclusion) saveExclusions() {
	file, err := os.Create(r.filePath)
	if err != nil {
		fmt.Println("Error writing exclusions file:", err)
		return
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(r.exclusionMap)
	if err != nil {
		fmt.Println("Error encoding exclusions to file:", err)
	}
}

func (r *RandomWithExclusion) Generate(start, end int) (int, error) {
	// Create a slice of valid numbers
	validNumbers := []int{}
	for i := start; i <= end; i++ {
		if !r.exclusionMap[i] {
			validNumbers = append(validNumbers, i)
		}
	}

	// Check if there are valid numbers left
	if len(validNumbers) == 0 {
		return 0, errors.New("no valid numbers available within the given range")
	}

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Select a random number from valid numbers
	randomIndex := rand.Intn(len(validNumbers))
	randomNumber := validNumbers[randomIndex]

	// Persist the exclusion
	r.exclusionMap[randomNumber] = true
	r.saveExclusions()

	return randomNumber, nil
}

func (r *RandomWithExclusion) ResetExclusions() {
	r.exclusionMap = make(map[int]bool)
	r.saveExclusions()
}

func main() {
	filePath := "exclusions.json"
	generator := NewRandomWithExclusion(filePath)

	randomNumber, err := generator.Generate(2, 301)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Generated:", randomNumber)
}
