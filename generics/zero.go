package generics

func ZeroOf[T any]() T {
	var zero T
	return zero
}
