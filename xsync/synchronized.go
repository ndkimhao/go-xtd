package xsync

import (
	"sync"
)

// Synchronized forces user to lock a shared object before using
type Synchronized[T any] struct {
	Mutex sync.Mutex
	Value T
}

func (s *Synchronized[T]) Lock() *T {
	s.Mutex.Lock()
	return &s.Value
}

func (s *Synchronized[T]) Unlock() {
	s.Mutex.Unlock()
}

// RWSynchronized forces user to lock a shared object before using
type RWSynchronized[T any] struct {
	RWMutex sync.RWMutex
	Value   T
}

func (s *RWSynchronized[T]) Lock() *T {
	s.RWMutex.Lock()
	return &s.Value
}

func (s *RWSynchronized[T]) Unlock() {
	s.RWMutex.Unlock()
}

func (s *RWSynchronized[T]) RLock() *T {
	s.RWMutex.RLock()
	return &s.Value
}

func (s *RWSynchronized[T]) RUnlock() {
	s.RWMutex.RUnlock()
}
