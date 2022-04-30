package xmap

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return map[T]struct{}{}
}

func (s Set[T]) Add(k T) {
	s[k] = struct{}{}
}

func (s Set[T]) TryAdd(k T) (added bool) {
	if _, found := s[k]; found {
		return false
	}
	s[k] = struct{}{}
	return true
}

func (s Set[T]) Delete(k T) {
	delete(s, k)
}

func (s Set[T]) TryDelete(k T) (deleted bool) {
	if _, found := s[k]; !found {
		return false
	}
	delete(s, k)
	return true
}

func (s Set[T]) Has(k T) bool {
	_, found := s[k]
	return found
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s *Set[T]) Clear() {
	*s = map[T]struct{}{}
}

func (s Set[T]) Raw() map[T]struct{} {
	return s
}

func (s Set[T]) Values() []T {
	return Keys(s)
}
