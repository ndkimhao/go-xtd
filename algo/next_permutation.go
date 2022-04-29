package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xfn"
)

func NextPermutationIterators[T any, It iter.RandomIterator[T, It]](first, last It, comp xfn.Comparator[T]) bool {
	rFirst := iter.ReverseRandom[T](last)
	rLast := iter.ReverseRandom[T](first)
	left := IsSortedUntil(rFirst, rLast, comp)
	if !left.Equal(rLast) {
		right := UpperBoundIterators[T](rFirst, left, left.Value(), comp)
		Swap[T](left, right)
	}
	Reverse[T](left.Base(), last)
	return !left.Equal(rLast)
}

func NextPermutation[T constraints.Ordered, It iter.RandomIterator[T, It]](r iter.Range[T, It]) bool {
	return NextPermutationIterators(r.Begin, r.End, xfn.Less[T])
}

func NextPermutationAny[T any, It iter.RandomIterator[T, It]](r iter.Range[T, It], comp xfn.Comparator[T]) bool {
	return NextPermutationIterators(r.Begin, r.End, comp)
}
