package xrand_test

import (
	math_rand "math/rand"
	"runtime"
	"testing"

	"github.com/ndkimhao/go-xtd/xrand"
)

func BenchmarkUint64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		v := uint64(0)
		for pb.Next() {
			v += xrand.Uint64()
		}
		runtime.KeepAlive(&v)
	})
}

func BenchmarkUint64_Local(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		rng := math_rand.New(math_rand.NewSource(0))
		v := uint64(0)
		for pb.Next() {
			v += rng.Uint64()
		}
		runtime.KeepAlive(&v)
	})
}

func BenchmarkUint64_Mutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		v := uint64(0)
		for pb.Next() {
			v += math_rand.Uint64()
		}
		runtime.KeepAlive(&v)
	})
}
