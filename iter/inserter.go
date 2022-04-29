package iter

import (
	"github.com/ndkimhao/go-xtd/generics"
)

type Appender[T any] interface {
	Append(value T)
}

func Append[T any](a Appender[T]) AppendIterator[T] {
	return AppendIterator[T]{a: a}
}

type AppendIterator[T any] struct {
	a Appender[T]
}

func (it AppendIterator[T]) Next() AppendIterator[T] {
	return it
}

func (it AppendIterator[T]) Value() T {
	return generics.ZeroOf[T]()
}

func (it AppendIterator[T]) Equal(other AppendIterator[T]) bool {
	return it.a == other.a
}

func (it AppendIterator[T]) SetValue(value T) {
	it.a.Append(value)
}

type Prepender[T any] interface {
	Prepend(value T)
}

type PrependIterator[T any] struct {
	p Prepender[T]
}

func Prepend[T any](p Prepender[T]) PrependIterator[T] {
	return PrependIterator[T]{p: p}
}

func (it PrependIterator[T]) Next() PrependIterator[T] {
	return it
}

func (it PrependIterator[T]) Value() T {
	return generics.ZeroOf[T]()
}

func (it PrependIterator[T]) Equal(other PrependIterator[T]) bool {
	return it.p == other.p
}

func (it PrependIterator[T]) SetValue(value T) {
	it.p.Prepend(value)
}
