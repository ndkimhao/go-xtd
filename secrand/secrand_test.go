package secrand_test

import (
	"math"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/secrand"
)

// Basic test, ensure no sticky 0/1 bits
func _testSticky(t *testing.T, max uint64, f func() uint64) {
	or := uint64(0)
	and := max
	xor := uint64(0)
	for i := 0; i < 100; i++ {
		v := f()
		or |= v
		and &= v
		xor = xor ^ v
	}
	assert.Equal(t, max, or)
	assert.Equal(t, uint64(0), and)
	t.Log("ones_count(xor) = ", bits.OnesCount64(xor))
	assert.InDelta(t, 32, bits.OnesCount64(xor), 10)
}

func TestUint64(t *testing.T) {
	_testSticky(t, math.MaxUint64, secrand.Uint64)
}

func TestInt63(t *testing.T) {
	_testSticky(t, math.MaxInt64, func() uint64 { return uint64(secrand.Int63()) })
}
