package slice

import (
	"github.com/ndkimhao/go-xtd/iter"
)

type Slice[T any] []T

func New[T any]() Slice[T] {
	return nil
}

func Of[T any](values ...T) Slice[T] {
	return values
}

func OfSlice[T any](values []T) Slice[T] {
	return values
}

func (s *Slice[T]) Append(value T) {
	*s = append(*s, value)
}

func (s *Slice[T]) AppendMany(values ...T) {
	*s = append(*s, values...)
}

func (s *Slice[T]) Delete(i int) {
	*s = Delete(*s, i)
}

func (s *Slice[T]) DeleteLast() {
	*s = DeleteLast(*s)
}

func (s *Slice[T]) UnorderedDelete(i int) {
	*s = UnorderedDelete(*s, i)
}

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) Cap() int {
	return cap(s)
}

func (s Slice[T]) Sub(start, end int) Slice[T] {
	return s[start:end]
}

func (s Slice[T]) Slice() []T {
	return s
}

func (s Slice[T]) At(i int) T {
	return s[i]
}

func (s Slice[T]) Set(i int, x T) {
	s[i] = x
}

func (s Slice[T]) First() T {
	return s[0]
}

func (s Slice[T]) Last() T {
	return s[len(s)-1]
}

func (s Slice[T]) Begin() Iterator[T] {
	return Iterator[T]{s: s, p: 0}
}

func (s Slice[T]) End() Iterator[T] {
	return Iterator[T]{s: s, p: s.Len()}
}

func (s Slice[T]) RBegin() iter.ReverseRandom[T, Iterator[T]] {
	return iter.ReverseRandomIterator[T](s.End())
}

func (s Slice[T]) REnd() iter.ReverseRandom[T, Iterator[T]] {
	return iter.ReverseRandomIterator[T](s.Begin())
}
