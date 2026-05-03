package main

import (
	"fmt"
	"sync"
)

type InMemoryDB struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]string),
	}
}

// SET
func (db *InMemoryDB) Set(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

// GET
func (db *InMemoryDB) Get(key string) (string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	val, ok := db.data[key]
	return val, ok
}

// DELETE
func (db *InMemoryDB) Delete(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, key)
}

func main() {
	db := NewDB()

	db.Set("user:1", "Kumail")
	db.Set("user:2", "Alice")

	val, ok := db.Get("user:1")
	if ok {
		fmt.Println("Found:", val)
	}

	db.Delete("user:2")
}
