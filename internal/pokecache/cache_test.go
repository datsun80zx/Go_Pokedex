package pokecache

import (
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	cache := NewCache(5 * time.Minute)
	
	// Test cases
	testCases := []struct {
		key   string
		value []byte
	}{
		{"test-key-1", []byte("test value 1")},
		{"test-key-2", []byte("test value 2")},
		{"test-key-3", []byte("test value 3")},
	}
	
	// Add entries to cache
	for _, tc := range testCases {
		cache.Add(tc.key, tc.value)
	}
	
	// Check if entries can be retrieved
	for _, tc := range testCases {
		val, found := cache.Get(tc.key)
		if !found {
			t.Errorf("Expected to find key %s in cache, but it was not found", tc.key)
			continue
		}
		
		// Check if the value matches
		if string(val) != string(tc.value) {
			t.Errorf("Expected value %s for key %s, but got %s", string(tc.value), tc.key, string(val))
		}
	}
	
	// Test getting a non-existent key
	_, found := cache.Get("non-existent-key")
	if found {
		t.Errorf("Expected not to find non-existent key, but it was found")
	}
}

func TestCacheReap(t *testing.T) {
	// Create a cache with a very short interval for testing
	cache := NewCache(10 * time.Millisecond)
	
	// Add an entry
	cache.Add("test-key", []byte("test value"))
	
	// Verify the entry exists
	_, found := cache.Get("test-key")
	if !found {
		t.Errorf("Expected to find key in cache immediately after adding")
	}
	
	// Wait for the reaper to run
	time.Sleep(20 * time.Millisecond)
	
	// The entry should be gone now
	_, found = cache.Get("test-key")
	if found {
		t.Errorf("Expected key to be removed from cache after expiration")
	}
}