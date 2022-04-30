package slice

import (
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xtd"
)

// Iterator implements iter.RandomIterator
var _ iter.RandomIterator[int, Iterator[int]] = Iterator[int]{}

// Iterator represents a slice iterator
type Iterator[T any] struct {
	_ xtd.NoCompare

	pos int
	s   Slice[T]
}

func (iter Iterator[T]) Ref() *T {
	return &iter.s[iter.pos]
}

func (iter Iterator[T]) Get() T {
	return iter.s[iter.pos]
}

func (iter Iterator[T]) Set(val T) {
	iter.s[iter.pos] = val
}

func (iter Iterator[T]) Next() Iterator[T] {
	iter.Inc()
	return iter
}

func (iter Iterator[T]) Prev() Iterator[T] {
	iter.Dec()
	return iter
}

func (iter Iterator[T]) Add(offset int) Iterator[T] {
	iter.Advance(offset)
	return iter
}

func (iter *Iterator[T]) Inc() {
	if iter.pos >= len(iter.s) {
		panic("increment out of bound")
	}
	iter.pos++
}

func (iter *Iterator[T]) Dec() {
	if iter.pos <= 0 {
		panic("decrement out of bound")
	}
	iter.pos--
}

func (iter *Iterator[T]) Advance(offset int) {
	k := iter.pos + offset
	if k < 0 || len(iter.s) < k {
		panic("offset out of bound")
	}
	iter.pos = k
}

func (iter Iterator[T]) Pos() int {
	return iter.pos
}

func (iter Iterator[T]) Equal(other Iterator[T]) bool {
	if !ReferenceEqual(iter.s, other.s) {
		panic("compare iterator of different slices")
	}
	return other.pos == iter.pos
}
