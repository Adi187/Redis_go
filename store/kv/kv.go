package kv

//when we increment the set method,if the given key exists do not increment the count of keys,
//if its a new key,increment the count of keys.

import (
	"errors"
	"fmt"
	"sort"
)

func (s *Store) Len() int {
	return len(s.data)
}

var ErrMaxKeysReached = errors.New("maximum number of keys reached")

// Store is a struct that represents a simple key-value store.
type Store struct {
	data    map[string]string
	maxSize int //0=unlimited
}

//set method to add a key-value pair to the store.

func (s *Store) Set(key, value string) error {
	// Check if the key already exists in the store
	_, exists := s.data[key]
	// If the key does not exist, check if the store has reached its maximum size
	if s.maxSize > 0 && len(s.data) >= s.maxSize && !exists {
		//throw an error

		return fmt.Errorf("Set(%q):%w", key, ErrMaxKeysReached)
	}

	s.data[key] = value
	fmt.Printf("Set(%q, %q) called. Current store: %v\n", key, value, s.data)
	return nil
}

func (s *Store) Get(key string) (string, bool, error) {
	if value, exists := s.data[key]; exists {
		return value, true, nil
	}

	return "", false, errors.New("key not found")
}

// NewStore creates a new instance of Store and initializes the data map.
//returns a pointer to the newly created Store instance.
func NewStore(maxSize int) *Store {
	//constructor
	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
}

// Clone returns a store with an independent copy of the data.
func Clone(s *Store) *Store {
	cloned := &Store{
		data:    make(map[string]string, len(s.data)),
		maxSize: s.maxSize,
	}
	for key, value := range s.data {
		cloned.data[key] = value
	}
	return cloned
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	//sort the keys alphabetically

	for key := range s.data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}

func Delete(s *Store, key string) {
	// Delete the key-value pair from the store if it exists
	if _, exists := s.data[key]; exists {
		delete(s.data, key)
		fmt.Println("Key deleted successfully.")
	} else {
		fmt.Println("Key not found.")
	}
}

func (s *Store) keys() []string { return s.Keys() }

func (s *Store) Delete(key string) { delete(s.data, key) }
