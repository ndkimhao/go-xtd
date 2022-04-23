package xfn

import (
	"github.com/ndkimhao/go-xtd/constraints"
)

func Plus[T constraints.Ordered](rhs T) UnaryOperator[T] {
	return func(lhs T) T {
		return lhs + rhs
	}
}

func Minus[T constraints.Number](rhs T) UnaryOperator[T] {
	return func(lhs T) T {
		return lhs - rhs
	}
}

func Mult[T constraints.Number](rhs T) UnaryOperator[T] {
	return func(lhs T) T {
		return lhs * rhs
	}
}

func Div[T constraints.Number](rhs T) UnaryOperator[T] {
	return func(lhs T) T {
		return lhs / rhs
	}
}

func Mod[T constraints.Integer](rhs T) UnaryOperator[T] {
	return func(lhs T) T {
		return lhs % rhs
	}
}

func PlusOp[T constraints.Ordered](lhs, rhs T) T {
	return lhs + rhs
}

func MultOp[T constraints.Number](lhs, rhs T) T {
	return lhs * rhs
}

func DivOp[T constraints.Number](lhs, rhs T) T {
	return lhs / rhs

}

func ModOp[T constraints.Integer](lhs, rhs T) T {
	return lhs % rhs
}
