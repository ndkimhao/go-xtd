package xhash

// Thomas Wang’s integer hash functions
// Based on https://naml.us/post/inverse-of-a-hash-function/
func mix64(v uint64) uint64 {
	v = (^v) + (v << 21) // v = (v << 21) - v - 1;
	v = v ^ (v >> 24)
	v = (v + (v << 3)) + (v << 8) // v * 265
	v = v ^ (v >> 14)
	v = (v + (v << 2)) + (v << 4) // v * 21
	v = v ^ (v >> 28)
	v = v + (v << 31)
	return v
}

type Fast64 struct {
	m0 uint64
	m1 uint64
}

const (
	fast64m0 uint64 = 8633297058295171728 // mix64(0)
	fast64m1 uint64 = 6614235796240398542 // mix64(1)
)

func NewFast64() Fast64 {
	return Fast64{m0: fast64m0, m1: fast64m1}
}

func (f *Fast64) WriteUint64(v uint64) {
	f.m0 = mix64(f.m1 ^ v)
	f.m1 = mix64(f.m0 ^ f.m1)
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
	return f.m0 ^ f.m1
}
