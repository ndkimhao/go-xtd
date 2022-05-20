package priority_queue

import (
	"container/heap"
	"errors"
	"fmt"

	"github.com/ndkimhao/go-xtd/ds/xslice"
	"github.com/ndkimhao/go-xtd/generics"
)

type XHeapSlice[T any] struct {
	Comp func(T, T) bool
	V    []T
}

func (h XHeapSlice[T]) Len() int           { return len(h.V) }
func (h XHeapSlice[T]) Less(i, j int) bool { return h.Comp(h.V[i], h.V[j]) }
func (h XHeapSlice[T]) Swap(i, j int)      { h.V[i], h.V[j] = h.V[j], h.V[i] }

func (h *XHeapSlice[T]) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.V = append(h.V, x.(T))
}

func (h *XHeapSlice[T]) Pop() any {
	x := h.V[len(h.V)-1]
	h.V = h.V[:len(h.V)-1]
	return x
}

type XHeap[T any] struct {
	h *XHeapSlice[T]
}

func (h XHeap[T]) Compare(a, b T) bool {
	return h.h.Comp(a, b)
}

func (h *XHeap[T]) Push(value T) error {
	heap.Push(h.h, value)
	return nil
}

func (h *XHeap[T]) PushTest(value T) error {
	h.h.Push(value)
	return nil
}

func (h *XHeap[T]) PushMany(values ...T) error {
	for _, x := range values {
		err := h.Push(x)

		if err != nil {
			return err
		}
	}

	return nil
}

func (h XHeap[T]) Front() (T, error) {
	if len(h.h.V) == 0 {
		return generics.ZeroOf[T](), errors.New(fmt.Sprintf("Get front of queue of length %v failed", len(h.h.V)))
	}

	return h.h.V[0], nil
}

func (h *XHeap[T]) Pop() error {
	heap.Remove(h.h, 0)
	return nil
}

func (h *XHeap[T]) PopTest() error {
	//heap.Remove(h.h, 0)
	h.h.Pop()
	return nil
}

func (h XHeap[T]) Empty() bool {
	if len(h.h.V) != 0 {
		return false
	}

	return true
}

func (h XHeap[T]) Size() int {
	return len(h.h.V)
}

func (h XHeap[T]) GetSlice() xslice.Slice[T] {
	return h.h.V
}
