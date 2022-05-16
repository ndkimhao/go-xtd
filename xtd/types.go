package xtd

type Comparable[T any] interface {
	Equal(other T) bool
}

type Ordered[T any] interface {
	// CompareTo returns a negative integer, zero, or a positive integer
	// as this object is less than, equal to, or greater than the specified object.
	CompareTo(other T) int
}
