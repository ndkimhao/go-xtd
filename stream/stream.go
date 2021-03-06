package stream

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ndkimhao/go-xtd/generics"
	"github.com/ndkimhao/go-xtd/xfn"
	"github.com/ndkimhao/go-xtd/xtd"
)

type Consumer[T any] func(value T)
type Generator[T any] func() (value T)
type BoundedGenerator[T any] func() (value T, ok bool)

type Source[T any] interface {
	// Next returns item from stream. When reaching end-of-stream,
	// returns `value` is zero value of T and `ok` is false.
	Next() (value T, ok bool)

	// SkipNext skips next `n` items. Returns number of skipped items,
	// or returns -1 if the underlying stream evaluates lazily.
	SkipNext(n int) (skipped int)
}

type typeTransformerStream[T any, R any] struct {
	s *Stream[T]
	t xfn.Function[T, R]
}

func (tts *typeTransformerStream[T, R]) Next() (value R, ok bool) {
	old, ok := tts.s.Next()
	if !ok {
		var zero R
		return zero, false
	}
	return tts.t(old), true
}

func (tts *typeTransformerStream[T, R]) SkipNext(n int) (skipped int) {
	return tts.s.SkipNext(n)
}

// use int to avoid allocation when storing in `ops []any`
type predSkip int
type predLimit int

type Stream[T any] struct {
	_ xtd.NoCopy

	src Source[T]
	ops []any
	buf [8]any

	hasPred bool
}

func New[T any](source Source[T]) *Stream[T] {
	s := &Stream[T]{src: source}
	s.ops = s.buf[:0] // small slice optimization
	return s
}

// Source interface

func (s *Stream[T]) Next() (value T, ok bool) {
	if s.empty() {
		goto end_of_stream
	}
loop_src:
	for v, hasNext := s.src.Next(); hasNext; v, hasNext = s.src.Next() {
		for i, oAny := range s.ops {
			switch op := oAny.(type) {
			case xfn.Predicate[T]:
				if !op(v) {
					continue loop_src // Skip item because xfn.Predicate[T] returns false
				}
			case predSkip:
				if op > 0 {
					s.ops[i] = op - 1
					continue loop_src // Skip item
				}
			case predLimit:
				if op > 0 {
					s.ops[i] = op - 1 // Under limit, take item
				} else {
					goto end_of_stream // Run out of limit, stop
				}
			case xfn.UnaryOperator[T]:
				v = op(v)
			case Consumer[T]:
				op(v)
			default:
				panic("Stream.Next: internal error: invalid op")
			}
		}
		return v, true
	}
end_of_stream:
	s.clear()
	return generics.ZeroOf[T](), false
}

func (s *Stream[T]) SkipNext(n int) (skipped int) {
	s.Skip(n)
	return -1
}

// Intermediate operations

func (s *Stream[T]) Map(transformer xfn.UnaryOperator[T]) *Stream[T] {
	s.ops = append(s.ops, transformer)
	return s
}

func (s *Stream[T]) Filter(predicate xfn.Predicate[T]) *Stream[T] {
	s.hasPred = true
	s.ops = append(s.ops, predicate)
	return s
}

func (s *Stream[T]) Skip(n int) *Stream[T] {
	if n < 0 {
		panic("Stream.Skip: negative value")
	}
	if n == 0 {
		return s
	}
	if !s.hasPred {
		s.src.SkipNext(n)
		return s
	}
	s.hasPred = true
	s.ops = append(s.ops, predSkip(n))
	return s
}

func (s *Stream[T]) Limit(n int) *Stream[T] {
	if n < 0 {
		panic("Stream.Limit: negative value")
	}
	s.hasPred = true
	s.ops = append(s.ops, predLimit(n))
	return s
}

func (s *Stream[T]) Peek(consumer Consumer[T]) *Stream[T] {
	s.ops = append(s.ops, consumer)
	return s
}

func Map[R any, T any](s *Stream[T], transformer xfn.Function[T, R]) *Stream[R] {
	return New[R](&typeTransformerStream[T, R]{s: s, t: transformer})
}

func (s *Stream[T]) clear() {
	s.src = nil
	s.ops = nil
}

func (s *Stream[T]) empty() bool {
	return s.src == nil
}

func (s *Stream[T]) MapToInt(transformer xfn.Function[T, int]) *Stream[int] {
	return Map(s, transformer)
}

func (s *Stream[T]) MapToString(transformer xfn.Function[T, string]) *Stream[string] {
	return Map(s, transformer)
}

func (s *Stream[T]) MapToBytes(transformer xfn.Function[T, []byte]) *Stream[[]byte] {
	return Map(s, transformer)
}

func (s *Stream[T]) MapToFloat(transformer xfn.Function[T, float64]) *Stream[float64] {
	return Map(s, transformer)
}

func (s *Stream[T]) MapToBool(transformer xfn.Function[T, bool]) *Stream[bool] {
	return Map(s, transformer)
}

func (s *Stream[T]) MapToAny(transformer xfn.Function[T, any]) *Stream[any] {
	return Map(s, transformer)
}

// Terminating operations

func (s *Stream[T]) Collect(consumer Consumer[T]) {
	for v, ok := s.Next(); ok; v, ok = s.Next() {
		consumer(v)
	}
}

func (s *Stream[T]) All(predicate xfn.Predicate[T]) bool {
	for v, ok := s.Next(); ok; v, ok = s.Next() {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func (s *Stream[T]) Any(predicate xfn.Predicate[T]) bool {
	for v, ok := s.Next(); ok; v, ok = s.Next() {
		if predicate(v) {
			return true
		}
	}
	return false
}

func (s *Stream[T]) None(predicate xfn.Predicate[T]) bool {
	return !s.Any(predicate)
}

func (s *Stream[T]) Count() int {
	count := 0
	for _, ok := s.Next(); ok; _, ok = s.Next() {
		count++
	}
	return count
}

func (s *Stream[T]) First() T {
	v, ok := s.Next()
	if !ok {
		panic("Stream.First: stream is empty")
	}
	return v
}

func (s *Stream[T]) Slice() []T {
	var r []T
	for v, ok := s.Next(); ok; v, ok = s.Next() {
		r = append(r, v)
	}
	return r
}

func (s *Stream[T]) Reduce(identity T, accumulator xfn.BinaryOperator[T]) T {
	r := identity
	for v, ok := s.Next(); ok; v, ok = s.Next() {
		r = accumulator(r, v)
	}
	return r
}

func (s *Stream[T]) String() string {
	var sb strings.Builder
	sb.WriteString("Stream[")
	sb.WriteString(reflect.TypeOf((*T)(nil)).Elem().Name())
	sb.WriteString("]{")
	first := true
	s.Collect(func(value T) {
		if first {
			first = false
		} else {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprint(value))
	})
	sb.WriteString("}")
	return sb.String()
}

func Reduce[A, T any](s *Stream[T], identity A, accumulator xfn.BiFunction[A, T, A]) A {
	r := identity
	for v, ok := s.Next(); ok; v, ok = s.Next() {
		r = accumulator(r, v)
	}
	return r
}
