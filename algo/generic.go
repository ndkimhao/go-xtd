package algo

import (
	"github.com/ndkimhao/go-xtd/ds/iter"
)

// Swap swaps the value of two iterator
func Swap[T any, It iter.Iterator[T, It]](a, b It) {
	if !a.Equal(b) {
		va := a.Get()
		vb := b.Get()
		a.Set(vb)
		b.Set(va)
	}
}

// Reverse the elements in the range [first, last]
func Reverse[T any, It iter.BidirIterator[T, It]](first, last It) {
	for !first.Equal(last) {
		last = last.Prev()
		if first.Equal(last) {
			break
		}
		Swap[T](first, last)
		first = first.Next()
	}
}

//// Count returns the number of elements that their value is equal to value in range [first, last)
//func Count[T comparable](first, last iter.ConstIterator[T], value T) int {
//	var count int
//	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
//		if iter.Get() == value {
//			count++
//		}
//	}
//	return count
//}
//
//// CountIf returns the number of elements are satisfied the function f in range [first, last)
//func CountIf[T any](first, last iter.ConstIterator[T], f xfn.Predicate[iter.ConstIterator[T]]) int {
//	var count int
//	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
//		if f(iter) {
//			count++
//		}
//	}
//	return count
//}
//
//// Find finds the first element that its value is equal to value in range [first, last), and returns its iterator, or last if not found
//func Find[T comparable](first, last iter.ConstIterator[T], value T) iter.ConstIterator[T] {
//	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
//		if iter.Get() == value {
//			return iter
//		}
//	}
//	return last
//}
//
//// FindIf finds the first element that is satisfied the function f, and returns its iterator, or last if not found
//func FindIf[T any](first, last iter.ConstIterator[T], f xfn.Predicate[iter.ConstIterator[T]]) iter.ConstIterator[T] {
//	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
//		if f(iter) {
//			return iter
//		}
//	}
//	return last
//}
//
//// MaxElement returns an Iterator to the largest element in the range [first, last). If several elements in the range are equivalent to the largest element, returns the iterator to the first such element. Returns last if the range is empty.
//func MaxElement[T constraints.Ordered](first, last iter.ConstIterator[T]) iter.ConstIterator[T] {
//	largest := first
//	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
//		if iter.Get() > largest.Get() {
//			largest = iter.Clone()
//		}
//	}
//	return largest
//}
//
//// MinElement returns an Iterator to the smallest element value in the range [first, last). If several elements in the range are equivalent to the smallest element value, returns the iterator to the first such element. Returns last if the range is empty.
//func MinElement[T constraints.Ordered](first, last iter.ConstIterator[T]) iter.ConstIterator[T] {
//	smallest := first
//	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
//		if iter.Get() < smallest.Get() {
//			smallest = iter.Clone()
//		}
//	}
//	return smallest
//}
