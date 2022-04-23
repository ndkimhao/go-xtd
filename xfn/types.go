package xfn

type Predicate[T any] func(T) bool
type BiPredicate[T, U any] func(T, U) bool

type UnaryOperator[T any] func(T) T
type BinaryOperator[T any] func(T, T) T

type Function[T, R any] func(T) R
type BiFunction[T, U, R any] func(T, U) R
