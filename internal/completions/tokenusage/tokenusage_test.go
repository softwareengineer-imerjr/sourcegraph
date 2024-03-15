package tokenusage_test

import (
	"time"
)

// MockCache is a simple in-memory cache for testing purposes.
// It now includes a basic TTL mechanism for keys.
type MockCache struct {
	store map[string]cachedValue
}

type cachedValue struct {
	value     int
	expiredAt time.Time
}

func NewMockCache() *MockCache {
	return &MockCache{store: make(map[string]cachedValue)}
}

func (m *MockCache) GetInt(key string) (int, bool) {
	val, exists := m.store[key]
	if !exists || val.expiredAt.Before(time.Now()) {
		return 0, false
	}
	return val.value, true
}

func (m *MockCache) SetInt(key string, val int) {
	// For simplicity, we're not handling TTL here, assuming values do not expire.
	// To handle TTL, you would need to adjust this method to accept and process a TTL value.
	m.store[key] = cachedValue{value: val, expiredAt: time.Now().Add(24 * time.Hour)}
}

func (m *MockCache) ListAllKeys() []string {
	keys := make([]string, 0, len(m.store))
	for key, val := range m.store {
		if val.expiredAt.After(time.Now()) {
			keys = append(keys, key)
		}
	}
	return keys
}
