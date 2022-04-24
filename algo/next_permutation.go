package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xfn"
)

func NextPermutation[T constraints.Ordered, It iter.RandomIterator[T, It]](first, last It) bool {
	return NextPermutationComp(first, last, xfn.LessComparatorOf[T])
}

func NextPermutationComp[T any, It iter.RandomIterator[T, It]](first, last It, comp xfn.LessComparator[T]) bool {
	rFirst := iter.ReverseRandomIterator[T](last)
	rLast := iter.ReverseRandomIterator[T](first)
	left := IsSortedUntilComp(rFirst, rLast, comp)
	if !left.Equal(rLast) {
		right := UpperBoundComp[T](rFirst, left, left.Value(), comp)
		Swap[T](left, right)
	}
	Reverse[T](left.Base(), last)
	return !left.Equal(rLast)
}
