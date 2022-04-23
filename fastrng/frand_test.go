package fastrng_test

import (
	cryto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
	"runtime"
	"testing"

	"github.com/ndkimhao/go-xtd/fastrng"
)

func TestRNG_Basic(t *testing.T) {
	r := fastrng.New([32]byte{1, 2, 3})
	for i := 0; i < 4; i++ {
		t.Logf("1/%d: %d", i, r.Uint64())
	}
	r = fastrng.New([32]byte{})
	for i := 0; i < 4; i++ {
		t.Logf("2/%d: %d", i, r.Uint64())
	}
	r = fastrng.New([32]byte{1, 2, 3})
	for i := 0; i < 4; i++ {
		t.Logf("3/%d: %d", i, r.Uint64())
	}
}

func BenchmarkMath_Uint64(b *testing.B) {
	v := uint64(0)
	for i := 0; i < b.N; i++ {
		v += math_rand.Uint64()
	}
	runtime.KeepAlive(&v)
}

func BenchmarkMathLocal_Uint64(b *testing.B) {
	rng := math_rand.NewSource(1).(math_rand.Source64)
	v := uint64(0)
	for i := 0; i < b.N; i++ {
		v += rng.Uint64()
	}
	runtime.KeepAlive(&v)
}

func BenchmarkCrypto_Uint64(b *testing.B) {
	v := uint64(0)
	for i := 0; i < b.N; i++ {
		var b [8]byte
		_, _ = cryto_rand.Read(b[:])
		v += binary.LittleEndian.Uint64(b[:])
	}
	runtime.KeepAlive(&v)
}

func BenchmarkFRand_Uint64(b *testing.B) {
	rng := fastrng.New([32]byte{1, 2})
	v := uint64(0)
	for i := 0; i < b.N; i++ {
		v += rng.Uint64()
	}
	runtime.KeepAlive(&v)
}

func BenchmarkMath_1M(b *testing.B) {
	sz := int64(1 << 20)
	b.SetBytes(sz)
	buf := make([]byte, sz)
	for i := 0; i < b.N; i++ {
		math_rand.Read(buf)
	}
}

func BenchmarkMathLocal_1M(b *testing.B) {
	rng := math_rand.New(math_rand.NewSource(1))
	sz := int64(1 << 20)
	b.SetBytes(sz)
	buf := make([]byte, sz)
	for i := 0; i < b.N; i++ {
		rng.Read(buf)
	}
}

func BenchmarkCrypto_1M(b *testing.B) {
	sz := int64(1 << 20)
	b.SetBytes(sz)
	buf := make([]byte, sz)
	for i := 0; i < b.N; i++ {
		_, _ = cryto_rand.Read(buf)
	}
}

func BenchmarkFRand_1M(b *testing.B) {
	rng := fastrng.New([32]byte{1, 2})
	sz := int64(1 << 20)
	b.SetBytes(sz)
	buf := make([]byte, sz)
	for i := 0; i < b.N; i++ {
		_, _ = rng.Read(buf)
	}
}
