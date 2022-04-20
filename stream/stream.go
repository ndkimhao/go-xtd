package stream

import (
	"github.com/ndkimhao/gstl/xtd"
)

type Predicate[T any] func(value T) bool
type Mapper[T any, R any] func(old T) (new R)

type Stream[T any] struct {
	_ xtd.NoCopy

	it xtd.Iterator[T]
}

func NewStream[T any](source xtd.Iterator[T]) *Stream[T] {
	return &Stream[T]{it: source}
}

func (s *Stream[T]) Map(mapper Mapper[T, T]) *Stream[T] {
	panic("implement me")
}

func (s *Stream[T]) Filter(predicate Predicate[T]) *Stream[T] {
	return NewStream[T](newFilter(s.it, predicate))
}

func Map[R any, T any](s *Stream[T], mapper Mapper[T, R]) *Stream[R] {
	panic("implement me")
}
func (s *Stream[T]) All(predicate Predicate[T]) bool {
	for it := s.it; it.Has(); it = it.Next() {
		if !predicate(it.Get()) {
			return false
		}
	}
	return true
}

func (s *Stream[T]) Any(predicate Predicate[T]) bool {
	for it := s.it; it.Has(); it = it.Next() {
		if predicate(it.Get()) {
			return true
		}
	}
	return false
}
