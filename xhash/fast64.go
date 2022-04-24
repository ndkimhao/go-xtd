package xhash

import (
	"math/bits"
)

type Fast64 struct {
	h uint64
}

const (
	key64_000 uint64 = 0xbe4ba423396cfeb8
	key64_008 uint64 = 0x1cad21f72c81017c
	key64_016 uint64 = 0xdb979083e96dd4de
	key64_024 uint64 = 0x1f67b3b7a4a44072
	key64_032 uint64 = 0x78e5c0cc4ee679cb
)

func NewFast64() Fast64 {
	return Fast64{h: key64_000}
}

func rrmxmx(h64 uint64, len uint64) uint64 {
	h64 ^= bits.RotateLeft64(h64, 49) ^ bits.RotateLeft64(h64, 24)
	h64 *= 0x9fb21c651e98df25
	h64 ^= (h64 >> 35) + len
	h64 *= 0x9fb21c651e98df25
	h64 ^= h64 >> 28
	return h64
}

func (f *Fast64) WriteUint64(v uint64) {
	keyed := v ^ (key64_008 ^ key64_016)
	f.h = rrmxmx(keyed, f.h)
}

func (f *Fast64) Write(p []byte) (n int, err error) {
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

func (f *Fast64) Sum(b []byte) []byte {
	var a [8]byte
	putUint64(a[:], f.Sum64())
	return append(b, a[:]...)
}

func (f *Fast64) Reset() {
	*f = NewFast64()
}

func (f *Fast64) Size() int {
	return 8
}

func (f *Fast64) BlockSize() int {
	return 8
}

func (f *Fast64) Sum64() uint64 {
	keyed := f.h ^ (key64_024 ^ key64_032)
	return rrmxmx(keyed, 8)
}
