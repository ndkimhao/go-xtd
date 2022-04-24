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

// ConstKvIterator is an interface of const key-value type iterator
type ConstKvIterator[T any, It any] interface {
	ConstIterator[T, It]
	Key() T
}

// KvIterator is an interface of mutable key-value type iterator
type KvIterator[T any, It any] interface {
	ConstKvIterator[T, It]
	SetValue(value T)
}

// ConstBidIterator is an interface of const bidirectional iterator
type ConstBidIterator[T any, It any] interface {
	ConstIterator[T, It]
	Prev() It
}

// BidIterator is an interface of mutable bidirectional iterator
type BidIterator[T any, It any] interface {
	ConstBidIterator[T, It]
	SetValue(value T)
}

// ConstKvBidIterator is an interface of const key-value type bidirectional iterator
type ConstKvBidIterator[T any, It any] interface {
	ConstKvIterator[T, It]
	Prev() It
}

// KvBidIterator is an interface of mutable key-value type bidirectional iterator
type KvBidIterator[T any, It any] interface {
	ConstKvIterator[T, It]
	Prev() It
	SetValue(value T)
}

// RandomAccessIterator is an interface of mutable random access iterator
type RandomAccessIterator[T any, It any] interface {
	BidIterator[T, It]
	Add(offset int) It
	Position() int
}
