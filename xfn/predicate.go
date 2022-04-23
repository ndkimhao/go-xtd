package xfn

import (
	"bytes"

	"github.com/ndkimhao/go-xtd/constraints"
)

func (p Predicate[T]) F() func(T) bool {
	return p
}

func (p Predicate[T]) Or(fns ...Predicate[T]) Predicate[T] {
	return Or(append([]Predicate[T]{p}, fns...)...)
}

func (p Predicate[T]) And(fns ...Predicate[T]) Predicate[T] {
	return And(append([]Predicate[T]{p}, fns...)...)
}

func (p Predicate[T]) Neg() Predicate[T] {
	return func(v T) bool { return !p(v) }
}

func Eq[T comparable](rhs T) Predicate[T] {
	return func(lhs T) bool {
		return lhs == rhs
	}
}

func EqAny[T comparable](rhs ...T) Predicate[T] {
	return func(v T) bool {
		for _, r := range rhs {
			if v == r {
				return true
			}
		}
		return false
	}
}

func Neq[T comparable](rhs T) Predicate[T] {
	return func(lhs T) bool {
		return lhs != rhs
	}
}

func NeqAny[T comparable](rhs ...T) Predicate[T] {
	return func(v T) bool {
		for _, r := range rhs {
			if v == r {
				return false
			}
		}
		return true
	}
}

func True[T any]() Predicate[T] {
	return func(T) bool { return true }
}

func False[T any]() Predicate[T] {
	return func(T) bool { return false }
}

func Or[T any](fns ...Predicate[T]) Predicate[T] {
	if len(fns) == 0 {

	}
	return func(v T) bool {
		for _, f := range fns {
			if f(v) {
				return true
			}
		}
		return false
	}
}

func And[T any](fns ...Predicate[T]) Predicate[T] {
	return func(v T) bool {
		for _, f := range fns {
			if f(v) {
				return true
			}
		}
		return false
	}
}

func Not[T any](fns ...Predicate[T]) Predicate[T] {
	return func(v T) bool {
		for _, f := range fns {
			if f(v) {
				return false
			}
		}
		return true
	}
}

func Greater[T constraints.Ordered](rhs T) Predicate[T] {
	return func(v T) bool {
		return v > rhs
	}
}

func GreaterEq[T constraints.Ordered](rhs T) Predicate[T] {
	return func(v T) bool {
		return v >= rhs
	}
}

func Less[T constraints.Ordered](rhs T) Predicate[T] {
	return func(v T) bool {
		return v < rhs
	}
}

func LessEq[T constraints.Ordered](rhs T) Predicate[T] {
	return func(v T) bool {
		return v <= rhs
	}
}

func Prefix(prefix string) Predicate[string] {
	return func(v string) bool {
		i := len(prefix)
		return len(v) >= i && v[:i] == prefix
	}
}

func Suffix(suffix string) Predicate[string] {
	return func(v string) bool {
		i := len(suffix)
		return len(v) >= i && v[len(v)-i:] == suffix
	}
}

func PrefixBytes(prefix []byte) Predicate[[]byte] {
	return func(v []byte) bool {
		i := len(prefix)
		return len(v) >= i && bytes.Equal(v[:i], prefix)
	}
}

func SuffixBytes(suffix []byte) Predicate[[]byte] {
	return func(v []byte) bool {
		i := len(suffix)
		return len(v) >= i && bytes.Equal(v[len(v)-i:], suffix)
	}
}

func EqOp[T comparable](lhs, rhs T) bool {
	return lhs == rhs
}

func NeqEqOp[T comparable](lhs, rhs T) bool {
	return lhs != rhs
}

func GreaterOp[T constraints.Ordered](lhs, rhs T) bool {
	return lhs > rhs
}

func GreaterEqOp[T constraints.Ordered](lhs, rhs T) bool {
	return lhs >= rhs
}

func LessOp[T constraints.Ordered](lhs, rhs T) bool {
	return lhs < rhs
}

func LessEqOp[T constraints.Ordered](lhs, rhs T) bool {
	return lhs <= rhs
}
