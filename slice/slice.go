package slice

import (
	"github.com/ndkimhao/go-xtd/iter"
)

type Slice[T any] []T

func New[T any]() Slice[T] {
	return nil
}

func NewLen[T any](len int) Slice[T] {
	return make([]T, len)
}

func NewLenCap[T any](len, cap int) Slice[T] {
	return make([]T, cap)
}

func Of[T any](values ...T) Slice[T] {
	return values
}

func OfSlice[T any](values []T) Slice[T] {
	return values
}

func Copy[T any](values []T) Slice[T] {
	if len(values) == 0 {
		return nil
	}
	return append([]T(nil), values...)
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

func (s *Slice[T]) Insert(i int, value T) {
	*s = Insert(*s, i, value)
}

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) Cap() int {
	return cap(s)
}

func (s Slice[T]) Empty() bool {
	return len(s) == 0
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

func (s Slice[T]) IteratorAt(pos int) Iterator[T] {
	slen := s.Len()
	if pos < 0 || slen < pos {
		panic("out of bound")
	}
	var beg *T
	if slen > 0 {
		beg = &s[0]
	}
	return Iterator[T]{pos: pos, len: slen, beg: beg}
}

func (s Slice[T]) Begin() Iterator[T] {
	return s.IteratorAt(0)
}

func (s Slice[T]) End() Iterator[T] {
	return s.IteratorAt(s.Len())
}

func (s Slice[T]) RBegin() iter.ReverseRandom[T, Iterator[T]] {
	return iter.MakeReverseRandom[T](s.End())
}

func (s Slice[T]) REnd() iter.ReverseRandom[T, Iterator[T]] {
	return iter.MakeReverseRandom[T](s.Begin())
}

func (s Slice[T]) Range() iter.Range[T, Iterator[T]] {
	return iter.MakeRange[T](s.Begin(), s.End())
}

func (s Slice[T]) SubRange(first, last int) iter.Range[T, Iterator[T]] {
	return iter.MakeRange[T](s.IteratorAt(first), s.IteratorAt(last))
}

func (s Slice[T]) ReverseRange() iter.Range[T, iter.ReverseRandom[T, Iterator[T]]] {
	return iter.MakeRange[T](s.RBegin(), s.REnd())
}

func (s Slice[T]) ReverseSubRange(first, last int) iter.Range[T, iter.ReverseRandom[T, Iterator[T]]] {
	return iter.MakeRange[T](s.RBegin().Add(first), s.RBegin().Add(last))
}

func (s Slice[T]) Reversed() Slice[T] {
	if len(s) == 0 {
		return nil
	}
	r := make([]T, len(s))
	last := len(s) - 1
	for i, v := range s {
		r[last-i] = v
	}
	return r
}
