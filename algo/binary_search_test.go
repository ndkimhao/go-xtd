package algo_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ndkimhao/go-xtd/algo"
	"github.com/ndkimhao/go-xtd/ds/xslice"
)

func TestUpperBound(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		//            0  1  2  3  4  5  6
		s := xslice.Of(1, 1, 2, 2, 3, 5, 7)
		check := func(pos, val int) {
			require.Equal(t, s.Begin().Add(pos), algo.UpperBound(s.Range(), val))
		}
		check(0, 0)
		check(2, 1)
		check(4, 2)
		check(5, 3)
		check(5, 4)
		check(6, 5)
		check(6, 6)
		check(7, 7)
		check(7, 8)
	})
	t.Run("Reversed", func(t *testing.T) {
		//            0  1  2  3  4  5  6
		s := xslice.Of(7, 5, 3, 2, 2, 1, 1)
		check := func(pos, val int) {
			require.Equal(t, s.Begin().Add(pos+1), algo.UpperBound(s.ReverseRange(), val).Base())
		}
		check(6, 0)
		check(4, 1)
		check(2, 2)
		check(1, 3)
		check(1, 4)
		check(0, 5)
		check(0, 6)
		check(-1, 7)
		check(-1, 8)
	})
	t.Run("Empty", func(t *testing.T) {
		s := xslice.Of[int]()
		require.Equal(t, s.End(), algo.UpperBound(s.Range(), 1))
	})
	t.Run("Empty Reversed", func(t *testing.T) {
		s := xslice.Of[int]()
		require.Equal(t, s.REnd(), algo.UpperBound(s.ReverseRange(), 1))
	})
}

func TestLowerBound(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		//            0  1  2  3  4  5  6
		s := xslice.Of(1, 1, 2, 2, 3, 5, 7)
		check := func(pos, val int) {
			require.Equal(t, s.Begin().Add(pos), algo.LowerBound(s.Range(), val))
		}
		check(0, 0)
		check(0, 1)
		check(2, 2)
		check(4, 3)
		check(5, 4)
		check(5, 5)
		check(6, 6)
		check(6, 7)
		check(7, 8)
	})
	t.Run("Reversed", func(t *testing.T) {
		//            0  1  2  3  4  5  6
		s := xslice.Of(7, 5, 3, 2, 2, 1, 1)
		check := func(pos, val int) {
			require.Equal(t, s.Begin().Add(pos+1), algo.LowerBound(s.ReverseRange(), val).Base())
		}
		check(6, 0)
		check(6, 1)
		check(4, 2)
		check(2, 3)
		check(1, 4)
		check(1, 5)
		check(0, 6)
		check(0, 7)
		check(-1, 8)
	})
	t.Run("Empty", func(t *testing.T) {
		s := xslice.Of[int]()
		require.Equal(t, s.End(), algo.LowerBound(s.Range(), 1))
	})
	t.Run("Empty Reversed", func(t *testing.T) {
		s := xslice.Of[int]()
		require.Equal(t, s.REnd(), algo.LowerBound(s.ReverseRange(), 1))
	})
}
