package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xfn"
)

func UpperBound[T any, It iter.ConstRandomIterator[T, It]](first, last It, value T, comp xfn.Comparator[T]) It {
	count := iter.Distance(first, last)
	if count < 0 {
		panic("invalid range")
	}
	for count > 0 {
		it := first
		step := count / 2
		it = it.Add(step)
		if !comp(value, it.Value()) {
			it = it.Next()
			first = it
			count -= step + 1
		} else {
			count = step
		}
	}
	return first
}

func LowerBound[T any, It iter.ConstRandomIterator[T, It]](first, last It, value T, comp xfn.Comparator[T]) It {
	count := iter.Distance(first, last)
	if count < 0 {
		panic("invalid range")
	}
	for count > 0 {
		it := first
		step := count / 2
		it = it.Add(step)
		if comp(it.Value(), value) {
			it = it.Next()
			first = it
			count -= step + 1
		} else {
			count = step
		}
	}
	return first
}

func UpperBoundOrdered[T constraints.Ordered, It iter.ConstRandomIterator[T, It]](first, last It, value T) It {
	return UpperBound(first, last, value, xfn.Less[T])
}

func LowerBoundOrdered[T constraints.Ordered, It iter.ConstRandomIterator[T, It]](first, last It, value T) It {
	return LowerBound(first, last, value, xfn.Less[T])
}
