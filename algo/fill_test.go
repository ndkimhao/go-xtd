package algo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/algo"
	"github.com/ndkimhao/go-xtd/ds/xring"
	"github.com/ndkimhao/go-xtd/iter"
	"github.com/ndkimhao/go-xtd/slice"
)

func TestFillN(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		s := slice.Of(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		it := algo.FillN(s.Begin().Add(1), 4, -1)
		assert.True(t, it.Equal(s.Begin().Add(5)))
		assert.Equal(t, slice.Of(0, -1, -1, -1, -1, 5, 6, 7, 8, 9), s)
	})
	t.Run("Append", func(t *testing.T) {
		s := slice.Of(0, 2)
		algo.FillN(iter.Append[int](&s), 3, -1)
		assert.Equal(t, slice.Of(0, 2, -1, -1, -1), s)
	})
	t.Run("Prepend", func(t *testing.T) {
		s := xring.Of(0, 2)
		algo.FillN(iter.Prepend[int](&s), 3, -1)
		assert.Equal(t, slice.Of(-1, -1, -1, 0, 2).Slice(), s.ToSlice(nil))
	})
	t.Run("Insert", func(t *testing.T) {
		s := slice.Of(0, 1, 2, 3, 4)
		algo.FillN(iter.Insert[int, slice.Iterator[int]](s.IteratorAt(2), &s), 3, -1)
		assert.Equal(t, slice.Of(0, 1, -1, -1, -1, 2, 3, 4), s)
	})
}
