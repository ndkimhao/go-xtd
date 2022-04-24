package secrand_test

import (
	"math"
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/secrand"
)

func TestUint64(t *testing.T) {
	// Basic test, ensure no sticky 0/1 bits
	or := uint64(0)
	and := uint64(math.MaxUint64)
	xor := uint64(0)
	for i := 0; i < 100; i++ {
		v := secrand.Uint64()
		or |= v
		and &= v
		xor = xor ^ v
	}
	assert.Equal(t, uint64(math.MaxUint64), or)
	assert.Equal(t, uint64(0), and)
	t.Log("one_count(xor) = ", bits.OnesCount64(xor))
	assert.InDelta(t, 32, bits.OnesCount64(xor), 10)
}

func TestInt63(t *testing.T) {
	// Basic test, ensure no sticky 0/1 bits
	or := uint64(0)
	and := uint64(math.MaxUint64)
	xor := uint64(0)
	for i := 0; i < 100; i++ {
		v := uint64(secrand.Int63())
		or |= v
		and &= v
		xor = xor ^ v
	}
	assert.Equal(t, uint64(math.MaxInt64), or)
	assert.Equal(t, uint64(0), and)
	t.Log("one_count(xor) = ", bits.OnesCount64(xor))
	assert.InDelta(t, 32, bits.OnesCount64(xor), 10)
}
