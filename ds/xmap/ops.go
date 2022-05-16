package xmap

func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	var entries []Entry[K, V]
	for k, v := range m {
		entries = append(entries, Entry[K, V]{Key: k, Value: v})
	}
	return entries
}

func Keys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	var values []V
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func Delete[K comparable, V any](m map[K]V, key K) {
	delete(m, key)
}

func Clear[K comparable, V any](m map[K]V) {
	// Go 1.15 optimizes this loop into runtime.mapclear()
	// Might not work for interface key type
	// https://github.com/golang/go/issues/45328
	for k := range m {
		delete(m, k)
	}
}
