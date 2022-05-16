package xmap

import (
	"github.com/ndkimhao/go-xtd/ds/xslice"
	"github.com/ndkimhao/go-xtd/xfn"
)

type Custom[K, V any] struct {
	main  map[uint64]Entry[K, V]
	ovf   map[uint64][]Entry[K, V]
	hash  func(K) uint64
	equal func(K, K) bool
}

func NewCustom[K, V any](hasher xfn.Function[K, uint64], comparator xfn.BiPredicate[K, K]) Custom[K, V] {
	return Custom[K, V]{
		main:  map[uint64]Entry[K, V]{},
		ovf:   nil,
		hash:  hasher,
		equal: comparator,
	}
}

func (m Custom[K, V]) Get(key K) (value V, ok bool) {
	hash := m.hash(key)
	main, foundMain := m.main[hash]
	if !foundMain {
		return main.Value, false // not found in main
	}
	if m.equal(main.Key, key) {
		return main.Value, true // found in main
	}
	overflow, foundOverflow := m.ovf[hash]
	if !foundOverflow {
		return main.Value, false // overflow not available
	}
	for _, entry := range overflow {
		if m.equal(entry.Key, key) {
			return entry.Value, true // found in overflow
		}
	}
	return main.Value, false // not found in overflow too
}

func (m Custom[K, V]) Set(key K, value V) (added bool) {
	hash := m.hash(key)
	mainEntry, foundMain := m.main[hash]
	if !foundMain || m.equal(mainEntry.Key, key) {
		m.main[hash] = NewEntry(key, value)
		return !foundMain // new in main OR update existing key in main
	}
	overflow, foundOverflow := m.ovf[hash]
	if !foundOverflow {
		if m.ovf == nil {
			m.ovf = map[uint64][]Entry[K, V]{}
		}
		m.ovf[hash] = []Entry[K, V]{NewEntry(key, value)}
		return true // add new to overflow
	}
	for i, entry := range overflow {
		if m.equal(entry.Key, key) {
			overflow[i] = NewEntry(key, value)
			return false // update existing key in overflow
		}
	}
	m.ovf[hash] = append(overflow, NewEntry(key, value))
	return true // add new entry to overflow
}

func (m Custom[K, V]) Delete(key K) (ok bool) {
	hash := m.hash(key)
	main, foundMain := m.main[hash]
	if !foundMain {
		return false // not found in main
	}
	if m.equal(main.Key, key) {
		delete(m.main, hash)
		return true // found in main
	}
	overflow, foundOverflow := m.ovf[hash]
	if !foundOverflow {
		return false // overflow not available
	}
	for i, entry := range overflow {
		if m.equal(entry.Key, key) {
			m.ovf[hash] = xslice.UnorderedDelete(overflow, i)
			return true // found in overflow
		}
	}
	return false // not found in overflow too
}
