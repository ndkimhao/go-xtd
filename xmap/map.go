package xmap

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Entries() []Entry[K, V] {
	var entries []Entry[K, V]
	for k, v := range m {
		entries = append(entries, Entry[K, V]{Key: k, Value: v})
	}
	return entries
}

func (m Map[K, V]) Keys() []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Map[K, V]) Values() []V {
	var values []V
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func (m Map[K, V]) Delete(key K) {
	delete(m, key)
}

func (m Map[K, V]) Map() map[K]V {
	return m
}
