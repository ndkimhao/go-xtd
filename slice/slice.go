package slice

import (
	"fmt"
)

type Slice[T any] []T

func New[T any]() Slice[T] {
	return nil
}

func Of[T any](values ...T) Slice[T] {
	return values
}

func OfSlice[T any](values []T) Slice[T] {
	return values
}

func (v *Slice[T]) Append(value T) {
	*v = append(*v, value)
}

func (v *Slice[T]) AppendMany(values ...T) {
	*v = append(*v, values...)
}

func (v *Slice[T]) Delete(i int) {
	s := *v
	if i < len(s)-1 {
		copy(s[i:], s[i+1:])
	}
	*v = s[:len(s)-1]
}

func (v *Slice[T]) DeleteLast() {
	s := *v
	*v = s[:len(s)-1]
}

func UnorderedDelete[T any](s []T, i int) []T {
	last := len(s) - 1
	if i < last {
		s[i], s[last] = s[last], s[i]
	}
	return s[:last]
}

func (v *Slice[T]) UnorderedDelete(i int) {
	*v = UnorderedDelete(*v, i)
}

func (v Slice[T]) Len() int {
	return len(v)
}

func (v Slice[T]) Cap() int {
	return cap(v)
}

func (v Slice[T]) Sub(start, end int) Slice[T] {
	return v[start:end]
}

func (v Slice[T]) Slice() []T {
	return v
}

func (v Slice[T]) At(n int) T {
	if n < 0 || n >= len(v) {
		panic(fmt.Sprint("index out of bound: n=", n, " len=", len(v)))
	}
	return v[n]
}

func (v Slice[T]) First() T {
	return v[0]
}

func (v Slice[T]) Last() T {
	return v[len(v)-1]
}
