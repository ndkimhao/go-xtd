package secrand_test

import (
	"runtime"
	"testing"

	"github.com/ndkimhao/go-xtd/secrand"
	"github.com/ndkimhao/go-xtd/xrand"
)

func BenchmarkUint64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		v := uint64(0)
		for pb.Next() {
			v += secrand.Uint64()
		}
		runtime.KeepAlive(&v)
	})
}

func BenchmarkUint64_Local(b *testing.B) {
	var seed [32]byte
	b.RunParallel(func(pb *testing.PB) {
		rng := secrand.NewRNGFromSeed(seed)
		v := uint64(0)
		for pb.Next() {
			v += rng.Uint64()
		}
		runtime.KeepAlive(&v)
	})
}

func BenchmarkUint64_xrand(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		v := uint64(0)
		for pb.Next() {
			v += xrand.Uint64()
		}
		runtime.KeepAlive(&v)
	})
}
