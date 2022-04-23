package xrand_test

import (
	cryto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
	"runtime"
	"testing"

	"github.com/ndkimhao/go-xtd/fastrng"
	"github.com/ndkimhao/go-xtd/xrand"
)

func BenchmarkMath_Uint64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		v := uint64(0)
		for pb.Next() {
			v += math_rand.Uint64()
		}
		runtime.KeepAlive(v)
	})
}

func BenchmarkCrypto_Uint64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		v := uint64(0)
		for pb.Next() {
			var b [8]byte
			_, _ = cryto_rand.Read(b[:])
			v += binary.LittleEndian.Uint64(b[:])
		}
		runtime.KeepAlive(v)
	})
}

func BenchmarkFRand_Uint64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		v := uint64(0)
		for pb.Next() {
			v += xrand.Uint64()
		}
		runtime.KeepAlive(v)
	})
}

func BenchmarkMathLocal_Uint64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		rng := math_rand.New(math_rand.NewSource(1))
		v := uint64(0)
		for pb.Next() {
			v += rng.Uint64()
		}
		runtime.KeepAlive(v)
	})
}

func BenchmarkFRandLocal_Uint64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		rng := math_rand.New(fastrng.New([32]byte{1, 2, 3}))
		v := uint64(0)
		for pb.Next() {
			v += rng.Uint64()
		}
		runtime.KeepAlive(v)
	})
}
