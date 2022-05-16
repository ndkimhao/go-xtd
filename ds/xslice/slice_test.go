package xslice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/ds/xslice"
)

func TestOf(t *testing.T) {
	assert.Equal(t, []int(nil), xslice.Of[int]().Slice())
	assert.Equal(t, []int{1, 2, 3}, xslice.Of(1, 2, 3).Slice())
	assert.Equal(t, []string{"hello", "world"}, xslice.Of("hello", "world").Slice())

	assert.Equal(t, []int{1, 2, 3}, xslice.OfSlice([]int{1, 2, 3}).Slice())
}

func TestSlice_Zero(t *testing.T) {
	var v xslice.Slice[int]
	assert.Equal(t, []int(nil), v.Slice())
	assert.Equal(t, 0, v.Len())
}

func TestSlice_ForRange(t *testing.T) {
	for i, x := range xslice.Of(1, 2, 3) {
		assert.Equal(t, i+1, x)
	}
}

func TestSlice_Len(t *testing.T) {
	assert.Equal(t, 0, xslice.Of[int]().Len())
	assert.Equal(t, 3, xslice.Of(1, 2, 3).Len())
}

func TestSlice_Append(t *testing.T) {
	var v xslice.Slice[int]
	x := []int{}
	for i := 1; i < 100; i++ {
		v.Append(i)
		x = append(x, i)
		assert.Equal(t, x, v.Slice())
	}
}

func TestSlice_AppendMany(t *testing.T) {
	var v xslice.Slice[int]
	x := []int{}
	for i := 1; i < 100; i++ {
		v.AppendMany(i, i*2)
		x = append(x, i, i*2)
		assert.Equal(t, x, v.Slice())
	}
}

func TestSlice_EraseEnd(t *testing.T) {
	var v xslice.Slice[int]
	for i := 1; i < 100; i++ {
		v.Append(i)
	}
	for i := 99; i >= 1; i-- {
		assert.Equal(t, i, v.Last())
		v.DeleteLast()
	}
}

func TestSlice_Slice(t *testing.T) {
	assert.Equal(t, []string(nil), xslice.Of[string]().Slice())
	assert.Equal(t, []string{"a"}, xslice.Of("a").Slice())
}

func TestSlice_Index(t *testing.T) {
	v := xslice.Of(1, 2, 3)
	assert.Equal(t, 1, v[0])
	assert.Equal(t, 2, v[1])
	assert.Equal(t, 3, v[2])
	assert.Panics(t, func() {
		_ = v[3]
	})
}

func TestSlice_At(t *testing.T) {
	v := xslice.Of(1, 2, 3)
	assert.Equal(t, 1, v.At(0))
	assert.Equal(t, 2, v.At(1))
	assert.Equal(t, 3, v.At(2))
	assert.Panics(t, func() { v.At(-1) })
	assert.Panics(t, func() { v.At(3) })
}

func TestSlice_Front(t *testing.T) {
	assert.Equal(t, 1, xslice.Of(1, 2, 3).First())
	assert.Panics(t, func() {
		xslice.Of[int]().First()
	})
}

func TestSlice_Back(t *testing.T) {
	assert.Equal(t, 3, xslice.Of(1, 2, 3).Last())
	assert.Panics(t, func() {
		xslice.Of[int]().Last()
	})
}

func TestSlice_Reversed(t *testing.T) {
	assert.Equal(t, xslice.Of[int](), xslice.Of[int]().Reversed())
	assert.Equal(t, xslice.Of(1), xslice.Of(1).Reversed())
	assert.Equal(t, xslice.Of(1, 2, 3), xslice.Of(3, 2, 1).Reversed())
	assert.Equal(t, xslice.Of(1, 2, 3, 4), xslice.Of(4, 3, 2, 1).Reversed())
}

func TestSlice_Insert(t *testing.T) {
	t.Run("From Empty", func(t *testing.T) {
		s := xslice.Of[int]()
		s.InsertAt(0, 5)
		assert.Equal(t, xslice.Of(5), s)
	})
	t.Run("Begin", func(t *testing.T) {
		s := xslice.Of(1, 2, 3)
		s.InsertAt(0, 5)
		assert.Equal(t, xslice.Of(5, 1, 2, 3), s)
	})
	t.Run("Last", func(t *testing.T) {
		s := xslice.Of(1, 2, 3)
		s.InsertAt(3, 5)
		assert.Equal(t, xslice.Of(1, 2, 3, 5), s)
	})
	t.Run("Middle", func(t *testing.T) {
		s := xslice.Of(1, 2, 3, 4, 5, 6)
		s.InsertAt(2, 10)
		assert.Equal(t, xslice.Of(1, 2, 10, 3, 4, 5, 6), s)
	})
	t.Run("Panic", func(t *testing.T) {
		s := xslice.Of(1, 2, 3)
		assert.Panics(t, func() { s.InsertAt(-1, 10) })
		assert.Panics(t, func() { s.InsertAt(4, 10) })
	})
}

func TestInsert(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		s := xslice.Of(1, 2)
		v := xslice.Of(1, 2)
		assert.PanicsWithValue(t, "iterator does not belongs to this slice", func() { s.Insert(v.Begin(), 0) })
	})
	t.Run("Normal", func(t *testing.T) {
		s := xslice.Of(1, 2, 3)
		it := s.Insert(s.Begin().Add(1), 4)
		assert.Equal(t, 4, it.Get())
		assert.True(t, it.Equal(s.Begin().Add(1)))
		assert.Equal(t, xslice.Of(1, 4, 2, 3), s)
	})
}
