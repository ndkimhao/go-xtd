package iter

// ConstIterator is an interface of const iterator
type ConstIterator[T any, It any] interface {
	Next() It
	Get() T
	Equal(other It) bool
}

// Iterator is an interface of mutable iterator
type Iterator[T any, It any] interface {
	ConstIterator[T, It]
	Set(value T)
}

// ConstBidirIterator is an interface of const bidirectional iterator
type ConstBidirIterator[T any, It any] interface {
	ConstIterator[T, It]
	Prev() It
}

// BidirIterator is an interface of mutable bidirectional iterator
type BidirIterator[T any, It any] interface {
	ConstBidirIterator[T, It]
	Set(value T)
}

// ConstRandomIterator is an interface of random access iterator
type ConstRandomIterator[T any, It any] interface {
	ConstBidirIterator[T, It]
	Add(offset int) It
	Pos() int
}

// RandomIterator is an interface of mutable random access iterator
type RandomIterator[T any, It any] interface {
	ConstRandomIterator[T, It]
	Set(value T)
}
