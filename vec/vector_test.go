package vec_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/vec"
)

func TestOf(t *testing.T) {
	assert.Equal(t, []int(nil), vec.Of[int]().Slice())
	assert.Equal(t, []int{1, 2, 3}, vec.Of(1, 2, 3).Slice())
	assert.Equal(t, []string{"hello", "world"}, vec.Of("hello", "world").Slice())

	assert.Equal(t, []int{1, 2, 3}, vec.OfSlice([]int{1, 2, 3}).Slice())
}

func TestVector_Zero(t *testing.T) {
	var v vec.Vector[int]
	assert.Equal(t, []int(nil), v.Slice())
	assert.Equal(t, 0, v.Len())
}

func TestVector_ForRange(t *testing.T) {
	for i, x := range vec.Of(1, 2, 3) {
		assert.Equal(t, i+1, x)
	}
}

func TestVector_Len(t *testing.T) {
	assert.Equal(t, 0, vec.Of[int]().Len())
	assert.Equal(t, 3, vec.Of(1, 2, 3).Len())
}

func TestVector_Append(t *testing.T) {
	var v vec.Vector[int]
	x := []int{}
	for i := 1; i < 100; i++ {
		v.Append(i)
		x = append(x, i)
		assert.Equal(t, x, v.Slice())
	}
}

func TestVector_AppendMany(t *testing.T) {
	var v vec.Vector[int]
	x := []int{}
	for i := 1; i < 100; i++ {
		v.AppendMany(i, i*2)
		x = append(x, i, i*2)
		assert.Equal(t, x, v.Slice())
	}
}

func TestVector_EraseEnd(t *testing.T) {
	var v vec.Vector[int]
	for i := 1; i < 100; i++ {
		v.Append(i)
	}
	for i := 99; i >= 1; i-- {
		assert.Equal(t, i, v.Last())
		v.DeleteLast()
	}
}

func TestVector_Slice(t *testing.T) {
	assert.Equal(t, []string(nil), vec.Of[string]().Slice())
	assert.Equal(t, []string{"a"}, vec.Of("a").Slice())
}

func TestVector_Index(t *testing.T) {
	v := vec.Of(1, 2, 3)
	assert.Equal(t, 1, v[0])
	assert.Equal(t, 2, v[1])
	assert.Equal(t, 3, v[2])
	assert.Panics(t, func() {
		_ = v[3]
	})
}

func TestVector_At(t *testing.T) {
	v := vec.Of(1, 2, 3)
	assert.Equal(t, 1, v.At(0))
	assert.Equal(t, 2, v.At(1))
	assert.Equal(t, 3, v.At(2))
	assert.PanicsWithValue(t, "index out of bound: n=-1 len=3", func() {
		_ = v.At(-1)
	})
	assert.PanicsWithValue(t, "index out of bound: n=3 len=3", func() {
		_ = v.At(3)
	})
}

func TestVector_Front(t *testing.T) {
	assert.Equal(t, 1, vec.Of(1, 2, 3).First())
	assert.Panics(t, func() {
		vec.Of[int]().First()
	})
}

func TestVector_Back(t *testing.T) {
	assert.Equal(t, 3, vec.Of(1, 2, 3).Last())
	assert.Panics(t, func() {
		vec.Of[int]().Last()
	})
}
