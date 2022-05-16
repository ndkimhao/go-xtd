package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/ds/iter"
	"github.com/ndkimhao/go-xtd/xfn"
)

func IsSortedUntilIterators[T any, It iter.ConstIterator[T, It]](first, last It, comp xfn.Comparator[T]) It {
	if !first.Equal(last) {
		next := first.Next()
		for !next.Equal(last) {
			if comp(next.Get(), first.Get()) {
				return next
			}
			first = next
			next = next.Next()
		}
	}
	return last
}

func IsSortedIterators[T any, It iter.ConstIterator[T, It]](first, last It, comp xfn.Comparator[T]) bool {
	return IsSortedUntilIterators(first, last, comp).Equal(last)
}

func IsSortedUntil[T constraints.Ordered, It iter.ConstIterator[T, It]](r iter.Range[T, It]) It {
	return IsSortedUntilIterators[T, It](r.Begin, r.End, xfn.Less[T])
}

func IsSorted[T constraints.Ordered, It iter.ConstIterator[T, It]](r iter.Range[T, It]) bool {
	return IsSortedIterators(r.Begin, r.End, xfn.Less[T])
}

func IsSortedUntilAny[T any, It iter.ConstIterator[T, It]](r iter.Range[T, It], comp xfn.Comparator[T]) It {
	return IsSortedUntilIterators[T, It](r.Begin, r.End, comp)
}

func IsSortedAny[T any, It iter.ConstIterator[T, It]](r iter.Range[T, It], comp xfn.Comparator[T]) bool {
	return IsSortedIterators(r.Begin, r.End, comp)
}
