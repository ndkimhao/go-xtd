package stream

import (
	"github.com/ndkimhao/gstl/xtd"
)

type filter[T any] struct {
	s    *Stream[T]
	pred Predicate[T]

	has  bool
	last T
}

func newFilter[T any](s *Stream[T], pred Predicate[T]) *filter[T] {
	return &filter[T]{s: s, pred: pred}
}

func (f *filter[T]) fetch() {
	if f.has {
		return
	}
	for it := f.s.src; it.Has(); it = it.Next() {
		if v := it.Get(); f.pred(v) {
			f.has = true
			f.last = v
			return
		}
	}
}

func (f *filter[T]) Get() T {
	f.fetch() // lazy fetch
	if !f.has {
		panic("filter.Get: end of stream")
	}
	return f.last
}

func (f *filter[T]) Has() bool {
	f.fetch() // lazy fetch
	return f.has
}

func (f *filter[T]) Next() xtd.Iterator[T] {
	f.fetch() // lazy fetch
	if !f.has {
		panic("filter.Next: end of stream")
	}
	var zero T
	f.has, f.last = false, zero
	f.fetch()
	return f
}

func (f *filter[T]) Skip(n int) (it xtd.Iterator[T], skipped int) {
	f.fetch() // lazy fetch
	cnt := 0
	for i := 0; i < n; i++ {
		if !f.has {
			break
		}
		f.has = false
		f.fetch()
		cnt++
	}
	return f, cnt
}
