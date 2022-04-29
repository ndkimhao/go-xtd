package iter

// Const Iterator

type ConstReverseIterator[T any, It ConstBidirIterator[T, It]] struct {
	i It
}

func ConstReverse[T any, It ConstBidirIterator[T, It]](i It) ConstReverseIterator[T, It] {
	return ConstReverseIterator[T, It]{i: i}
}

func (c ConstReverseIterator[T, It]) Base() It {
	return c.i
}

func (c ConstReverseIterator[T, It]) Next() ConstReverseIterator[T, It] {
	return ConstReverseIterator[T, It]{i: c.i.Prev()}
}

func (c ConstReverseIterator[T, It]) Prev() ConstReverseIterator[T, It] {
	return ConstReverseIterator[T, It]{i: c.i.Next()}
}

func (c ConstReverseIterator[T, It]) Value() T {
	return c.i.Prev().Value()
}

func (c ConstReverseIterator[T, It]) Equal(other ConstReverseIterator[T, It]) bool {
	return c.i.Equal(other.i)
}

// Iterator

type ReverseItarator[T any, It BidirIterator[T, It]] struct {
	i It
}

func Reverse[T any, It BidirIterator[T, It]](i It) ReverseItarator[T, It] {
	return ReverseItarator[T, It]{i: i}
}

func (c ReverseItarator[T, It]) Base() It {
	return c.i
}

func (c ReverseItarator[T, It]) Next() ReverseItarator[T, It] {
	return ReverseItarator[T, It]{i: c.i.Prev()}
}

func (c ReverseItarator[T, It]) Prev() ReverseItarator[T, It] {
	return ReverseItarator[T, It]{i: c.i.Next()}
}

func (c ReverseItarator[T, It]) Value() T {
	return c.i.Prev().Value()
}

func (c ReverseItarator[T, It]) SetValue(value T) {
	c.i.Prev().SetValue(value)
}

func (c ReverseItarator[T, It]) Equal(other ReverseItarator[T, It]) bool {
	return c.i.Equal(other.i)
}

// Random

type ConstReverseRandomIterator[T any, It ConstRandomIterator[T, It]] struct {
	i It
}

func ConstReverseRandom[T any, It ConstRandomIterator[T, It]](i It) ConstReverseRandomIterator[T, It] {
	return ConstReverseRandomIterator[T, It]{i: i}
}

func (c ConstReverseRandomIterator[T, It]) Base() It {
	return c.i
}

func (c ConstReverseRandomIterator[T, It]) Next() ConstReverseRandomIterator[T, It] {
	return ConstReverseRandomIterator[T, It]{i: c.i.Prev()}
}

func (c ConstReverseRandomIterator[T, It]) Prev() ConstReverseRandomIterator[T, It] {
	return ConstReverseRandomIterator[T, It]{i: c.i.Next()}
}

func (c ConstReverseRandomIterator[T, It]) Add(offset int) ConstReverseRandomIterator[T, It] {
	return ConstReverseRandomIterator[T, It]{i: c.i.Add(-offset)}
}

func (c ConstReverseRandomIterator[T, It]) Position() int {
	return -c.i.Position()
}

func (c ConstReverseRandomIterator[T, It]) Value() T {
	return c.i.Prev().Value()
}

func (c ConstReverseRandomIterator[T, It]) Equal(other ConstReverseRandomIterator[T, It]) bool {
	return c.i.Equal(other.i)
}

// Random

type ReverseRandomIterator[T any, It RandomIterator[T, It]] struct {
	i It
}

func ReverseRandom[T any, It RandomIterator[T, It]](i It) ReverseRandomIterator[T, It] {
	return ReverseRandomIterator[T, It]{i: i}
}

func (c ReverseRandomIterator[T, It]) Base() It {
	return c.i
}

func (c ReverseRandomIterator[T, It]) Next() ReverseRandomIterator[T, It] {
	return ReverseRandomIterator[T, It]{i: c.i.Prev()}
}

func (c ReverseRandomIterator[T, It]) Prev() ReverseRandomIterator[T, It] {
	return ReverseRandomIterator[T, It]{i: c.i.Next()}
}

func (c ReverseRandomIterator[T, It]) Add(offset int) ReverseRandomIterator[T, It] {
	return ReverseRandomIterator[T, It]{i: c.i.Add(-offset)}
}

func (c ReverseRandomIterator[T, It]) Position() int {
	return -c.i.Position()
}

func (c ReverseRandomIterator[T, It]) Value() T {
	return c.i.Prev().Value()
}

func (c ReverseRandomIterator[T, It]) SetValue(value T) {
	c.i.Prev().SetValue(value)
}

func (c ReverseRandomIterator[T, It]) Equal(other ReverseRandomIterator[T, It]) bool {
	return c.i.Equal(other.i)
}
