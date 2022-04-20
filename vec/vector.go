package vec

import (
	"fmt"
)

type Vector[T any] []T

func Of[T any](values ...T) Vector[T] {
	return values
}

func OfSlice[T any](values []T) Vector[T] {
	return values
}

func (v *Vector[T]) PushBack(value T) {
	*v = append(*v, value)
}

func (v *Vector[T]) PushBackMany(values ...T) {
	*v = append(*v, values...)
}

func (v *Vector[T]) PopBack() (removedValue T) {
	last := len(*v) - 1
	val := (*v)[last]
	*v = (*v)[:last]
	return val
}

func (v Vector[T]) Size() int {
	return len(v)
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

func (v Vector[T]) Front() T {
	return v[0]
}

func (v Vector[T]) Back() T {
	return v[len(v)-1]
}
