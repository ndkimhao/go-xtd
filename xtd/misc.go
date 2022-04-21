package xtd

func Addr[T any](value T) *T {
	return &value
}
