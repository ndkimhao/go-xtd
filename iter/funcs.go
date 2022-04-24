package iter

func Distance[T interface{ Position() int }](a, b T) int {
	return b.Position() - a.Position()
}
