package atomic

import (
	"sync/atomic"
	"unsafe"

	"github.com/ndkimhao/gstl/xtd"
)

type Ptr[T any] struct {
	_ xtd.NoCopy

	v unsafe.Pointer
}

// NewPtr creates a new Ptr[T].
func NewPtr[T any](val *T) *Ptr[T] {
	return &Ptr[T]{v: unsafe.Pointer(val)}
}

// Load atomically loads the wrapped pointer.
func (p *Ptr[T]) Load() *T {
	return (*T)(atomic.LoadPointer(&p.v))
}

// Store atomically stores the passed pointer.
func (p *Ptr[T]) Store(val *T) {
	atomic.StorePointer(&p.v, unsafe.Pointer(val))
}

// Swap atomically swaps the wrapped pointer and returns the old value.
func (p *Ptr[T]) Swap(val *T) (old *T) {
	return (*T)(atomic.SwapPointer(&p.v, unsafe.Pointer(val)))
}

// CompareAndSwap executes the compare-and-swap operation for the wrapped pointer.
func (p *Ptr[T]) CompareAndSwap(old, new *T) (swapped bool) {
	return atomic.CompareAndSwapPointer(&p.v, unsafe.Pointer(old), unsafe.Pointer(new))
}
