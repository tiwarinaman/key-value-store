package storage

import (
	"sync"
	"time"
)

type Storage struct {
	data map[string]interface{}
	mu   sync.RWMutex
	ttl  map[string]time.Time
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]interface{}),
		ttl:  make(map[string]time.Time),
	}
}

func (s *Storage) Set(key string, value interface{}, expiration int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
	if expiration > 0 {
		s.ttl[key] = time.Now().Add(time.Duration(expiration) * time.Second)
	}
}

func (s *Storage) Get(key string) interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if expiry, exists := s.ttl[key]; exists && time.Now().After(expiry) {
		delete(s.data, key)
		delete(s.ttl, key)
		return nil
	}

	return s.data[key]
}
