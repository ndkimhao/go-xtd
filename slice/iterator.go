package slice

import (
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xtd"
)

// Iterator is a implementation of iter.RandomAccessIterator
var _ iter.RandomAccessIterator[int, Iterator[int]] = Iterator[int]{}

// Iterator represents a slice iterator
type Iterator[T any] struct {
	_ xtd.NoCompare

	s Slice[T]
	p int
}

func (iter Iterator[T]) Value() T {
	return iter.s.At(iter.p)
}

func (iter Iterator[T]) SetValue(val T) {
	iter.s.Set(iter.p, val)
}

func (iter Iterator[T]) Next() Iterator[T] {
	if iter.p >= iter.s.Len() {
		panic("Iterator.Next: past end of slice")
	}
	return Iterator[T]{s: iter.s, p: iter.p + 1}
}

func (iter Iterator[T]) Prev() Iterator[T] {
	if iter.p <= 0 {
		panic("Iterator.Prev: past start of slice")
	}
	return Iterator[T]{s: iter.s, p: iter.p - 1}
}

func (iter Iterator[T]) Add(offset int) Iterator[T] {
	k := iter.p + offset
	if k < 0 && iter.s.Len() < k {
		panic("Iterator.Add: out of bound")
	}
	return Iterator[T]{s: iter.s, p: k}
}

func (iter Iterator[T]) Position() int {
	return iter.p
}

func (iter Iterator[T]) Equal(other Iterator[T]) bool {
	if !ReferenceEqual(iter.s, other.s) {
		panic("Iterator.Equal: compare iterator of different slices")
	}
	return other.p == iter.p
}
