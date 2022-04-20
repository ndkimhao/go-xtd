package xtd

type Iterator[T any] interface {
	Get() T
	Has() bool
	Next() Iterator[T]
	Skip(n int) (it Iterator[T], skipped int)
}
