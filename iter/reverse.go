package iter

// Const Iterator

type ReverseConst[T any, It ConstBidirIterator[T, It]] struct {
	i It
}

func ReverseConstIterator[T any, It ConstBidirIterator[T, It]](i It) ReverseConst[T, It] {
	return ReverseConst[T, It]{i: i}
}

func (c ReverseConst[T, It]) Base() It {
	return c.i
}

func (c ReverseConst[T, It]) Next() ReverseConst[T, It] {
	return ReverseConst[T, It]{i: c.i.Prev()}
}

func (c ReverseConst[T, It]) Prev() ReverseConst[T, It] {
	return ReverseConst[T, It]{i: c.i.Next()}
}

func (c ReverseConst[T, It]) Value() T {
	return c.i.Prev().Value()
}

func (c ReverseConst[T, It]) Equal(other ReverseConst[T, It]) bool {
	return c.i.Equal(other.i)
}

// Iterator

type Reverse[T any, It BidirIterator[T, It]] struct {
	i It
}

func ReverseIterator[T any, It BidirIterator[T, It]](i It) Reverse[T, It] {
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

type ReverseConstRandom[T any, It ConstRandomIterator[T, It]] struct {
	i It
}

func ReverseConstRandomIterator[T any, It ConstRandomIterator[T, It]](i It) ReverseConstRandom[T, It] {
	return ReverseConstRandom[T, It]{i: i}
}

func (c ReverseConstRandom[T, It]) Base() It {
	return c.i
}

func (c ReverseConstRandom[T, It]) Next() ReverseConstRandom[T, It] {
	return ReverseConstRandom[T, It]{i: c.i.Prev()}
}

func (c ReverseConstRandom[T, It]) Prev() ReverseConstRandom[T, It] {
	return ReverseConstRandom[T, It]{i: c.i.Next()}
}

func (c ReverseConstRandom[T, It]) Add(offset int) ReverseConstRandom[T, It] {
	return ReverseConstRandom[T, It]{i: c.i.Add(-offset)}
}

func (c ReverseConstRandom[T, It]) Position() int {
	return -c.i.Position()
}

func (c ReverseConstRandom[T, It]) Value() T {
	return c.i.Prev().Value()
}

func (c ReverseConstRandom[T, It]) Equal(other ReverseConstRandom[T, It]) bool {
	return c.i.Equal(other.i)
}

// Random

type ReverseRandom[T any, It RandomIterator[T, It]] struct {
	i It
}

func ReverseRandomIterator[T any, It RandomIterator[T, It]](i It) ReverseRandom[T, It] {
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
