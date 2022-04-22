package stream

import (
	"github.com/ndkimhao/go-xtd/vec"
	"github.com/ndkimhao/go-xtd/xmap"
)

func OfSlice[T any](slice []T) *Stream[T] {
	return New[T](&sliceIter[T]{a: slice})
}

func Of[T any](values ...T) *Stream[T] {
	return OfSlice(values)
}

func OfVec[T any](v vec.Vector[T]) *Stream[T] {
	return OfSlice(v.Slice())
}

func OfMap[K comparable, V any](m map[K]V) *Stream[xmap.Entry[K, V]] {
	return OfSlice(xmap.Entries(m))
}
