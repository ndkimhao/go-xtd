package atomic_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/atomic"
)

func TestNewPtr(t *testing.T) {
	i := 1
	p := atomic.NewPtr(&i)
	assert.Equal(t, &i, p.Load())
}

func runTestPtr[T any](t *testing.T, typeName string) {
	a, b, c := new(T), new(T), new(T)

	tests := []struct {
		name   string
		newPtr func() *atomic.Ptr[T]
		init   *T
	}{
		{
			name:   "Normal",
			newPtr: func() *atomic.Ptr[T] { return atomic.NewPtr[T](a) },
			init:   a,
		},
		{
			name:   "Nil",
			newPtr: func() *atomic.Ptr[T] { return atomic.NewPtr[T](nil) },
			init:   nil,
		},
		{
			name:   "Zero",
			newPtr: func() *atomic.Ptr[T] { return new(atomic.Ptr[T]) },
			init:   nil,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(typeName, "_", test.name), func(t *testing.T) {
			t.Run("Load", func(t *testing.T) {
				p := test.newPtr()
				assert.Equal(t, test.init, p.Load())
			})
			t.Run("Store", func(t *testing.T) {
				p := test.newPtr()
				p.Store(b)
				assert.Equal(t, b, p.Load())
			})
			t.Run("Swap", func(t *testing.T) {
				p := test.newPtr()
				assert.Equal(t, test.init, p.Swap(b)) // Swap returns old value
				assert.Equal(t, b, p.Load())          // p has new value now
			})
			t.Run("CAS_fail", func(t *testing.T) {
				p := test.newPtr()
				assert.False(t, p.CompareAndSwap(c, b))
				assert.Equal(t, test.init, p.Load())
			})
			t.Run("CAS_ok", func(t *testing.T) {
				p := test.newPtr()
				assert.True(t, p.CompareAndSwap(test.init, b))
				assert.Equal(t, b, p.Load())
			})
		})
	}
}

func TestPtr(t *testing.T) {
	runTestPtr[int](t, "int")
	runTestPtr[time.Time](t, "time.Time")
}
