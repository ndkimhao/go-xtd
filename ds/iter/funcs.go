package iter

func Distance[T any, It ConstRandomIterator[T, It]](a, b It) int {
	_ = a.Equal(b)
	return b.Pos() - a.Pos()
}
