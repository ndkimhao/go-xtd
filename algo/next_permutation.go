package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xfn"
)

func NextPermutation[T any, It iter.RandomIterator[T, It]](first, last It, comp xfn.Comparator[T]) bool {
	rFirst := iter.ReverseRandomIterator[T](last)
	rLast := iter.ReverseRandomIterator[T](first)
	left := IsSortedUntil(rFirst, rLast, comp)
	if !left.Equal(rLast) {
		right := UpperBound[T](rFirst, left, left.Value(), comp)
		Swap[T](left, right)
	}
	Reverse[T](left.Base(), last)
	return !left.Equal(rLast)
}

func NextPermutationOrdered[T constraints.Ordered, It iter.RandomIterator[T, It]](first, last It) bool {
	return NextPermutation(first, last, xfn.Less[T])
}
