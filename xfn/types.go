package xfn

type Predicate[T any] func(T) bool
type UnaryOperator[T any] func(T) T
type BinaryOperator[T any] func(T, T) T
type BiFunction[T, U, R any] func(T, U) R
