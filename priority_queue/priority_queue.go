package priority_queue

import (
	"github.com/ndkimhao/go-xtd/ds/xslice"
)

type PriorityQueue[T any] interface {
	Compare(T, T) bool
	Push(T) error
	PushMany(values ...T) error
	Front() (T, error)
	Pop() error
	Empty() bool
	Size() int
	GetSlice() xslice.Slice[T]
}

func NewXPriorityQueue[T any](comp func(T, T) bool, data ...T) (*XPriorityQueue[T], error) {
	var q XPriorityQueue[T]
	q.comp = comp

	if err := q.PushMany(data...); err != nil {
		return nil, err
	}

	return &q, nil
}

func NewXHeap[T any](comp func(T, T) bool, data ...T) (*XHeap[T], error) {
	var h XHeap[T]
	h.h = &XHeapSlice[T]{comp, nil}
	if err := h.PushMany(data...); err != nil {
		return nil, err
	}

	return &h, nil
}
