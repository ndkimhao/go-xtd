package ring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRing(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		r := New[int]()
		assert.Panics(t, func() { r.DeleteLast() })
		assert.Panics(t, func() { r.DeleteFirst() })
		assert.Panics(t, func() { r.First() })
		assert.Panics(t, func() { r.Last() })
		assert.Panics(t, func() { r.At(0) })
		assert.Equal(t, []int(nil), r.ToSlice(nil))
		r.Append(1)
		assert.Equal(t, []int{1}, r.ToSlice(nil))
		r.Append(1)
		r.Append(2)
		assert.Equal(t, 2, r.Last())
		assert.Equal(t, 1, r.First())
		r.Append(3)
		r.Append(4)
		assert.Equal(t, []int{1, 1, 2, 3, 4}, r.ToSlice(nil))
		r.Prepend(5)
		assert.Equal(t, 5, r.First())
		assert.Equal(t, []int{5, 1, 1, 2, 3, 4}, r.ToSlice(nil))
	})
}
