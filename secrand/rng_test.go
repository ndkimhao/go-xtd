package secrand_test

import (
	"math"
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
