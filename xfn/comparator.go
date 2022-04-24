package xfn

import (
	"github.com/ndkimhao/go-xtd/constraints"
)

// Comparator returns true if a < b
type Comparator[T any] func(T, T) bool

// FullComparator returns a negative integer, zero, or a positive integer
// as the first object is less than, equal to, or greater than the second object.
type FullComparator[T any] func(T, T) int

func FullComparatorOf[T constraints.Ordered](a, b T) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
