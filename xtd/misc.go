package xtd

func PtrOf[T any](value T) *T {
	return &value
}
