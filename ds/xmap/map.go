package xmap

type Entry[K any, V any] struct {
	Key   K
	Value V
}

func NewEntry[K, V any](key K, value V) Entry[K, V] {
	return Entry[K, V]{Key: key, Value: value}
}

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() Map[K, V] {
	return map[K]V{}
}

func (m Map[K, V]) Entries() []Entry[K, V] {
	return Entries(m)
}

func (m Map[K, V]) Keys() []K {
	return Keys(m)
}

func (m Map[K, V]) Values() []V {
	return Values(m)
}

func (m Map[K, V]) Delete(key K) {
	delete(m, key)
}

func (m Map[K, V]) Map() map[K]V {
	return m
}

func (m *Map[K, V]) Clear() {
	*m = map[K]V{}
}
