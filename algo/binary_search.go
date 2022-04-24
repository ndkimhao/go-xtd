package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xfn"
)

func UpperBound[T constraints.Ordered, It iter.ConstRandomIterator[T, It]](first, last It, value T) It {
	return UpperBoundComp(first, last, value, xfn.Less[T])
}

func UpperBoundComp[T any, It iter.ConstRandomIterator[T, It]](first, last It, value T, comp xfn.Comparator[T]) It {
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

func LowerBound[T constraints.Ordered, It iter.ConstRandomIterator[T, It]](first, last It, value T) It {
	return LowerBoundComp(first, last, value, xfn.Less[T])
}

func LowerBoundComp[T any, It iter.ConstRandomIterator[T, It]](first, last It, value T, comp xfn.Comparator[T]) It {
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
