package algo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/algo"
	"github.com/ndkimhao/go-xtd/ds/xslice"
)

func TestReverse(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		s := xslice.Of(3, 2, 1)
		algo.Reverse[int](s.Begin(), s.End())
		assert.Equal(t, xslice.Of(1, 2, 3), s)
	})
	t.Run("Reversed", func(t *testing.T) {
		s := xslice.Of(3, 2, 1)
		algo.Reverse[int](s.RBegin(), s.REnd())
		assert.Equal(t, xslice.Of(1, 2, 3), s)
	})
	t.Run("Empty", func(t *testing.T) {
		s := xslice.Of[int]()
		algo.Reverse[int](s.Begin(), s.End())
		assert.Equal(t, xslice.Of[int](), s)
	})
	t.Run("Empty Reversed", func(t *testing.T) {
		s := xslice.Of[int]()
		algo.Reverse[int](s.RBegin(), s.REnd())
		assert.Equal(t, xslice.Of[int](), s)
	})
}

func TestSwap(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		s := xslice.Of(1, 2, 3, 4, 5, 6)
		algo.Swap[int](s.Begin().Add(1), s.Begin().Add(3))
		assert.Equal(t, xslice.Of(1, 4, 3, 2, 5, 6), s)
	})
	t.Run("Reversed", func(t *testing.T) {
		s := xslice.Of(1, 2, 3, 4, 5, 6)
		algo.Swap[int](s.RBegin().Add(1), s.RBegin().Add(3))
		assert.Equal(t, xslice.Of(1, 2, 5, 4, 3, 6), s)
	})
	t.Run("Same", func(t *testing.T) {
		s := xslice.Of(1, 2)
		algo.Swap[int](s.Begin(), s.Begin())
		assert.Equal(t, xslice.Of(1, 2), s)
	})
}
