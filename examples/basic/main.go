package main

import (
	"fmt"
	kvstore "github.com/abdullahkabakk/kv"
	"log"
)

func main() {
	// Create a new store
	store := kvstore.New()

	// Set some values
	store.Set("name", "John Doe")
	store.Set("age", 30)
	store.Set("languages", []string{"Go", "Python", "JavaScript"})

	// Get a value
	if name, exists := store.Get("name"); exists {
		fmt.Printf("Name: %v\n", name)
	}

	// Get all keys
	keys := store.Keys()
	fmt.Println("All keys:", keys)

	// Save to a file
	err := store.SaveToFile("data.json")
	if err != nil {
		log.Fatalf("Failed to save store: %v", err)
	}

	// Create a new store and load from file
	newStore := kvstore.New()
	err = newStore.LoadFromFile("data.json")
	if err != nil {
		log.Fatalf("Failed to load store: %v", err)
	}

	// Verify loaded data
	if age, exists := newStore.Get("age"); exists {
		fmt.Printf("Age: %v\n", age)
	}
}
