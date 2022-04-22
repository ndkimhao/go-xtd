package stream

import (
	"github.com/ndkimhao/go-xtd/constraints"
)

type sliceIter[T any] struct {
	a []T
	i int
}

func (s *sliceIter[T]) Next() (value T, ok bool) {
	if s.i == len(s.a) {
		var zero T
		return zero, false
	}
	x := s.i
	s.i++
	return s.a[x], true
}

func (s *sliceIter[T]) SkipNext(n int) (skipped int) {
	if n < 0 {
		panic("sliceIter.SkipNext: negative value")
	}
	sz := len(s.a)
	if s.i+n >= sz { // TODO: handle overflow
		rem := sz - s.i
		s.i = sz
		return rem
	}
	s.i += n
	return n
}

func From[T any](slice []T) *Stream[T] {
	return New[T](&sliceIter[T]{a: slice})
}

func Of[T any](values ...T) *Stream[T] {
	return From(values)
}

// TODO: handle integer overflow/underflow
type intIter[T constraints.Integer] struct {
	start T
	end   T
	step  T
	cur   T
}

func (i *intIter[T]) Next() (value T, ok bool) {
	if (i.step > 0 && i.cur > i.end) || (i.step < 0 && i.cur < i.end) {
		return 0, false
	}
	x := i.cur
	i.cur += i.step
	return x, true
}

func (i *intIter[T]) SkipNext(n int) (skipped int) {
	i.cur += T(n) * i.step // TODO: handle overflow/underflow
	return -1
}

func Range[T constraints.Integer](start, end, step T) *Stream[T] {
	if step == 0 {
		panic("stream.Range: step is 0")
	}
	return New[T](&intIter[T]{
		start: start,
		end:   end,
		step:  step,
		cur:   start,
	})
}

func RangeN[T constraints.Integer](n T) *Stream[T] {
	return Range[T](0, n-1, 1)
}
