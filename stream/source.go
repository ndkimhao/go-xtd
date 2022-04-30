package stream

import (
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/ds/iter"
	"github.com/ndkimhao/go-xtd/generics"
	"github.com/ndkimhao/go-xtd/xmath"
)

type sliceSource[T any] struct {
	a []T
	i int
}

func (s *sliceSource[T]) Next() (value T, ok bool) {
	if s.i == len(s.a) {
		var zero T
		return zero, false
	}
	x := s.i
	s.i++
	return s.a[x], true
}

func (s *sliceSource[T]) SkipNext(n int) (skipped int) {
	if n < 0 {
		panic("sliceSource.SkipNext: negative value")
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
type intSource[T constraints.Integer] struct {
	start T
	end   T
	step  T
	cur   T
	inf   bool
}

func (i *intSource[T]) Next() (value T, ok bool) {
	if !i.inf && ((i.step > 0 && i.cur > i.end) || (i.step < 0 && i.cur < i.end)) {
		return 0, false
	}
	x := i.cur
	i.cur += i.step
	return x, true
}

func (i *intSource[T]) SkipNext(n int) (skipped int) {
	i.cur += T(n) * i.step // TODO: handle overflow/underflow
	return -1
}

func Range[T constraints.Integer](start, end, step T) *Stream[T] {
	if step == 0 {
		panic("stream.Range: step is 0")
	}
	return New[T](&intSource[T]{
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
	return New[T](&intSource[T]{
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

type iteratorSource[T any, It iter.ConstIterator[T, It]] struct {
	cur It
	end It
}

func OfRange[T any, It iter.ConstIterator[T, It]](r iter.Range[T, It]) *Stream[T] {
	_ = r.Begin.Equal(r.End)
	return New[T](&iteratorSource[T, It]{cur: r.Begin, end: r.End})
}

func OfIterators[T any, It iter.ConstIterator[T, It]](first, last It) *Stream[T] {
	_ = first.Equal(last)
	return New[T](&iteratorSource[T, It]{cur: first, end: last})
}

func (s *iteratorSource[T, It]) Next() (value T, ok bool) {
	if s.cur.Equal(s.end) {
		return generics.ZeroOf[T](), false
	}
	v := s.cur.Get()
	s.cur = s.cur.Next()
	return v, true
}

func (s *iteratorSource[T, It]) SkipNext(n int) (skipped int) {
	if rCur, ok := any(s.cur).(iter.ConstRandomIterator[T, It]); ok {
		rEnd := any(s.end).(iter.ConstRandomIterator[T, It])
		skipped = xmath.Min(n, rEnd.Pos()-rCur.Pos())
		s.cur = rCur.Add(skipped)
		return
	}

	skipped = 0
	for ; skipped < n; skipped++ {
		if _, ok := s.Next(); !ok {
			break
		}
	}
	return
}
