package ttl

import (
	"redis_go/store"
	"sort"
	"time"
	"errors"
)

type ttlEntry struct {
	value     string
	expiresAt time.Time
}

type TTLStore struct {
	data map[string]ttlEntry
	ttl  time.Duration
}

func NewTTLStore(ttl time.Duration) *TTLStore {
	return &TTLStore{data: make(map[string]ttlEntry), ttl: ttl}
}

func (t *TTLStore) Set(key, value string) error {
	if key == "" {
		return store.ErrEmptykey
	}
	t.data[key] = ttlEntry{value: value, expiresAt: time.Now().Add(t.ttl)}
	return nil
}

func (t *TTLStore) Get(key string) (string, bool,error) {
	if key == "" {
		return "", false, store.ErrEmptykey	
	}
	entry, ok := t.data[key]
	if !ok || !time.Now().Before(entry.expiresAt) {
		delete(t.data, key)
		return "", false, errors.New("key not found or expired")
	}
	return entry.value, true, nil
}

func (t *TTLStore) Keys() []string {
	keys := make([]string, 0, len(t.data))
	for key := range t.data {
		if _, ok, _ := t.Get(key); ok {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	return keys
}

func (t *TTLStore) Delete(key string) {
	delete(t.data, key)
}

func (t *TTLStore) Len() int {
	return len(t.Keys())
}
