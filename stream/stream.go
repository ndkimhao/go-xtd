package stream

import (
	"github.com/ndkimhao/go-xtd/xtd"
)

type Predicate[T any] func(value T) bool
type Transformer[T any] func(old T) (new T)
type TypeTransformer[T any, R any] func(old T) (new R)
type Consumer[T any] func(value T)

type Iterator[T any] interface {
	// Next returns item from stream. When reaching end-of-stream,
	// returns `value` is zero value of T and `ok` is false.
	Next() (value T, ok bool)

	// SkipNext skips next `n` items. Returns number of skipped items,
	// or returns -1 if the underlying stream evaluates lazily.
	SkipNext(n int) (skipped int)
}

type typeTransformerStream[T any, R any] struct {
	s *Stream[T]
	t TypeTransformer[T, R]
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

type predSkip[T any] struct {
	skip int
}

type predLimit[T any] struct {
	limit int
}

func (ps *predSkip[T]) keep() bool {
	if ps.skip > 0 {
		ps.skip--
		return false
	}
	return true
}

func (pl *predLimit[T]) keep() bool {
	if pl.limit > 0 {
		pl.limit--
		return true
	}
	return false
}

type Stream[T any] struct {
	_ xtd.NoCopy

	src Iterator[T]
	ops []any

	hasPred bool
}

func NewStream[T any](source Iterator[T]) *Stream[T] {
	return &Stream[T]{src: source}
}

// Iterator interface

func (s *Stream[T]) Next() (value T, ok bool) {
	if s.empty() {
		goto end_of_stream
	}
loop_src:
	for v, hasNext := s.src.Next(); hasNext; v, hasNext = s.src.Next() {
		for _, oAny := range s.ops {
			switch op := oAny.(type) {
			case Predicate[T]:
				if !op(v) {
					continue loop_src
				}
			case *predSkip[T]:
				if !op.keep() {
					continue loop_src
				}
			case *predLimit[T]:
				if !op.keep() {
					continue loop_src
				}
			case Transformer[T]:
				v = op(v)
			default:
				panic("Stream.Next: internal error: invalid op")
			}
		}
		return v, true
	}
end_of_stream:
	s.clear()
	var zero T
	return zero, false
}

func (s *Stream[T]) SkipNext(n int) (skipped int) {
	s.Skip(n)
	return -1
}

// Immediate operations

func (s *Stream[T]) Map(transformer Transformer[T]) *Stream[T] {
	if s.empty() {
		return s
	}
	s.ops = append(s.ops, transformer)
	return s
}

func (s *Stream[T]) Filter(predicate Predicate[T]) *Stream[T] {
	if s.empty() {
		return s
	}
	s.hasPred = true
	s.ops = append(s.ops, predicate)
	return s
}

func (s *Stream[T]) Skip(n int) *Stream[T] {
	if n < 0 {
		panic("Stream.Skip: negative value")
	}
	if s.empty() || n == 0 {
		return s
	}
	if !s.hasPred {
		s.src.SkipNext(n)
		return s
	}
	s.hasPred = true
	s.ops = append(s.ops, &predSkip[T]{skip: n})
	return s
}

func (s *Stream[T]) Limit(n int) *Stream[T] {
	if n < 0 {
		panic("Stream.Limit: negative value")
	}
	if s.empty() {
		return s
	}
	s.hasPred = true
	s.ops = append(s.ops, &predLimit[T]{limit: n})
	return s
}

func Map[R any, T any](s *Stream[T], transformer TypeTransformer[T, R]) *Stream[R] {
	return NewStream[R](&typeTransformerStream[T, R]{s: s, t: transformer})
}

func (s *Stream[T]) clear() {
	s.src = nil
	s.ops = nil
}

func (s *Stream[T]) empty() bool {
	return s.src == nil
}

// Terminating operations

func (s *Stream[T]) ForEach(consumer Consumer[T]) {
	for v, ok := s.Next(); ok; v, ok = s.Next() {
		consumer(v)
	}
}

func (s *Stream[T]) All(predicate Predicate[T]) bool {
	for v, ok := s.Next(); ok; v, ok = s.Next() {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func (s *Stream[T]) Any(predicate Predicate[T]) bool {
	for v, ok := s.Next(); ok; v, ok = s.Next() {
		if predicate(v) {
			return true
		}
	}
	return false
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
