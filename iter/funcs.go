package iter

func Distance[T interface {
	Position() int
	Equal(T) bool
}](a, b T) int {
	_ = a.Equal(b)
	return b.Position() - a.Position()
}
