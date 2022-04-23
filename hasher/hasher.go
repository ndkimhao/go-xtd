package hasher

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"unsafe"
)

var defaultSeed uintptr

func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot read from crypto/rand.Read")
	}
	defaultSeed = uintptr(binary.LittleEndian.Uint64(b[:]))
}

// Based on https://mdlayher.com/blog/go-generics-draft-design-building-a-hashtable/

type Hasher[T comparable] func(data unsafe.Pointer, seed uintptr) uintptr

func (h Hasher[T]) Hash(value T) uint64 {
	return uint64(h(unsafe.Pointer(&value), defaultSeed))
}

func (h Hasher[T]) HashSeed(value T, seed uint64) uint64 {
	return uint64(h(unsafe.Pointer(&value), uintptr(seed)))
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
