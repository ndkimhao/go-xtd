package secrand_test

import (
	"runtime"
	"testing"

	"github.com/ndkimhao/go-xtd/secrand"
	"github.com/ndkimhao/go-xtd/xrand"
)

func BenchmarkUint64_parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		v := uint64(0)
		for pb.Next() {
			v += secrand.Uint64()
		}
		runtime.KeepAlive(&v)
	})
}

func BenchmarkXrand_Uint64_parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		v := uint64(0)
		for pb.Next() {
			v += xrand.Uint64()
		}
		runtime.KeepAlive(&v)
	})
}
