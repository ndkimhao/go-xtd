package iter

// ConstReverseIterator implements ConstIterator
var _ ConstIterator[int, ConstReverseIterator[int, DummyIterator[int]]] = ConstReverseIterator[int, DummyIterator[int]]{}

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

func (c ConstReverseIterator[T, It]) Get() T {
	return c.i.Prev().Get()
}

func (c ConstReverseIterator[T, It]) Equal(other ConstReverseIterator[T, It]) bool {
	return c.i.Equal(other.i)
}

// ReverseIterator implements Iterator
var _ Iterator[int, ReverseIterator[int, DummyIterator[int]]] = ReverseIterator[int, DummyIterator[int]]{}

type ReverseIterator[T any, It BidirIterator[T, It]] struct {
	i It
}

func Reverse[T any, It BidirIterator[T, It]](i It) ReverseIterator[T, It] {
	return ReverseIterator[T, It]{i: i}
}

func (c ReverseIterator[T, It]) Base() It {
	return c.i
}

func (c ReverseIterator[T, It]) Next() ReverseIterator[T, It] {
	return ReverseIterator[T, It]{i: c.i.Prev()}
}

func (c ReverseIterator[T, It]) Prev() ReverseIterator[T, It] {
	return ReverseIterator[T, It]{i: c.i.Next()}
}

func (c ReverseIterator[T, It]) Get() T {
	return c.i.Prev().Get()
}

func (c ReverseIterator[T, It]) Set(value T) {
	c.i.Prev().Set(value)
}

func (c ReverseIterator[T, It]) Equal(other ReverseIterator[T, It]) bool {
	return c.i.Equal(other.i)
}

// ConstReverseRandomIterator implements ConstRandomIterator
var _ ConstRandomIterator[int, ConstReverseRandomIterator[int, DummyIterator[int]]] = ConstReverseRandomIterator[int, DummyIterator[int]]{}

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

func (c ConstReverseRandomIterator[T, It]) Pos() int {
	return -c.i.Pos()
}

func (c ConstReverseRandomIterator[T, It]) Get() T {
	return c.i.Prev().Get()
}

func (c ConstReverseRandomIterator[T, It]) Equal(other ConstReverseRandomIterator[T, It]) bool {
	return c.i.Equal(other.i)
}

// ReverseRandomIterator implements RandomIterator
var _ RandomIterator[int, ReverseRandomIterator[int, DummyIterator[int]]] = ReverseRandomIterator[int, DummyIterator[int]]{}

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

func (c ReverseRandomIterator[T, It]) Pos() int {
	return -c.i.Pos()
}

func (c ReverseRandomIterator[T, It]) Get() T {
	return c.i.Prev().Get()
}

func (c ReverseRandomIterator[T, It]) Set(value T) {
	c.i.Prev().Set(value)
}

func (c ReverseRandomIterator[T, It]) Equal(other ReverseRandomIterator[T, It]) bool {
	return c.i.Equal(other.i)
}
