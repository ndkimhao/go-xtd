package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/iterator"
	"github.com/ndkimhao/go-xtd/xfn"
)

//NextPermutation transform range [first last) to next permutation,return true if success, or false if failure
func NextPermutation[T constraints.Ordered](first, last iterator.RandomAccessIterator[T]) bool {
	return NextPermutationCustom(first, last, xfn.ComparatorOf[T])
}

func NextPermutationCustom[T constraints.Ordered](first, last iterator.RandomAccessIterator[T], cmp xfn.Comparator[T]) bool {
	rangeLen := last.Position() - first.Position()
	endPos := first.Position() + rangeLen - 1
	cur := endPos
	pre := cur - 1

	endIter := first.IteratorAt(endPos)
	for cur > first.Position() && cmp(first.IteratorAt(pre).Value(), first.IteratorAt(cur).Value()) >= 0 {
		cur--
		pre--
	}
	if cur <= first.Position() {
		reverse(first, endIter)
		return false
	}
	for cur = endPos; cmp(first.IteratorAt(cur).Value(),
		first.IteratorAt(pre).Value()) <= 0; cur-- {
	}
	Swap[T](first.IteratorAt(cur), first.IteratorAt(pre))
	reverse[T](first.IteratorAt(pre+1), endIter)
	return true
}

func reverse[T any](s, e iterator.RandomAccessIterator[T]) {
	for s.Position() < e.Position() {
		Swap[T](s, e)
		s.Next()
		e.Prev()
	}
}
