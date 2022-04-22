package vec

import (
	"fmt"
)

type Vector[T any] []T

func New[T any]() Vector[T] {
	return nil
}

func Of[T any](values ...T) Vector[T] {
	return values
}

func OfSlice[T any](values []T) Vector[T] {
	return values
}

func (v *Vector[T]) Append(value T) {
	*v = append(*v, value)
}

func (v *Vector[T]) AppendMany(values ...T) {
	*v = append(*v, values...)
}

func (v *Vector[T]) Delete(i int) {
	s := *v
	if i < len(s)-1 {
		copy(s[i:], s[i+1:])
	}
	*v = s[:len(s)-1]
}

func (v *Vector[T]) DeleteLast() {
	s := *v
	*v = s[:len(s)-1]
}

func (v *Vector[T]) UnorderedDelete(i int) {
	s := *v
	last := len(s) - 1
	if i < last {
		s[i], s[last] = s[last], s[i]
	}
	*v = s[:len(s)-1]
}

func (v Vector[T]) Len() int {
	return len(v)
}

func (v Vector[T]) Cap() int {
	return cap(v)
}

func (v Vector[T]) Sub(start, end int) Vector[T] {
	return v[start:end]
}

func (v Vector[T]) Slice() []T {
	return v
}

func (v Vector[T]) At(n int) T {
	if n < 0 || n >= len(v) {
		panic(fmt.Sprint("index out of bound: n=", n, " len=", len(v)))
	}
	return v[n]
}

func (v Vector[T]) First() T {
	return v[0]
}

func (v Vector[T]) Last() T {
	return v[len(v)-1]
}
