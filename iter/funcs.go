package iter

func Distance[T interface {
	Pos() int
	Equal(T) bool
}](a, b T) int {
	_ = a.Equal(b)
	return b.Pos() - a.Pos()
}
