package secrand_test

import (
	"math"
	math_rand "math/rand"
	"runtime"
	"testing"

	"github.com/ndkimhao/go-xtd/secrand"
)

func _newRNG() *secrand.RNG {
	var seed = [32]byte{1, 2, 3}
	return secrand.NewRNGFromSeed(seed)
}

func TestRNG_Uint64(t *testing.T) {
	rng := _newRNG()
	_testSticky(t, math.MaxUint64, rng.Uint64)
}

func TestRNG_Int63(t *testing.T) {
	rng := _newRNG()
	_testSticky(t, math.MaxInt64, func() uint64 { return uint64(rng.Int63()) })
}

func TestRNG_Uint64_print(t *testing.T) {
	t.Run("Default_Nonce", func(t *testing.T) {
		rng := _newRNG()
		for i := 0; i < 5; i++ {
			t.Log(rng.Uint64())
		}
	})
	t.Run("Zero_Nonce", func(t *testing.T) {
		var seed = [32]byte{1, 2, 3}
		var nonce [8]byte
		rng := secrand.NewRNGFromSeedNonce(seed, nonce)
		for i := 0; i < 5; i++ {
			t.Log(rng.Uint64())
		}
	})
}

func BenchmarkRNG_Uint64(b *testing.B) {
	var seed [32]byte
	var nonce [8]byte
	var rng *secrand.RNG
	f := func(b *testing.B) {
		v := uint64(0)
		for i := 0; i < b.N; i++ {
			v += rng.Uint64()
		}
		runtime.KeepAlive(&v)
	}

	rng = secrand.NewRNGFromSeed(seed)
	b.Run("1k buffer", f)
	rng = secrand.NewRNGCustom(seed, nonce, 128<<10, 12)
	b.Run("128k buffer", f)
}

func BenchmarkMathRand_Uint64(b *testing.B) {
	rng := math_rand.New(math_rand.NewSource(0))
	v := uint64(0)
	for i := 0; i < b.N; i++ {
		v += rng.Uint64()
	}
	runtime.KeepAlive(&v)
}
