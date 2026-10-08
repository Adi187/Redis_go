package kv

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	store := NewStore(0)
	store.Set("c", "3")
	store.Set("a", "1")
	got := store.keys()
	want := []string{"a", "c"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() = %v, want %v", got, want)
	}

}

func TestKeys_ReturnSortedKeys(t *testing.T) {
	store := NewStore(0)

	got := store.keys()
	if len(got) != 0 {
		t.Errorf("Keys() = %v, want empty slice", got)
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	//make a new store
	//set("hello","world")
	//3.Call Get("hello") - assert it returns "world" and true
	//4. Call Get("nonexistent") - assert it returns "" and false

	store := NewStore(0)
	store.Set("hello", "world")

	value, exists, _ := store.Get("hello")
	if !exists || value != "world" {
		t.Error("Get() failed")

	}

	// Test getting a non-existent key
	value, exists, _ = store.Get("nonexistent")
	if exists || value != "" {
		t.Error("Get() failed")

	}
}

// Table driven test for removing keys
func TestRemove(t *testing.T) {
	//define a test structure
	tests := []struct {
		name        string
		setupKey    map[string]string
		deleteKey   string
		expectedLen int
	}{
		{"delete existing key", map[string]string{"a": "1", "b": "2"}, "a", 1},
		{"delete non-existing key", map[string]string{"a": "1", "b": "2"}, "c", 2},
		{"delete from empty store", map[string]string{}, "a", 0},
		{"delete last key", map[string]string{"a": "1"}, "a", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewStore(0)
			for k, v := range tt.setupKey {
				store.Set(k, v)
			}
			Delete(store, tt.deleteKey)
			if len(store.keys()) != tt.expectedLen {
				t.Error("Remove() failed")
			}
		})
	}
}
