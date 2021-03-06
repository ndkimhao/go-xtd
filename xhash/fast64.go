package xhash

import (
	"math/bits"
)

type Incremental64 struct {
	h uint64
}

const (
	key64_000 uint64 = 0xbe4ba423396cfeb8
	key64_008 uint64 = 0x1cad21f72c81017c
	key64_016 uint64 = 0xdb979083e96dd4de
)

func NewFast64() Incremental64 {
	return Incremental64{h: key64_000}
}

func rrmxmx(h64 uint64) uint64 {
	h64 ^= bits.RotateLeft64(h64, 49) ^ bits.RotateLeft64(h64, 24)
	h64 *= 0x9fb21c651e98df25
	h64 ^= (h64 >> 35) + 8
	h64 *= 0x9fb21c651e98df25
	h64 ^= h64 >> 28
	return h64
}

func Uint64(v uint64) uint64 {
	keyed := v ^ (key64_008 ^ key64_016)
	return rrmxmx(keyed)
}

// Uint64Seed seed should be random (e.g., precompute seed by hashing it first)
func Uint64Seed(v uint64, seed uint64) uint64 {
	seed ^= uint64(bits.ReverseBytes32(uint32(seed))) << 32
	keyed := v ^ (key64_008 ^ key64_016 - seed)
	return rrmxmx(keyed)
}

func (f *Incremental64) WriteUint64(v uint64) {
	f.h = Uint64Seed(v, f.h)
}

func (f *Incremental64) Write(p []byte) (n int, err error) {
	i := 0
	f.WriteUint64(uint64(len(p)))
	for ; i < len(p); i += 8 {
		f.WriteUint64(toUint64(p[i : i+8]))
	}
	if i < len(p) {
		var rem [8]byte
		copy(rem[:], p[i:])
		f.WriteUint64(toUint64(rem[:]))
	}
	return i, nil
}

func (f *Incremental64) Sum(b []byte) []byte {
	var a [8]byte
	putUint64(a[:], f.Sum64())
	return append(b, a[:]...)
}

func (f *Incremental64) Reset() {
	*f = NewFast64()
}

func (f *Incremental64) Size() int {
	return 8
}

func (f *Incremental64) BlockSize() int {
	return 8
}

func (f *Incremental64) Sum64() uint64 {
	return Uint64(f.h)
}
