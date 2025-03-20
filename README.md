# Key Value Store

A simple, thread-safe key-value store library for Go with persistence support.

## Features

- Simple and intuitive API
- Thread-safe operations
- In-memory storage with JSON persistence
- No external dependencies
- Fully tested

## Installation

```
go get github.com/abdullahkabakk/kv
```

## Usage

### Basic Operations

```go
package main

import (
	"fmt"
	"github.com/abdullahkabakk/kv"
)

func main() {
	// Create a new store
	store := kvstore.New()

	// Set values
	store.Set("name", "John Doe")
	store.Set("age", 30)
	store.Set("languages", []string{"Go", "Python", "JavaScript"})

	// Get a value
	if name, exists := store.Get("name"); exists {
		fmt.Printf("Name: %v\n", name)
	}

	// List all keys
	keys := store.Keys()
	fmt.Println("All keys:", keys)

	// Delete a key
	store.Delete("age")

	// Clear the store
	store.Clear()
}
```

### Persistence

```go
package main

import (
	"fmt"
	"log"
	"github.com/abdullahkabakk/kv"
)

func main() {
	// Create and populate a store
	store := kvstore.New()
	store.Set("user", map[string]interface{}{
		"name": "Alice",
		"role": "Admin",
	})

	// Save to file
	err := store.SaveToFile("data.json")
	if err != nil {
		log.Fatalf("Failed to save: %v", err)
	}

	// Later, load the data
	newStore := kvstore.New()
	err = newStore.LoadFromFile("data.json")
	if err != nil {
		log.Fatalf("Failed to load: %v", err)
	}

	// Use the loaded data
	if user, exists := newStore.Get("user"); exists {
		fmt.Println("Loaded user data:", user)
	}
}
```

## API Reference

### `New() *Store`

Creates a new empty key-value store.

### `(s *Store) Set(key string, value interface{})`

Stores a value for the given key.

### `(s *Store) Get(key string) (interface{}, bool)`

Retrieves the value for the given key. Returns the value and a boolean indicating whether the key exists.

### `(s *Store) Delete(key string)`

Removes a key and its value from the store.

### `(s *Store) Keys() []string`

Returns all keys in the store.

### `(s *Store) SaveToFile(filename string) error`

Persists the store to a file in JSON format.

### `(s *Store) LoadFromFile(filename string) error`

Loads a store from a JSON file.

### `(s *Store) Clear()`

Removes all keys and values from the store.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request