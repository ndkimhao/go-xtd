package slice

import (
	"reflect"
	"unsafe"
)

func UnorderedDelete[T any](s []T, i int) []T {
	if i < 0 || len(s) <= i {
		panic("delete out of bound")
	}
	last := len(s) - 1
	if i < last {
		s[i], s[last] = s[last], s[i]
	}
	return s[:last]
}

func Delete[T any](s []T, i int) []T {
	if i < 0 || len(s) <= i {
		panic("delete out of bound")
	}
	if i < len(s)-1 {
		copy(s[i:], s[i+1:])
	}
	return s[:len(s)-1]
}

func DeleteLast[T any](s []T) []T {
	if len(s) == 0 {
		panic("delete from empty slice")
	}
	return s[:len(s)-1]
}

func Insert[T any](s []T, i int, value T) []T {
	if i < 0 || len(s) < i {
		panic("insert out of bound")
	}
	if i < len(s) {
		s = append(s, value)
		copy(s[i+1:], s[i:len(s)-1])
		s[i] = value
		return s
	} else {
		return append(s, value)
	}
}

func ReferenceEqual[T any](a []T, b []T) bool {
	ha := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	hb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return ha.Data == hb.Data
}
