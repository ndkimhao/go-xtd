package xrand_test

import (
	math_rand "math/rand"
	"runtime"
	"testing"

	"github.com/ndkimhao/go-xtd/xrand"
)

func BenchmarkMathRand_Uint64(b *testing.B) {
	rng := math_rand.New(math_rand.NewSource(0))
	v := uint64(0)
	for i := 0; i < b.N; i++ {
		v += rng.Uint64()
	}
	runtime.KeepAlive(&v)
}

func BenchmarkUint64_parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		v := uint64(0)
		for pb.Next() {
			v += xrand.Uint64()
		}
		runtime.KeepAlive(&v)
	})
}

func BenchmarkMathRand_Uint64_Mutex_parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		v := uint64(0)
		for pb.Next() {
			v += math_rand.Uint64()
		}
		runtime.KeepAlive(&v)
	})
}
