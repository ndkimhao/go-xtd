package algo_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ndkimhao/go-xtd/algo"
	"github.com/ndkimhao/go-xtd/slice"
	"github.com/ndkimhao/go-xtd/xfn"
)

func TestNextPermutation(t *testing.T) {
	t.Run("Normal 3", func(t *testing.T) {
		a := slice.Of(1, 2, 3)
		require.True(t, algo.NextPermutation(a.Range()))
		require.Equal(t, slice.Of(1, 3, 2), a)
		require.True(t, algo.NextPermutation(a.Range()))
		require.Equal(t, slice.Of(2, 1, 3), a)
		require.True(t, algo.NextPermutation(a.Range()))
		require.Equal(t, slice.Of(2, 3, 1), a)
		require.True(t, algo.NextPermutation(a.Range()))
		require.Equal(t, slice.Of(3, 1, 2), a)
		require.True(t, algo.NextPermutation(a.Range()))
		require.Equal(t, slice.Of(3, 2, 1), a)
		require.False(t, algo.NextPermutation(a.Range()))
		require.Equal(t, slice.Of(1, 2, 3), a)

		require.True(t, algo.NextPermutation(a.SubRange(1, 3)))
		require.Equal(t, slice.Of(1, 3, 2), a)
		require.False(t, algo.NextPermutation(a.SubRange(1, 3)))
		require.Equal(t, slice.Of(1, 2, 3), a)
	})
	t.Run("Normal 2", func(t *testing.T) {
		a := slice.Of(1, 2)
		require.True(t, algo.NextPermutation(a.Range()))
		require.Equal(t, slice.Of(2, 1), a)
		require.False(t, algo.NextPermutation(a.Range()))
		require.Equal(t, slice.Of(1, 2), a)
	})
	t.Run("Normal 3 Reversed", func(t *testing.T) {
		a := slice.Of(1, 2, 3).Reversed()
		require.True(t, algo.NextPermutation(a.ReverseRange()))
		require.Equal(t, slice.Of(1, 3, 2).Reversed(), a)
		require.True(t, algo.NextPermutation(a.ReverseRange()))
		require.Equal(t, slice.Of(2, 1, 3).Reversed(), a)
		require.True(t, algo.NextPermutation(a.ReverseRange()))
		require.Equal(t, slice.Of(2, 3, 1).Reversed(), a)
		require.True(t, algo.NextPermutation(a.ReverseRange()))
		require.Equal(t, slice.Of(3, 1, 2).Reversed(), a)
		require.True(t, algo.NextPermutation(a.ReverseRange()))
		require.Equal(t, slice.Of(3, 2, 1).Reversed(), a)
		require.False(t, algo.NextPermutation(a.ReverseRange()))
		require.Equal(t, slice.Of(1, 2, 3).Reversed(), a)

		require.True(t, algo.NextPermutation(a.ReverseSubRange(1, 3)))
		require.Equal(t, slice.Of(1, 3, 2).Reversed(), a)
		require.False(t, algo.NextPermutation(a.ReverseSubRange(1, 3)))
		require.Equal(t, slice.Of(1, 2, 3).Reversed(), a)
	})
	t.Run("Normal 2 Reversed", func(t *testing.T) {
		a := slice.Of(1, 2).Reversed()
		require.True(t, algo.NextPermutation(a.ReverseRange()))
		require.Equal(t, slice.Of(2, 1).Reversed(), a)
		require.False(t, algo.NextPermutation(a.ReverseRange()))
		require.Equal(t, slice.Of(1, 2).Reversed(), a)
	})
	t.Run("One Element", func(t *testing.T) {
		a := slice.Of(1)
		require.False(t, algo.NextPermutation(a.Range()))
		require.Equal(t, slice.Of(1), a)
	})
	t.Run("One Element Reversed", func(t *testing.T) {
		a := slice.Of(1)
		require.False(t, algo.NextPermutation(a.ReverseRange()))
		require.Equal(t, slice.Of(1), a)
	})
	t.Run("Empty", func(t *testing.T) {
		a := slice.Of[int]()
		require.False(t, algo.NextPermutation(a.ReverseRange()))
		require.Equal(t, slice.Of[int](), a)
	})
	t.Run("Empty Reversed", func(t *testing.T) {
		a := slice.Of[int]()
		require.False(t, algo.NextPermutation(a.ReverseRange()))
		require.Equal(t, slice.Of[int](), a)
	})
	t.Run("Normal 2 Range", func(t *testing.T) {
		a := slice.Of(1, 2)
		require.True(t, algo.NextPermutationAny[int](a.Range(), xfn.Less[int]))
		require.Equal(t, slice.Of(2, 1), a)
		require.False(t, algo.NextPermutationAny[int](a.Range(), xfn.Less[int]))
		require.Equal(t, slice.Of(1, 2), a)
	})
}
