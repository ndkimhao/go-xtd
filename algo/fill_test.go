package algo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/ds/iter"

	"github.com/ndkimhao/go-xtd/algo"
	"github.com/ndkimhao/go-xtd/ds/xring"
	xslice2 "github.com/ndkimhao/go-xtd/ds/xslice"
)

func TestFillN(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		s := xslice2.Of(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		it := algo.FillN(s.Begin().Add(1), 4, -1)
		assert.True(t, it.Equal(s.Begin().Add(5)))
		assert.Equal(t, xslice2.Of(0, -1, -1, -1, -1, 5, 6, 7, 8, 9), s)
	})
	t.Run("Append", func(t *testing.T) {
		s := xslice2.Of(0, 2)
		algo.FillN(iter.Append[int](&s), 3, -1)
		assert.Equal(t, xslice2.Of(0, 2, -1, -1, -1), s)
	})
	t.Run("Prepend", func(t *testing.T) {
		s := xring.Of(0, 2)
		algo.FillN(iter.Prepend[int](&s), 3, -1)
		assert.Equal(t, xslice2.Of(-1, -1, -1, 0, 2).Slice(), s.ToSlice(nil))
	})
	t.Run("Insert", func(t *testing.T) {
		s := xslice2.Of(0, 1, 2, 3, 4)
		algo.FillN(iter.Insert[int, xslice2.Iterator[int]](s.IteratorAt(2), &s), 3, -1)
		assert.Equal(t, xslice2.Of(0, 1, -1, -1, -1, 2, 3, 4), s)
	})
}
