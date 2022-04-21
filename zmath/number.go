package zmath

import (
	"github.com/ndkimhao/go-xtd/constraints"
)

func Min[T constraints.Ordered](a, b T) T {
	if a <= b {
		return a
	} else {
		return b
	}
}

func Max[T constraints.Ordered](a, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}
