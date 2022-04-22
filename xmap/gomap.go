package xmap

func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	return Map[K, V](m).Entries()
}
