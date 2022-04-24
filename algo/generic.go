package algo

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/iterator"
	"github.com/ndkimhao/go-xtd/xfn"
)

// Swap swaps the value of two iterator
func Swap[T any](a, b iterator.Iterator[T]) {
	va := a.Value()
	vb := b.Value()
	a.SetValue(vb)
	b.SetValue(va)
}

// Reverse reverse the elements in the range [first, last]
func Reverse[T any](first, last iterator.BidIterator[T]) {
	left := first.Clone().(iterator.BidIterator[T])
	right := last.Clone().(iterator.BidIterator[T]).Prev().(iterator.BidIterator[T])
	for !left.Equal(right) {
		Swap[T](left, right)
		left.Next()
		if left.Equal(right) {
			break
		}
		right.Prev()
	}
}

// Count returns the number of elements that their value is equal to value in range [first, last)
func Count[T comparable](first, last iterator.ConstIterator[T], value T) int {
	var count int
	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
		if iter.Value() == value {
			count++
		}
	}
	return count
}

// CountIf returns the number of elements are satisfied the function f in range [first, last)
func CountIf[T any](first, last iterator.ConstIterator[T], f xfn.Predicate[iterator.ConstIterator[T]]) int {
	var count int
	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
		if f(iter) {
			count++
		}
	}
	return count
}

// Find finds the first element that its value is equal to value in range [first, last), and returns its iterator, or last if not found
func Find[T comparable](first, last iterator.ConstIterator[T], value T) iterator.ConstIterator[T] {
	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
		if iter.Value() == value {
			return iter
		}
	}
	return last
}

// FindIf finds the first element that is satisfied the function f, and returns its iterator, or last if not found
func FindIf[T any](first, last iterator.ConstIterator[T], f xfn.Predicate[iterator.ConstIterator[T]]) iterator.ConstIterator[T] {
	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
		if f(iter) {
			return iter
		}
	}
	return last
}

// MaxElement returns an Iterator to the largest element in the range [first, last). If several elements in the range are equivalent to the largest element, returns the iterator to the first such element. Returns last if the range is empty.
func MaxElement[T constraints.Ordered](first, last iterator.ConstIterator[T]) iterator.ConstIterator[T] {
	largest := first
	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
		if iter.Value() > largest.Value() {
			largest = iter.Clone()
		}
	}
	return largest
}

// MinElement returns an Iterator to the smallest element value in the range [first, last). If several elements in the range are equivalent to the smallest element value, returns the iterator to the first such element. Returns last if the range is empty.
func MinElement[T constraints.Ordered](first, last iterator.ConstIterator[T]) iterator.ConstIterator[T] {
	smallest := first
	for iter := first.Clone(); !iter.Equal(last); iter.Next() {
		if iter.Value() < smallest.Value() {
			smallest = iter.Clone()
		}
	}
	return smallest
}
