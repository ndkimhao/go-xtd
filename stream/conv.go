package stream

import (
	xmap2 "github.com/ndkimhao/go-xtd/ds/xmap"
	"github.com/ndkimhao/go-xtd/ds/xslice"
)

func OfSlice[T any](slice []T) *Stream[T] {
	return New[T](&sliceIter[T]{a: slice})
}

func Of[T any](values ...T) *Stream[T] {
	return OfSlice(values)
}

func OfVec[T any](v xslice.Slice[T]) *Stream[T] {
	return OfSlice(v.Slice())
}

func OfMap[K comparable, V any](m map[K]V) *Stream[xmap2.Entry[K, V]] {
	return OfSlice(xmap2.Entries(m))
}
