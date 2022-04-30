package iter

import (
	"github.com/ndkimhao/go-xtd/generics"
)

type DummyIterator[T any] struct{}

func Dummy[T any]() DummyIterator[T] {
	return DummyIterator[T]{}
}

func (it DummyIterator[T]) Next() DummyIterator[T] {
	return it
}

func (it DummyIterator[T]) Get() T {
	return generics.ZeroOf[T]()
}

func (it DummyIterator[T]) Equal(other DummyIterator[T]) bool {
	return true
}

func (it DummyIterator[T]) Prev() DummyIterator[T] {
	return it
}

func (it DummyIterator[T]) Add(offset int) DummyIterator[T] {
	return it
}

func (it DummyIterator[T]) Pos() int {
	return 0
}

func (it DummyIterator[T]) Set(value T) {
}
