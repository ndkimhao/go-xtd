package algo_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ndkimhao/go-xtd/algo"
	"github.com/ndkimhao/go-xtd/slice"
)

func TestUpperBound(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		//            0  1  2  3  4  5  6
		s := slice.Of(1, 1, 2, 2, 3, 5, 7)
		check := func(pos, val int) {
			require.Equal(t, s.Begin().Add(pos), algo.UpperBound(s.Begin(), s.End(), val))
		}
		check(0, 0)
		check(2, 1)
		check(4, 2)

		check(5, 3)
		check(5, 4)

		check(6, 5)
		check(6, 6)

		check(7, 7)
		check(7, 10)
	})
	t.Run("Empty", func(t *testing.T) {
		s := slice.Of[int]()
		require.Equal(t, s.End(), algo.UpperBound(s.Begin(), s.End(), 1))
	})
}
