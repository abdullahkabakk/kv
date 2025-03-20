package kvstore

import (
	"os"
	"testing"
)

func TestBasicOperations(t *testing.T) {
	store := New()

	// Test Set and Get
	store.Set("key1", "value1")
	val, exists := store.Get("key1")
	if !exists {
		t.Fatalf("Expected key1 to exist")
	}
	if val != "value1" {
		t.Fatalf("Expected value1, got %v", val)
	}

	// Test non-existent key
	_, exists = store.Get("nonexistent")
	if exists {
		t.Fatalf("Expected nonexistent key to not exist")
	}

	// Test Delete
	store.Delete("key1")
	_, exists = store.Get("key1")
	if exists {
		t.Fatalf("Expected key1 to be deleted")
	}
}

func TestKeys(t *testing.T) {
	store := New()
	store.Set("key1", "value1")
	store.Set("key2", "value2")

	keys := store.Keys()
	if len(keys) != 2 {
		t.Fatalf("Expected 2 keys, got %d", len(keys))
	}

	// Check both keys exist in the slice
	foundKey1, foundKey2 := false, false
	for _, k := range keys {
		if k == "key1" {
			foundKey1 = true
		}
		if k == "key2" {
			foundKey2 = true
		}
	}

	if !foundKey1 || !foundKey2 {
		t.Fatalf("Not all expected keys were returned")
	}
}

func TestPersistence(t *testing.T) {
	filename := "test_store.json"
	// Clean up after test
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatalf("Failed to remove test file: %v", err)
		}
	}(filename)

	// Create and save store
	store := New()
	store.Set("key1", "value1")
	store.Set("key2", 42)

	err := store.SaveToFile(filename)
	if err != nil {
		t.Fatalf("Failed to save: %v", err)
	}

	// Load into new store
	newStore := New()
	err = newStore.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("Failed to load: %v", err)
	}

	// Verify values
	val1, exists := newStore.Get("key1")
	if !exists || val1 != "value1" {
		t.Fatalf("Expected key1=value1, got %v", val1)
	}

	val2, exists := newStore.Get("key2")
	if !exists {
		t.Fatalf("Expected key2 to exist")
	}

	if val2.(float64) != 42 {
		t.Fatalf("Expected key2=42, got %v (%T)", val2, val2)
	}
}
