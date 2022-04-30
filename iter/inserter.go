package iter

import (
	"github.com/ndkimhao/go-xtd/generics"
)

// AppendIterator implements Iterator
var _ Iterator[int, AppendIterator[int]] = AppendIterator[int]{}

type AppendIterator[T any] struct {
	a Appender[T]
}

type Appender[T any] interface {
	Append(value T)
}

func Append[T any](a Appender[T]) AppendIterator[T] {
	return AppendIterator[T]{a: a}
}

func (it AppendIterator[T]) Next() AppendIterator[T] {
	return it
}

func (it AppendIterator[T]) Get() T {
	return generics.ZeroOf[T]()
}

func (it AppendIterator[T]) Equal(other AppendIterator[T]) bool {
	return it.a == other.a
}

func (it AppendIterator[T]) Set(value T) {
	it.a.Append(value)
}

// PrependIterator implements Iterator
var _ Iterator[int, PrependIterator[int]] = PrependIterator[int]{}

type PrependIterator[T any] struct {
	p Prepender[T]
}

type Prepender[T any] interface {
	Prepend(value T)
}

func Prepend[T any](p Prepender[T]) PrependIterator[T] {
	return PrependIterator[T]{p: p}
}

func (it PrependIterator[T]) Next() PrependIterator[T] {
	return it
}

func (it PrependIterator[T]) Get() T {
	return generics.ZeroOf[T]()
}

func (it PrependIterator[T]) Equal(other PrependIterator[T]) bool {
	return it.p == other.p
}

func (it PrependIterator[T]) Set(value T) {
	it.p.Prepend(value)
}

// InsertIterator implements Iterator
var _ Iterator[int, *InsertIterator[int, DummyIterator[int]]] = &InsertIterator[int, DummyIterator[int]]{}

type InsertIterator[T any, It ConstIterator[T, It]] struct {
	ins Inserter[T, It]
	p   It
}

type Inserter[T any, It ConstIterator[T, It]] interface {
	Insert(pos It, value T) It
}

func Insert[T any, It ConstIterator[T, It]](pos It, i Inserter[T, It]) *InsertIterator[T, It] {
	return &InsertIterator[T, It]{ins: i, p: pos}
}

func (it *InsertIterator[T, It]) Next() *InsertIterator[T, It] {
	return it
}

func (it *InsertIterator[T, It]) Get() T {
	return generics.ZeroOf[T]()
}

func (it *InsertIterator[T, It]) Equal(other *InsertIterator[T, It]) bool {
	return it.ins == other.ins
}

func (it *InsertIterator[T, It]) Set(value T) {
	it.p = it.ins.Insert(it.p, value).Next()
}
