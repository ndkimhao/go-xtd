package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/slice"
)

func TestOf(t *testing.T) {
	assert.Equal(t, []int(nil), slice.Of[int]().Slice())
	assert.Equal(t, []int{1, 2, 3}, slice.Of(1, 2, 3).Slice())
	assert.Equal(t, []string{"hello", "world"}, slice.Of("hello", "world").Slice())

	assert.Equal(t, []int{1, 2, 3}, slice.OfSlice([]int{1, 2, 3}).Slice())
}

func TestSlice_Zero(t *testing.T) {
	var v slice.Slice[int]
	assert.Equal(t, []int(nil), v.Slice())
	assert.Equal(t, 0, v.Len())
}

func TestSlice_ForRange(t *testing.T) {
	for i, x := range slice.Of(1, 2, 3) {
		assert.Equal(t, i+1, x)
	}
}

func TestSlice_Len(t *testing.T) {
	assert.Equal(t, 0, slice.Of[int]().Len())
	assert.Equal(t, 3, slice.Of(1, 2, 3).Len())
}

func TestSlice_Append(t *testing.T) {
	var v slice.Slice[int]
	x := []int{}
	for i := 1; i < 100; i++ {
		v.Append(i)
		x = append(x, i)
		assert.Equal(t, x, v.Slice())
	}
}

func TestSlice_AppendMany(t *testing.T) {
	var v slice.Slice[int]
	x := []int{}
	for i := 1; i < 100; i++ {
		v.AppendMany(i, i*2)
		x = append(x, i, i*2)
		assert.Equal(t, x, v.Slice())
	}
}

func TestSlice_EraseEnd(t *testing.T) {
	var v slice.Slice[int]
	for i := 1; i < 100; i++ {
		v.Append(i)
	}
	for i := 99; i >= 1; i-- {
		assert.Equal(t, i, v.Last())
		v.DeleteLast()
	}
}

func TestSlice_Slice(t *testing.T) {
	assert.Equal(t, []string(nil), slice.Of[string]().Slice())
	assert.Equal(t, []string{"a"}, slice.Of("a").Slice())
}

func TestSlice_Index(t *testing.T) {
	v := slice.Of(1, 2, 3)
	assert.Equal(t, 1, v[0])
	assert.Equal(t, 2, v[1])
	assert.Equal(t, 3, v[2])
	assert.Panics(t, func() {
		_ = v[3]
	})
}

func TestSlice_At(t *testing.T) {
	v := slice.Of(1, 2, 3)
	assert.Equal(t, 1, v.At(0))
	assert.Equal(t, 2, v.At(1))
	assert.Equal(t, 3, v.At(2))
	assert.Panics(t, func() { v.At(-1) })
	assert.Panics(t, func() { v.At(3) })
}

func TestSlice_Front(t *testing.T) {
	assert.Equal(t, 1, slice.Of(1, 2, 3).First())
	assert.Panics(t, func() {
		slice.Of[int]().First()
	})
}

func TestSlice_Back(t *testing.T) {
	assert.Equal(t, 3, slice.Of(1, 2, 3).Last())
	assert.Panics(t, func() {
		slice.Of[int]().Last()
	})
}

func TestSlice_Reversed(t *testing.T) {
	assert.Equal(t, slice.Of[int](), slice.Of[int]().Reversed())
	assert.Equal(t, slice.Of(1), slice.Of(1).Reversed())
	assert.Equal(t, slice.Of(1, 2, 3), slice.Of(3, 2, 1).Reversed())
	assert.Equal(t, slice.Of(1, 2, 3, 4), slice.Of(4, 3, 2, 1).Reversed())
}

func TestSlice_Insert(t *testing.T) {
	t.Run("From Empty", func(t *testing.T) {
		s := slice.Of[int]()
		s.Insert(0, 5)
		assert.Equal(t, slice.Of(5), s)
	})
	t.Run("Begin", func(t *testing.T) {
		s := slice.Of(1, 2, 3)
		s.Insert(0, 5)
		assert.Equal(t, slice.Of(5, 1, 2, 3), s)
	})
	t.Run("Last", func(t *testing.T) {
		s := slice.Of(1, 2, 3)
		s.Insert(3, 5)
		assert.Equal(t, slice.Of(1, 2, 3, 5), s)
	})
	t.Run("Middle", func(t *testing.T) {
		s := slice.Of(1, 2, 3, 4, 5, 6)
		s.Insert(2, 10)
		assert.Equal(t, slice.Of(1, 2, 10, 3, 4, 5, 6), s)
	})
	t.Run("Panic", func(t *testing.T) {
		s := slice.Of(1, 2, 3)
		assert.Panics(t, func() { s.Insert(-1, 10) })
		assert.Panics(t, func() { s.Insert(4, 10) })
	})
}
