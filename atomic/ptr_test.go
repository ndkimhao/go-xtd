package atomic

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func runTestPtr[T any](t *testing.T, typeName string) {
	a, b, c := new(T), new(T), new(T)

	tests := []struct {
		name   string
		newPtr func() *Ptr[T]
		init   *T
	}{
		{
			name:   "Normal",
			newPtr: func() *Ptr[T] { return NewPtr[T](a) },
			init:   a,
		},
		{
			name:   "Nil",
			newPtr: func() *Ptr[T] { return NewPtr[T](nil) },
			init:   nil,
		},
		{
			name:   "Zero",
			newPtr: func() *Ptr[T] { return new(Ptr[T]) },
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
				assert.False(t, p.CAS(c, b))
				assert.Equal(t, test.init, p.Load())
			})
			t.Run("CAS_ok", func(t *testing.T) {
				p := test.newPtr()
				assert.True(t, p.CAS(test.init, b))
				assert.Equal(t, b, p.Load())
			})
		})
	}
}

func TestPtr(t *testing.T) {
	runTestPtr[int](t, "int")
	runTestPtr[time.Time](t, "time.Time")
}
