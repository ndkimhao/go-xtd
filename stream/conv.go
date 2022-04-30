package stream

func OfSlice[T any](slice []T) *Stream[T] {
	return New[T](&sliceSource[T]{a: slice})
}

func Of[T any](values ...T) *Stream[T] {
	return OfSlice(values)
}
