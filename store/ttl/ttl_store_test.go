//write all the tests for the ttl store methods
package ttl

import (
	"testing"
	"time"
)

func TestTTLStore_SetAndGet(t *testing.T) {
	store := NewTTLStore(2 * time.Second)

	err := store.Set("key1", "value1")
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}	
	
	value, ok := store.GetWithTTL("key1")
	if !ok || value != "value1" {
		t.Fatalf("GetWithTTL failed: expected value1, got %v", value)
	}
}

func TestTTLStore_Expiration(t *testing.T) {
	store := NewTTLStore(1 * time.Second)

	err := store.Set("key2", "value2")
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	time.Sleep(2 * time.Second) // Wait for the key to expire

	value, ok := store.GetWithTTL("key2")
	if ok || value != "" {
		t.Fatalf("GetWithTTL should have expired, but got value: %v", value)
	}
}

func TestTTLStore_EmptyKey(t *testing.T) {
	store := NewTTLStore(1 * time.Second)

	err := store.Set("", "value")
	if err == nil {
		t.Fatalf("Set should have failed for empty key")
	}

	value, ok := store.GetWithTTL("")
	if ok || value != "" {
		t.Fatalf("GetWithTTL should have failed for empty key, but got value: %v", value)
	}
}		

