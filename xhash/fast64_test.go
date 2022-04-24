package xhash_test

import (
	"math"
	"math/bits"
	"math/rand"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/xhash"
	"github.com/ndkimhao/go-xtd/xmap"
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

func TestFast64_Write(t *testing.T) {
	t.Run("Partial Write", func(t *testing.T) {
		seen := xmap.NewSet[uint64]()
		b := make([]byte, 32)
		for i := 0; i <= len(b); i++ {
			h1 := xhash.NewFast64()
			_, _ = h1.Write(b[:i])
			assert.True(t, seen.TryAdd(h1.Sum64()))

			if i > 0 {
				h2 := xhash.NewFast64()
				b[i-1] = 1
				_, _ = h2.Write(b[:i])
				b[i-1] = 0
				assert.True(t, seen.TryAdd(h2.Sum64()))
			}
		}
	})
}

func TestFast64_Reset(t *testing.T) {
	h := xhash.NewFast64()
	h.WriteUint64(1)
	old := h.Sum64()
	h.Reset()
	h.WriteUint64(1)
	assert.Equal(t, old, h.Sum64())
}

func BenchmarkFast64_Write(b *testing.B) {
	b.ReportAllocs()
	sz := int64(4 << 10)
	b.SetBytes(sz)
	buf := make([]byte, sz)
	h := xhash.NewFast64()
	for i := 0; i < b.N; i++ {
		_, _ = h.Write(buf)
	}
	runtime.KeepAlive(h.Sum64())
}

func BenchmarkFast64_WriteUint64(b *testing.B) {
	b.ReportAllocs()
	h := xhash.NewFast64()
	for i := 0; i < b.N; i++ {
		h.WriteUint64(uint64(i))
	}
	runtime.KeepAlive(h.Sum64())
}

func BenchmarkFast64_Sum64(b *testing.B) {
	b.ReportAllocs()
	h := xhash.NewFast64()
	v := uint64(0)
	for i := 0; i < b.N; i++ {
		v += h.Sum64()
	}
	runtime.KeepAlive(v)
}

func BenchmarkFast64(b *testing.B) {
	b.ReportAllocs()
	v := uint64(0)
	for i := 0; i < b.N; i++ {
		h := xhash.NewFast64()
		h.WriteUint64(uint64(i))
		v += h.Sum64()
	}
	runtime.KeepAlive(v)
}
