package iter

// ConstIterator is an interface of const iterator
type ConstIterator[T any, It any] interface {
	Next() It
	Value() T
	Equal(other It) bool
}

// Iterator is an interface of mutable iterator
type Iterator[T any, It any] interface {
	ConstIterator[T, It]
	SetValue(value T)
}

// ConstBidirIterator is an interface of const bidirectional iterator
type ConstBidirIterator[T any, It any] interface {
	ConstIterator[T, It]
	Prev() It
}

// BidirIterator is an interface of mutable bidirectional iterator
type BidirIterator[T any, It any] interface {
	ConstBidirIterator[T, It]
	SetValue(value T)
}

// ConstRandomIterator is an interface of random access iterator
type ConstRandomIterator[T any, It any] interface {
	ConstBidirIterator[T, It]
	Add(offset int) It
	Position() int
}

// RandomIterator is an interface of mutable random access iterator
type RandomIterator[T any, It any] interface {
	ConstRandomIterator[T, It]
	SetValue(value T)
}
