package algo

import (
	"github.com/ndkimhao/go-xtd/iter"
)

func FillN[T any, It iter.Iterator[T, It]](first It, count int, value T) It {
	for i := 0; i < count; i++ {
		first.Set(value)
		first = first.Next()
	}
	return first
}
