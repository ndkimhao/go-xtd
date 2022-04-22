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

// TODO: handle integer overflow/underflow
type intIter[T constraints.Integer] struct {
	start T
	end   T
	step  T
	cur   T
	inf   bool
}

func (i *intIter[T]) Next() (value T, ok bool) {
	if !i.inf && ((i.step > 0 && i.cur > i.end) || (i.step < 0 && i.cur < i.end)) {
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
		inf:   false,
	})
}

func RangeN[T constraints.Integer](n T) *Stream[T] {
	return Range[T](0, n-1, 1)
}

func IntegerSequence[T constraints.Integer](start, step T) *Stream[T] {
	if step == 0 {
		panic("stream.IntegerSequence: step is 0")
	}
	return New[T](&intIter[T]{
		start: start,
		step:  step,
		cur:   start,
		inf:   true,
	})
}

func (g Generator[T]) Next() (value T, ok bool) {
	return g(), true
}

func (g Generator[T]) SkipNext(n int) (skipped int) {
	for i := 0; i < n; i++ {
		g()
	}
	return n
}

func (g BoundedGenerator[T]) Next() (value T, ok bool) {
	return g()
}

func (g BoundedGenerator[T]) SkipNext(n int) (skipped int) {
	for i := 0; i < n; i++ {
		_, ok := g()
		if !ok {
			return i
		}
	}
	return n
}

func Generate[T any](generator Generator[T]) *Stream[T] {
	return New[T](generator)
}

func BoundedGenerate[T any](generator BoundedGenerator[T]) *Stream[T] {
	return New[T](generator)
}
