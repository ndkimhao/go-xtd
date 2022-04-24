package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xfn"
)

func IsSortedUntil[T constraints.Ordered, It iter.ConstIterator[T, It]](first, last It) It {
	return IsSortedUntilCustom[T, It](first, last, xfn.LessComparatorOf[T])
}

func IsSortedUntilCustom[T any, It iter.ConstIterator[T, It]](first, last It, comp xfn.LessComparator[T]) It {
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
