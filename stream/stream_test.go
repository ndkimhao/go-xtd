package stream_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ndkimhao/go-xtd/slice"
	"github.com/ndkimhao/go-xtd/stream"
)

type MockIterator[T any] struct {
	mock.Mock
}

func (m *MockIterator[T]) Next() (value T, ok bool) {
	args := m.Called()
	return args.Get(0).(T), args.Bool(1)
}

func (m *MockIterator[T]) SkipNext(n int) (skipped int) {
	return m.Called(n).Int(0)
}

func TestNew(t *testing.T) {
	assert.Equal(t, []int(nil), stream.New[int](nil).Slice())
	assert.Equal(t, []int{1, 2, 3}, stream.OfSlice([]int{1, 2, 3}).Slice())
}

func TestOf(t *testing.T) {
	assert.Equal(t, []int(nil), stream.Of[int]().Slice())
	assert.Equal(t, []int{1, 2, 3}, stream.Of(1, 2, 3).Slice())
}

func TestStream_Source(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		it := &MockIterator[int]{}
		it.On("Next").Return(1, true).Once()
		it.On("Next").Return(2, true).Once()
		it.On("Next").Return(0, false).Once()
		assert.Equal(t, []int{1, 2}, stream.New[int](it).Slice())
		it.AssertExpectations(t)
	})
	t.Run("Empty", func(t *testing.T) {
		it := &MockIterator[int]{}
		it.On("Next").Return(0, false).Once()
		assert.Equal(t, []int(nil), stream.New[int](it).Slice())
		it.AssertExpectations(t)
	})
	t.Run("Skip", func(t *testing.T) {
		it := &MockIterator[int]{}
		it.On("SkipNext", 5).Return(-1).Once()
		it.On("Next").Return(1, true).Once()
		it.On("Next").Return(0, false).Once()
		assert.Equal(t, []int{1}, stream.New[int](it).Skip(5).Slice())
		it.AssertExpectations(t)
	})
	t.Run("NoSkip", func(t *testing.T) {
		it := &MockIterator[int]{}
		it.On("Next").Return(1, true).Once()
		it.On("Next").Return(2, true).Once()
		it.On("Next").Return(3, true).Once()
		it.On("Next").Return(0, false).Once()
		c1, c2 := 0, 0
		r := stream.New[int](it).
			Filter(func(int) bool {
				c1++
				return true
			}).
			Skip(1).
			Filter(func(int) bool {
				c2++
				return true
			}).
			Slice()
		assert.Equal(t, []int{2, 3}, r)
		assert.Equal(t, 3, c1)
		assert.Equal(t, 2, c2)
		it.AssertExpectations(t)
	})
}

func TestStream_Map(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		r := stream.
			Range(0, 9, 3).
			Map(func(x int) int { return x * x }).
			Slice()
		assert.Equal(t, []int{0, 9, 36, 81}, r)
	})
	t.Run("Complex_1", func(t *testing.T) {
		r := stream.
			Range(0, 50, 3).
			Map(func(x int) int { return x * x }).
			Filter(func(x int) bool { return x%2 == 1 }).
			Map(func(x int) int { return x - 1 }).
			Map(func(x int) int { return x / 2 }).
			Skip(1).
			Limit(2).
			Slice()
		assert.Equal(t, []int{40, 112}, r)
	})
}

func TestStream_Peek(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		a := slice.New[int]()
		r := slice.New[int]()
		stream.
			Range(0, 9, 3).
			Peek(a.Append).
			Map(func(x int) int { return x * x }).
			Peek(r.Append).
			Collect(stream.VoidConsumer[int])
		assert.Equal(t, []int{0, 3, 6, 9}, a.Slice())
		assert.Equal(t, []int{0, 9, 36, 81}, r.Slice())
	})
}

func TestStream_String(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		assert.Equal(t, "Stream[int]{1, 2, 3}", stream.Of(1, 2, 3).String())
	})
}

func TestMap(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		r := stream.Map(stream.Range(0, 9, 3), strconv.Itoa).Slice()
		assert.Equal(t, []string{"0", "3", "6", "9"}, r)
	})
}
