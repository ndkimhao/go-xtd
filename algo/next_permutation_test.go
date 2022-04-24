package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/slice"
)

func TestNextPermutation(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		a := slice.Of(1, 2, 3)
		assert.True(t, NextPermutation[int](a.Begin(), a.End()))
		assert.Equal(t, slice.Of(1, 3, 2), a)
		assert.True(t, NextPermutation[int](a.Begin(), a.End()))
		assert.Equal(t, slice.Of(2, 1, 3), a)
		assert.True(t, NextPermutation[int](a.Begin(), a.End()))
		assert.Equal(t, slice.Of(2, 3, 1), a)
		assert.True(t, NextPermutation[int](a.Begin(), a.End()))
		assert.Equal(t, slice.Of(3, 1, 2), a)
		assert.True(t, NextPermutation[int](a.Begin(), a.End()))
		assert.Equal(t, slice.Of(3, 2, 1), a)
		assert.False(t, NextPermutation[int](a.Begin(), a.End()))
		assert.Equal(t, slice.Of(1, 2, 3), a)
	})
	t.Run("One Element", func(t *testing.T) {
		a := slice.Of(1)
		assert.False(t, NextPermutation[int](a.Begin(), a.End()))
		assert.Equal(t, slice.Of(1), a)
	})
	t.Run("Empty", func(t *testing.T) {
		a := slice.Of[int]()
		assert.False(t, NextPermutation[int](a.Begin(), a.End()))
		assert.Equal(t, slice.Of[int](), a)
	})
}
