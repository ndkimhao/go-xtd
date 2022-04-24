package slice

import (
	"github.com/ndkimhao/go-xtd/iterator"
)

// Iterator is a implementation of iterator.RandomAccessIterator
var _ iterator.RandomAccessIterator[int] = (*Iterator[int])(nil)

// Iterator represents a slice iterator
type Iterator[T any] struct {
	s Slice[T]
	p int
}

// IsValid returns trus if the iterator is valid, othterwise return false
func (iter *Iterator[T]) IsValid() bool {
	if iter.p >= 0 && iter.p < iter.s.Len() {
		return true
	}
	return false
}

// Value returns the value of the iterator point to
func (iter *Iterator[T]) Value() T {
	return iter.s.At(iter.p)
}

// SetValue sets the value of the iterator point to
func (iter *Iterator[T]) SetValue(val T) {
	iter.s.Set(iter.p, val)
}

// Next moves the iterator's position to the next position, and returns itself
func (iter *Iterator[T]) Next() iterator.ConstIterator[T] {
	if iter.p < iter.s.Len() {
		iter.p++
	}
	return iter
}

// Prev move the iterator's position to the previous position, and return itself
func (iter *Iterator[T]) Prev() iterator.ConstBidIterator[T] {
	if iter.p >= 0 {
		iter.p--
	}
	return iter
}

// Clone clones the iterator into a new one
func (iter *Iterator[T]) Clone() iterator.ConstIterator[T] {
	return &Iterator[T]{s: iter.s, p: iter.p}
}

// IteratorAt creates an iterator with the passed position
func (iter *Iterator[T]) IteratorAt(position int) iterator.RandomAccessIterator[T] {
	return &Iterator[T]{s: iter.s, p: position}
}

// Position returns the position of the iterator
func (iter *Iterator[T]) Position() int {
	return iter.p
}

// Equal returns true if the iterator is equal to the passed iterator
func (iter *Iterator[T]) Equal(other iterator.ConstIterator[T]) bool {
	if otherIter, ok := other.(*Iterator[T]); !ok {
		return otherIter.p == iter.p
	}
	return false
}
