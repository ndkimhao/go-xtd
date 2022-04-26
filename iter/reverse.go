package iter

// Const Iterator

type ConstReverse[T any, It ConstBidirIterator[T, It]] struct {
	i It
}

func MakeConstReverse[T any, It ConstBidirIterator[T, It]](i It) ConstReverse[T, It] {
	return ConstReverse[T, It]{i: i}
}

func (c ConstReverse[T, It]) Base() It {
	return c.i
}

func (c ConstReverse[T, It]) Next() ConstReverse[T, It] {
	return ConstReverse[T, It]{i: c.i.Prev()}
}

func (c ConstReverse[T, It]) Prev() ConstReverse[T, It] {
	return ConstReverse[T, It]{i: c.i.Next()}
}

func (c ConstReverse[T, It]) Value() T {
	return c.i.Prev().Value()
}

func (c ConstReverse[T, It]) Equal(other ConstReverse[T, It]) bool {
	return c.i.Equal(other.i)
}

// Iterator

type Reverse[T any, It BidirIterator[T, It]] struct {
	i It
}

func MakeReverse[T any, It BidirIterator[T, It]](i It) Reverse[T, It] {
	return Reverse[T, It]{i: i}
}

func (c Reverse[T, It]) Base() It {
	return c.i
}

func (c Reverse[T, It]) Next() Reverse[T, It] {
	return Reverse[T, It]{i: c.i.Prev()}
}

func (c Reverse[T, It]) Prev() Reverse[T, It] {
	return Reverse[T, It]{i: c.i.Next()}
}

func (c Reverse[T, It]) Value() T {
	return c.i.Prev().Value()
}

func (c Reverse[T, It]) SetValue(value T) {
	c.i.Prev().SetValue(value)
}

func (c Reverse[T, It]) Equal(other Reverse[T, It]) bool {
	return c.i.Equal(other.i)
}

// Random

type ConstReverseRandom[T any, It ConstRandomIterator[T, It]] struct {
	i It
}

func MakeConstReverseRandom[T any, It ConstRandomIterator[T, It]](i It) ConstReverseRandom[T, It] {
	return ConstReverseRandom[T, It]{i: i}
}

func (c ConstReverseRandom[T, It]) Base() It {
	return c.i
}

func (c ConstReverseRandom[T, It]) Next() ConstReverseRandom[T, It] {
	return ConstReverseRandom[T, It]{i: c.i.Prev()}
}

func (c ConstReverseRandom[T, It]) Prev() ConstReverseRandom[T, It] {
	return ConstReverseRandom[T, It]{i: c.i.Next()}
}

func (c ConstReverseRandom[T, It]) Add(offset int) ConstReverseRandom[T, It] {
	return ConstReverseRandom[T, It]{i: c.i.Add(-offset)}
}

func (c ConstReverseRandom[T, It]) Position() int {
	return -c.i.Position()
}

func (c ConstReverseRandom[T, It]) Value() T {
	return c.i.Prev().Value()
}

func (c ConstReverseRandom[T, It]) Equal(other ConstReverseRandom[T, It]) bool {
	return c.i.Equal(other.i)
}

// Random

type ReverseRandom[T any, It RandomIterator[T, It]] struct {
	i It
}

func MakeReverseRandom[T any, It RandomIterator[T, It]](i It) ReverseRandom[T, It] {
	return ReverseRandom[T, It]{i: i}
}

func (c ReverseRandom[T, It]) Base() It {
	return c.i
}

func (c ReverseRandom[T, It]) Next() ReverseRandom[T, It] {
	return ReverseRandom[T, It]{i: c.i.Prev()}
}

func (c ReverseRandom[T, It]) Prev() ReverseRandom[T, It] {
	return ReverseRandom[T, It]{i: c.i.Next()}
}

func (c ReverseRandom[T, It]) Add(offset int) ReverseRandom[T, It] {
	return ReverseRandom[T, It]{i: c.i.Add(-offset)}
}

func (c ReverseRandom[T, It]) Position() int {
	return -c.i.Position()
}

func (c ReverseRandom[T, It]) Value() T {
	return c.i.Prev().Value()
}

func (c ReverseRandom[T, It]) SetValue(value T) {
	c.i.Prev().SetValue(value)
}

func (c ReverseRandom[T, It]) Equal(other ReverseRandom[T, It]) bool {
	return c.i.Equal(other.i)
}
