package stream

import (
	"github.com/ndkimhao/go-xtd/xtd"
)

type filter[T any] struct {
	_ xtd.NoCopy

	it   xtd.Iterator[T]
	pred Predicate[T]

	has  bool
	last T
}

func newFilter[T any](it xtd.Iterator[T], pred Predicate[T]) *filter[T] {
	return &filter[T]{it: it, pred: pred}
}

func (f *filter[T]) fetch() {
	if f.has {
		return
	}
	for it := f.it; it.Has(); it = it.Next() {
		if v := it.Get(); f.pred(v) {
			f.has, f.last = true, v
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

func (f *filter[T]) Skip(n int) (iterator xtd.Iterator[T], skipped int) {
	f.fetch() // lazy fetch
	x := 0
	it := f.it
	for ; x < n && it.Has(); it = it.Next() {
		if v := it.Get(); f.pred(v) {
			x++
			if x == n {
				f.has, f.last = true, v
				return f, x
			}
		}
	}
	var zero T
	f.has, f.last = false, zero
	return it, x
}
