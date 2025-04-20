
package cache

import (
	"testing"
	"time"
)

type TestStruct struct {
	Name string
	Age  int
}

func TestRedisCache(t *testing.T) {
	cache, err := NewRedisCache("redis://localhost:6379")
	if err != nil {
		t.Skip("Redis not available:", err)
	}

	// Test Set and Get
	testData := TestStruct{Name: "Test", Age: 25}
	err = cache.Set("test", testData, time.Minute)
	if err != nil {
		t.Error("Failed to set cache:", err)
	}

	var result TestStruct
	err = cache.Get("test", &result)
	if err != nil {
		t.Error("Failed to get cache:", err)
	}
	if result.Name != testData.Name || result.Age != testData.Age {
		t.Error("Cache data mismatch")
	}

	// Test Delete
	err = cache.Delete("test")
	if err != nil {
		t.Error("Failed to delete cache:", err)
	}

	err = cache.Get("test", &result)
	if err == nil {
		t.Error("Expected error getting deleted cache")
	}
}
