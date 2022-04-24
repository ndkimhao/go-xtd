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
	mix uint64
}

const (
	fast64MixInit uint64 = 8633297058295171728 // u64Hash(0)
)

func NewFast64() Fast64 {
	return Fast64{mix: fast64MixInit}
}

func (f *Fast64) WriteUint64(v uint64) {
	v = mix64(v) ^ v
	f.mix = mix64(f.mix ^ v)
}

func (f *Fast64) Write(p []byte) (n int, err error) {
	i := 0
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
	return mix64(f.mix) ^ f.mix
}
