package kvstore

import (
	"encoding/json"
	"os"
	"sync"
)

// Store represents a simple key-value store
type Store struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

// New creates a new empty key-value store
func New() *Store {
	return &Store{
		data: make(map[string]interface{}),
	}
}

// Set stores a value for the given key
func (s *Store) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Get retrieves the value for the given key
func (s *Store) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, exists := s.data[key]
	return val, exists
}

// Delete removes a key and its value from the store
func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

// Keys returns all keys in the store
func (s *Store) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

// SaveToFile persists the store to a file in JSON format
func (s *Store) SaveToFile(filename string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	encoder := json.NewEncoder(file)
	return encoder.Encode(s.data)
}

// LoadFromFile loads a store from a JSON file
func (s *Store) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	s.mu.Lock()
	defer s.mu.Unlock()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&s.data)
}

// Clear removes all keys and values from the store
func (s *Store) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make(map[string]interface{})
}
