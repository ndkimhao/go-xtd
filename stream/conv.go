package stream

import (
	"github.com/ndkimhao/go-xtd/vec"
)

func OfSlice[T any](slice []T) *Stream[T] {
	return New[T](&sliceIter[T]{a: slice})
}

func Of[T any](values ...T) *Stream[T] {
	return OfSlice(values)
}

func OfVec[T any](v vec.Vector[T]) *Stream[T] {
	return New[T](OfSlice(v.Slice()))
}
