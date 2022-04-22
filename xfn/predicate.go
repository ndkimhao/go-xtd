package xfn

type Predicate[T any] func(T) bool

func (p Predicate[T]) F() func(T) bool {
	return p
}

func (p Predicate[T]) Or(fns ...Predicate[T]) Predicate[T] {
	return Or(append([]Predicate[T]{p}, fns...)...)
}

func (p Predicate[T]) And(fns ...Predicate[T]) Predicate[T] {
	return And(append([]Predicate[T]{p}, fns...)...)
}

func (p Predicate[T]) Negate() Predicate[T] {
	return func(v T) bool { return !p(v) }
}

func Equal[T comparable](rhs T) Predicate[T] {
	return func(lhs T) bool {
		return lhs == rhs
	}
}

func EqualAny[T comparable](rhs ...T) Predicate[T] {
	return func(v T) bool {
		for _, r := range rhs {
			if v == r {
				return true
			}
		}
		return false
	}
}

func NotEqual[T comparable](rhs T) Predicate[T] {
	return func(lhs T) bool {
		return lhs != rhs
	}
}

func NotEqualAll[T comparable](rhs ...T) Predicate[T] {
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
