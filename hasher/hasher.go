package hasher

import (
	"unsafe"

	"github.com/ndkimhao/go-xtd/xrand"
)

var defaultSeed = uintptr(xrand.Crypto().Uint64())

// Based on https://mdlayher.com/blog/go-generics-draft-design-building-a-hashtable/

type Hasher[T comparable] func(data unsafe.Pointer, seed uintptr) uintptr

func (h Hasher[T]) Hash(value T) uint64 {
	return uint64(h(unsafe.Pointer(&value), defaultSeed))
}

func (h Hasher[T]) HashSeed(value T, seed uint64) uint64 {
	return uint64(h(unsafe.Pointer(&value), uintptr(seed)))
}

func (h Hasher[T]) HashPointer(value *T) uint64 {
	return uint64(h(unsafe.Pointer(value), defaultSeed))
}

func (h Hasher[T]) HashPointerSeed(value *T, seed uint64) uint64 {
	return uint64(h(unsafe.Pointer(value), uintptr(seed)))
}

func Of[T comparable]() Hasher[T] {
	var m interface{} = (map[T]struct{})(nil)
	return (*maptype)(*(*unsafe.Pointer)(unsafe.Pointer(&m))).hasher
}

func Hash[T comparable](value T) uint64 {
	return Of[T]().Hash(value)
}

func HashSeed[T comparable](value T, seed uint64) uint64 {
	return Of[T]().HashSeed(value, seed)
}

func HashOp[T comparable]() func(T) uint64 {
	return Of[T]().Hash
}
