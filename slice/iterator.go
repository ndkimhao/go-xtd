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

func (iter Iterator[T]) Ref() *T {
	if iter.pos < 0 || iter.len <= iter.pos {
		panic("ref out of bound")
	}
	return (*T)(unsafe.Add(unsafe.Pointer((*T)(iter.beg)), uintptr(iter.pos)*unsafe.Sizeof(*(*T)(nil))))
}

func (iter Iterator[T]) Value() T {
	return *iter.Ref()
}

func (iter Iterator[T]) SetValue(val T) {
	*iter.Ref() = val
}

func (iter Iterator[T]) addUnchecked(offset int) Iterator[T] {
	return Iterator[T]{pos: iter.pos + offset, len: iter.len, beg: iter.beg}
}

func (iter Iterator[T]) Next() Iterator[T] {
	if iter.pos >= iter.len {
		panic("next is out of bound")
	}
	return iter.addUnchecked(1)
}

func (iter Iterator[T]) Prev() Iterator[T] {
	if iter.pos <= 0 {
		panic("prev is out of bound")
	}
	return iter.addUnchecked(-1)
}

func (iter Iterator[T]) Add(offset int) Iterator[T] {
	k := iter.pos + offset
	if k < 0 || iter.len < k {
		panic("add out of bound")
	}
	return iter.addUnchecked(offset)
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
