package atomic

import (
	"github.com/ndkimhao/gstl/xtd"
)

// Value wrapped a copiable value atomically
// WARNING: Store, Swap, and CompareAndSwap operations do allocate
type Value[T comparable] struct {
	_ xtd.NoCopy

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
	p := v.p.Load()
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// Store atomically stores the passed value.
func (v *Value[T]) Store(val T) {
	// Taking address of function argument causes Go to allocate and copy it
	v.p.Store(&val)
}

// Swap atomically swaps the wrapped pointer and returns the old value.
func (v *Value[T]) Swap(val T) (old T) {
	p := v.p.Swap(&val)
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// CompareAndSwap executes the compare-and-swap operation for the wrapped pointer.
func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	p := v.p.Load()
	if p == nil {
		var zero T
		if zero != old {
			return false
		}
	} else if *p != old {
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
