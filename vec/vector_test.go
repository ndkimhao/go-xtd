package vec_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/gstl/vec"
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
	assert.Equal(t, 0, v.Size())
}

func TestVector_ForRange(t *testing.T) {
	for i, x := range vec.Of(1, 2, 3) {
		assert.Equal(t, i+1, x)
	}
}

func TestVector_Size(t *testing.T) {
	assert.Equal(t, 0, vec.Of[int]().Size())
	assert.Equal(t, 3, vec.Of(1, 2, 3).Size())
}

func TestVector_PushBack(t *testing.T) {
	var v vec.Vector[int]
	x := []int{}
	for i := 1; i < 100; i++ {
		v.PushBack(i)
		x = append(x, i)
		assert.Equal(t, x, v.Slice())
	}
}

func TestVector_PushBackMany(t *testing.T) {
	var v vec.Vector[int]
	x := []int{}
	for i := 1; i < 100; i++ {
		v.PushBackMany(i, i*2)
		x = append(x, i, i*2)
		assert.Equal(t, x, v.Slice())
	}
}

func TestVector_PopBack(t *testing.T) {
	var v vec.Vector[int]
	for i := 1; i < 100; i++ {
		v.PushBack(i)
	}
	for i := 99; i >= 1; i-- {
		assert.Equal(t, i, v.PopBack())
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
	assert.Equal(t, 1, vec.Of(1, 2, 3).Front())
	assert.Panics(t, func() {
		vec.Of[int]().Front()
	})
}

func TestVector_Back(t *testing.T) {
	assert.Equal(t, 3, vec.Of(1, 2, 3).Back())
	assert.Panics(t, func() {
		vec.Of[int]().Back()
	})
}
