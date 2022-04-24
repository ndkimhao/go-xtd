package slice

import (
	"reflect"
	"unsafe"
)

func UnorderedDelete[T any](s []T, i int) []T {
	last := len(s) - 1
	if i < last {
		s[i], s[last] = s[last], s[i]
	}
	return s[:last]
}

func Delete[T any](s []T, i int) []T {
	if i < len(s)-1 {
		copy(s[i:], s[i+1:])
	}
	return s[:len(s)-1]
}

func DeleteLast[T any](s []T) []T {
	return s[:len(s)-1]
}

func ReferenceEqual[T any](a []T, b []T) bool {
	ha := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	hb := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return ha.Data == hb.Data
}
