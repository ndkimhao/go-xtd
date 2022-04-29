package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/xfn"
)

func UpperBoundIterators[T any, It iter.ConstRandomIterator[T, It]](first, last It, value T, comp xfn.Comparator[T]) It {
	count := iter.Distance(first, last)
	if count < 0 {
		panic("invalid range")
	}
	for count > 0 {
		it := first
		step := count / 2
		it = it.Add(step)
		if !comp(value, it.Get()) {
			it = it.Next()
			first = it
			count -= step + 1
		} else {
			count = step
		}
	}
	return first
}

func LowerBoundIterators[T any, It iter.ConstRandomIterator[T, It]](first, last It, value T, comp xfn.Comparator[T]) It {
	count := iter.Distance(first, last)
	if count < 0 {
		panic("invalid range")
	}
	for count > 0 {
		it := first
		step := count / 2
		it = it.Add(step)
		if comp(it.Get(), value) {
			it = it.Next()
			first = it
			count -= step + 1
		} else {
			count = step
		}
	}
	return first
}

func UpperBound[T constraints.Ordered, It iter.ConstRandomIterator[T, It]](r iter.Range[T, It], value T) It {
	return UpperBoundIterators[T](r.Begin, r.End, value, xfn.Less[T])
}

func LowerBound[T constraints.Ordered, It iter.ConstRandomIterator[T, It]](r iter.Range[T, It], value T) It {
	return LowerBoundIterators[T](r.Begin, r.End, value, xfn.Less[T])
}

func UpperBoundAny[T any, It iter.ConstRandomIterator[T, It]](r iter.Range[T, It], value T, comp xfn.Comparator[T]) It {
	return UpperBoundIterators[T](r.Begin, r.End, value, comp)
}

func LowerBoundAny[T any, It iter.ConstRandomIterator[T, It]](r iter.Range[T, It], value T, comp xfn.Comparator[T]) It {
	return LowerBoundIterators[T](r.Begin, r.End, value, comp)
}
