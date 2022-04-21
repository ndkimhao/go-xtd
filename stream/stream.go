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

type op[T any] interface {
	apply(valuePtr *T) (keep bool)
}

func (p Predicate[T]) apply(valuePtr *T) (keep bool) {
	return p(*valuePtr)
}

func (t Transformer[T]) apply(valuePtr *T) (keep bool) {
	*valuePtr = t(*valuePtr)
	return true
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

type predicateSkip[T any] struct {
	p Predicate[T]
	n int
}

func (ps *predicateSkip[T]) apply(valuePtr *T) (keep bool) {
	if ps.p(*valuePtr) {
		if ps.n <= 0 {
			return true
		} else {
			ps.n--
		}
	}
	return false
}

type Stream[T any] struct {
	_ xtd.NoCopy

	src Iterator[T]
	ops []op[T]

	lastPred int
}

func NewStream[T any](source Iterator[T]) *Stream[T] {
	return &Stream[T]{src: source, lastPred: -1}
}

// Iterator interface

func (s *Stream[T]) Next() (value T, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (s *Stream[T]) SkipNext(n int) (skipped int) {
	s.Skip(n)
	return -1
}

// Immediate operations

func (s *Stream[T]) Map(transformer Transformer[T]) *Stream[T] {
	s.ops = append(s.ops, transformer)
	return s
}

func (s *Stream[T]) Filter(predicate Predicate[T]) *Stream[T] {
	s.lastPred = len(s.ops)
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
	if s.lastPred == -1 {
		s.src.SkipNext(n)
		return s
	}
	p := s.ops[s.lastPred]
	if ps, ok := p.(*predicateSkip[T]); ok {
		ps.n += n
	} else {
		s.ops[s.lastPred] = &predicateSkip[T]{p: p.(Predicate[T]), n: n}
	}
	return s
}

func Map[R any, T any](s *Stream[T], transformer TypeTransformer[T, R]) *Stream[R] {
	return NewStream[R](&typeTransformerStream[T, R]{s: s, t: transformer})
}

// Terminating operations

func (s *Stream[T]) ForEach(consumer Consumer[T]) {
	for {
		v, ok := s.Next()
		if !ok {
			break
		}
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
