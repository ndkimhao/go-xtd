package slice

import (
	"unsafe"

	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xtd"
)

// Iterator is a implementation of iter.RandomIterator
var _ iter.RandomIterator[int, Iterator[int]] = Iterator[int]{}

// Iterator represents a slice iterator
type Iterator[T any] struct {
	_ xtd.NoCompare

	pos int
	len int
	beg *T
}

func (iter Iterator[T]) ValueRef() *T {
	if iter.pos < 0 || iter.len <= iter.pos {
		panic("ref out of bound")
	}
	return (*T)(unsafe.Add(unsafe.Pointer((*T)(iter.beg)), uintptr(iter.pos)*unsafe.Sizeof(*(*T)(nil))))
}

func (iter Iterator[T]) Value() T {
	return *iter.ValueRef()
}

func (iter Iterator[T]) SetValue(val T) {
	*iter.ValueRef() = val
}

func (iter Iterator[T]) Next() Iterator[T] {
	iter.Incr()
	return iter
}

func (iter Iterator[T]) Prev() Iterator[T] {
	iter.Decr()
	return iter
}

func (iter Iterator[T]) Add(offset int) Iterator[T] {
	iter.Advance(offset)
	return iter
}

func (iter *Iterator[T]) Incr() {
	if iter.pos >= iter.len {
		panic("increment out of bound")
	}
	iter.pos++
}

func (iter *Iterator[T]) Decr() {
	if iter.pos <= 0 {
		panic("decrement out of bound")
	}
	iter.pos--
}

func (iter *Iterator[T]) Advance(offset int) {
	k := iter.pos + offset
	if k < 0 || iter.len < k {
		panic("offset out of bound")
	}
	iter.pos = k
}

func (iter Iterator[T]) Position() int {
	return iter.pos
}

func (iter Iterator[T]) Equal(other Iterator[T]) bool {
	if iter.beg != other.beg {
		panic("compare iterator of different slices")
	}
	return other.pos == iter.pos
}
