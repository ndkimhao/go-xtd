package iter

type Range[T any, It ConstIterator[T, It]] struct {
	Begin It
	End   It
}

func MakeRange[T any, It ConstIterator[T, It]](begin, end It) Range[T, It] {
	return Range[T, It]{Begin: begin, End: end}
}

func ReverseRandomRange[T any, It RandomIterator[T, It]](r Range[T, It]) Range[T, ReverseRandomIterator[T, It]] {
	return MakeRange[T](ReverseRandom[T](r.Begin), ReverseRandom[T](r.End))
}

func SubRange[T any, It ConstRandomIterator[T, It]](r Range[T, It], first, last int) Range[T, It] {
	return MakeRange[T](r.Begin.Add(first), r.Begin.Add(last))
}

func RangeDistance[T any, It ConstRandomIterator[T, It]](r Range[T, It]) int {
	return r.End.Position() - r.Begin.Position()
}
