package iterator

// ConstIterator is an interface of const iterator
type ConstIterator[T any] interface {
	IsValid() bool
	Next() ConstIterator
	Value() T
	Clone() ConstIterator
	Equal(other ConstIterator) bool
}

// Iterator is an interface of mutable iterator
type Iterator[T any] interface {
	ConstIterator
	SetValue(value T)
}

// ConstKvIterator is an interface of const key-value type iterator
type ConstKvIterator[T any] interface {
	ConstIterator
	Key() T
}

// KvIterator is an interface of mutable key-value type iterator
type KvIterator[T any] interface {
	ConstKvIterator
	SetValue(value T)
}

// ConstBidIterator is an interface of const bidirectional iterator
type ConstBidIterator[T any] interface {
	ConstIterator
	Prev() ConstBidIterator
}

// BidIterator is an interface of mutable bidirectional iterator
type BidIterator[T any] interface {
	ConstBidIterator
	SetValue(value T)
}

// ConstKvBidIterator is an interface of const key-value type bidirectional iterator
type ConstKvBidIterator[T any] interface {
	ConstKvIterator
	Prev() ConstBidIterator
}

// KvBidIterator is an interface of mutable key-value type bidirectional iterator
type KvBidIterator[T any] interface {
	ConstKvIterator
	Prev() ConstBidIterator
	SetValue(value T)
}

// RandomAccessIterator is an interface of mutable random access iterator
type RandomAccessIterator[T any] interface {
	BidIterator
	//IteratorAt returns a new iterator at position
	IteratorAt(position int) RandomAccessIterator
	Position() int
}
