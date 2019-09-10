package structure

import "sync"

// HashTable https://en.wikipedia.org/wiki/Hash_table
type HashTable interface {
	Set(key string, value int)
	Get(key string) int
	Remove(key string)
}

// NewHashTable returns new struct implementing HashTable
func NewHashTable() HashTable {
	return &hashTable{
		data: make(map[string]int),
	}
}

type hashTable struct {
	mu   sync.Mutex
	data map[string]int
}

func (s *hashTable) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *hashTable) Get(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[key]
}

func (s *hashTable) Remove(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}
