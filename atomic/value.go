package atomic

import (
	"github.com/ndkimhao/gstl/misc"
)

// Value is simpler than Go atomic.Value, however,
// it does not support CompareAndSwap operation
type Value[T comparable] struct {
	_ misc.NoCmpCopy

	p Ptr[T]
}

// NewValue creates a new Value[T].
func NewValue[T comparable](val T) *Value[T] {
	v := &Value[T]{}
	v.Store(val)
	return v
}

// Load atomically loads the wrapped value.
func (v *Value[T]) Load() T {
	return *v.p.Load()
}

// Store atomically stores the passed value.
func (v *Value[T]) Store(val T) {
	// Taking address of function argument causes Go to allocate and copy it
	v.p.Store(&val)
}

// Swap atomically swaps the wrapped pointer and returns the old value.
func (v *Value[T]) Swap(val T) (old T) {
	return *v.p.Swap(&val)
}

// CompareAndSwap executes the compare-and-swap operation for the wrapped pointer.
func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	p := v.p.Load()
	if *p != old {
		return false
	}
	if new == old {
		return true
	}
	// We don't need a CAS loop here
	// v.p is changed only if the underlying value is changed,
	// so if v.p.CompareAndSwap fails, then the value was changed
	return v.p.CompareAndSwap(p, &new)
}
