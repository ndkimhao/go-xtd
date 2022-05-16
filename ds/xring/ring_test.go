package xring_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ndkimhao/go-xtd/ds/xring"
)

func TestRing(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		r := xring.New[int]()
		require.Panics(t, func() { r.DeleteLast() })
		require.Panics(t, func() { r.DeleteFirst() })
		require.Panics(t, func() { r.First() })
		require.Panics(t, func() { r.Last() })
		require.Panics(t, func() { r.At(0) })
		require.Equal(t, []int(nil), r.ToSlice(nil))
		r.Append(1)
		require.Equal(t, []int{1}, r.ToSlice(nil))
		r.Append(1)
		r.Append(2)
		require.Equal(t, 2, r.Last())
		require.Equal(t, 1, r.First())
		r.Append(3)
		r.Append(4)
		require.Equal(t, []int{1, 1, 2, 3, 4}, r.ToSlice(nil))
		r.Prepend(5)
		require.Equal(t, 5, r.First())
		require.Equal(t, []int{5, 1, 1, 2, 3, 4}, r.ToSlice(nil))
		require.Equal(t, 8, r.Cap())
		require.Panics(t, func() { r.At(-1) })
		require.Panics(t, func() { r.At(6) })
	})
}

func TestRing_Large(t *testing.T) {
	n := 100000
	if testing.Short() {
		n = 100
	}
	r := xring.New[int]()
	for i := 1; i <= n; i++ {
		r.Append(i)
		r.Prepend(-i)
		require.Equal(t, -i, r.At(0))
		require.Equal(t, i, r.At(r.Len()-1))
	}
	require.GreaterOrEqual(t, r.Cap(), n*2)
	for i := n; i >= n/2; i-- {
		require.Equal(t, -i, r.First())
		require.Equal(t, i, r.Last())
		r.DeleteFirst()
		r.DeleteLast()
	}
	for i := 1; i <= n; i++ {
		r.Append(i)
		r.Prepend(-i)
	}
	for i := n; i >= 1; i-- {
		require.Equal(t, -i, r.First())
		require.Equal(t, i, r.Last())
		r.DeleteFirst()
		r.DeleteLast()
	}
	for i := n/2 - 1; i >= 1; i-- {
		require.Equal(t, -i, r.First())
		require.Equal(t, i, r.Last())
		r.DeleteFirst()
		r.DeleteLast()
	}
	require.Equal(t, []int(nil), r.ToSlice(nil))
	require.Less(t, r.Cap(), 1024)
}
