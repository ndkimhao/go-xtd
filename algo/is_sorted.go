package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xfn"
)

func IsSortedUntil[T any, It iter.ConstIterator[T, It]](first, last It, comp xfn.Comparator[T]) It {
	if !first.Equal(last) {
		next := first.Next()
		for !next.Equal(last) {
			if comp(next.Value(), first.Value()) {
				return next
			}
			first = next
			next = next.Next()
		}
	}
	return last
}

func IsSorted[T any, It iter.ConstIterator[T, It]](first, last It, comp xfn.Comparator[T]) bool {
	return IsSortedUntil(first, last, comp).Equal(last)
}

func IsSortedUntilOrdered[T constraints.Ordered, It iter.ConstIterator[T, It]](first, last It) It {
	return IsSortedUntil[T, It](first, last, xfn.Less[T])
}

func IsSortedOrdered[T constraints.Ordered, It iter.ConstIterator[T, It]](first, last It) bool {
	return IsSortedUntil(first, last, xfn.Less[T]).Equal(last)
}
