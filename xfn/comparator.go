package xfn

import (
	"github.com/ndkimhao/go-xtd/constraints"
)

// Comparator returns a negative integer, zero, or a positive integer
// as this object is less than, equal to, or greater than the specified object.
type Comparator[T any] func(T, T) int

func ComparatorOf[T constraints.Ordered](a, b T) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
