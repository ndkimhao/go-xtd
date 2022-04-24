package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xfn"
)

func IsSortedUntil[T constraints.Ordered, It iter.ConstIterator[T, It]](first, last It) It {
	return IsSortedUntilComp[T, It](first, last, xfn.Less[T])
}

func IsSortedUntilComp[T any, It iter.ConstIterator[T, It]](first, last It, comp xfn.Comparator[T]) It {
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

func IsSorted[T constraints.Ordered, It iter.ConstIterator[T, It]](first, last It) bool {
	return IsSortedUntilComp(first, last, xfn.Less[T]).Equal(last)
}

func IsSortedComp[T any, It iter.ConstIterator[T, It]](first, last It, comp xfn.Comparator[T]) bool {
	return IsSortedUntilComp(first, last, comp).Equal(last)
}
