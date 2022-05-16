package xsort

import (
	"math/bits"

	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/xfn"
)

type sortedHint int // hint for pdqsort when choosing the pivot

const (
	unknownHint sortedHint = iota
	increasingHint
	decreasingHint
)

// xorshift paper: https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
type xorshift uint64

func (r *xorshift) Next() uint64 {
	*r ^= *r << 13
	*r ^= *r >> 17
	*r ^= *r << 5
	return uint64(*r)
}

func nextPowerOfTwo(length int) uint {
	shift := uint(bits.Len(uint(length)))
	return uint(1 << shift)
}

func SortOrdered[T constraints.Ordered](data []T) {
	length := len(data)
	limit := bits.Len(uint(length))
	pdqsortOrdered(data, 0, length, limit)
}

func Sort[T any](data []T, less xfn.Comparator[T]) {
	length := len(data)
	limit := bits.Len(uint(length))
	pdqsortLessFunc(data, 0, length, limit, less)
}

func StableOrdered[T constraints.Ordered](data []T) {
	stableOrdered(data, len(data))
}

func Stable[T any](data []T, less xfn.Comparator[T]) {
	stableLessFunc(data, len(data), less)
}
