package xhash_test

import (
	"math"
	"math/bits"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/xhash"
)

// Basic test, ensure no sticky 0/1 bits
func _testSticky(t *testing.T, f func() uint64) {
	const max uint64 = math.MaxUint64
	or := uint64(0)
	and := max
	xor := uint64(0)
	h := xhash.NewFast64()
	for i := 0; i < 100; i++ {
		h.WriteUint64(f())
		v := h.Sum64()
		or |= v
		and &= v
		xor = xor ^ v
	}
	assert.Equal(t, max, or)
	assert.Equal(t, uint64(0), and)
	t.Log("ones_count(xor) = ", bits.OnesCount64(xor))
	assert.InDelta(t, 32, bits.OnesCount64(xor), 10)
}

func TestFast64_Sticky(t *testing.T) {
	t.Run("Sequence", func(t *testing.T) {
		i := uint64(0)
		_testSticky(t, func() uint64 {
			i++
			return i
		})
	})
	t.Run("Random", func(t *testing.T) {
		rng := rand.NewSource(0).(rand.Source64)
		_testSticky(t, rng.Uint64)
	})
	t.Run("Zeros", func(t *testing.T) {
		_testSticky(t, func() uint64 { return 0 })
	})
	t.Run("Ones", func(t *testing.T) {
		_testSticky(t, func() uint64 { return 1 })
	})
	t.Run("Max", func(t *testing.T) {
		_testSticky(t, func() uint64 { return math.MaxUint64 })
	})
}
