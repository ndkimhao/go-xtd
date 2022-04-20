// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic_test

import (
	"math/rand"
	"runtime"
	"sync"
	std_atomic "sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/atomic"
)

func TestNewValue(t *testing.T) {
	vInt := atomic.NewValue[int](1)
	assert.Equal(t, 1, vInt.Load())
	vString := atomic.NewValue[string]("hello")
	assert.Equal(t, "hello", vString.Load())
}

func TestValue_int(t *testing.T) {
	var v atomic.Value[int]
	assert.Equal(t, 0, v.Load())
	v.Store(42)
	assert.Equal(t, 42, v.Load())
	v.Store(84)
	assert.Equal(t, 84, v.Load())
}

func TestValue_string(t *testing.T) {
	var v atomic.Value[string]
	assert.Equal(t, "", v.Load())
	v.Store("foo")
	assert.Equal(t, "foo", v.Load())
	v.Store("barbaz")
	assert.Equal(t, "barbaz", v.Load())
}

func doTestValue_Concurrent[T comparable](t *testing.T, test []T) {
	p := 4 * runtime.GOMAXPROCS(0)
	N := int(1e5)
	if testing.Short() {
		p /= 2
		N = 1e3
	}
	var v atomic.Value[T]
	done := make(chan bool, p)
	for i := 0; i < p; i++ {
		go func() {
			r := rand.New(rand.NewSource(rand.Int63()))
			expected := true
		loop:
			for j := 0; j < N; j++ {
				x := test[r.Intn(len(test))]
				v.Store(x)
				x = v.Load()
				for _, x1 := range test {
					if x == x1 {
						continue loop
					}
				}
				t.Logf("loaded unexpected value %+v, want %+v", x, test)
				expected = false
				break
			}
			done <- expected
		}()
	}
	for i := 0; i < p; i++ {
		if !<-done {
			t.FailNow()
		}
	}
}

func TestValue_Concurrent(t *testing.T) {
	doTestValue_Concurrent(t, []uint16{uint16(0), ^uint16(0), uint16(1 + 2<<8), uint16(3 + 4<<8)})
	doTestValue_Concurrent(t, []uint32{uint32(0), ^uint32(0), uint32(1 + 2<<16), uint32(3 + 4<<16)})
	doTestValue_Concurrent(t, []uint64{uint64(0), ^uint64(0), uint64(1 + 2<<32), uint64(3 + 4<<32)})
	doTestValue_Concurrent(t, []complex64{complex(0, 0), complex(1, 2), complex(3, 4), complex(5, 6)})
}

func BenchmarkValueRead(b *testing.B) {
	var v atomic.Value[int]
	v.Store(123)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x := v.Load()
			if x != 123 {
				b.Fatalf("wrong value: got %v, want 123", x)
			}
		}
	})
}

func TestValue_Swap(t *testing.T) {
	// Int
	var vInt atomic.Value[int]
	assert.Equal(t, 0, vInt.Swap(1))
	assert.Equal(t, 1, vInt.Load())
	assert.Equal(t, 1, vInt.Swap(2))
	assert.Equal(t, 2, vInt.Load())

	// String
	var vString atomic.Value[string]
	assert.Equal(t, "", vString.Swap("hello"))
	assert.Equal(t, "hello", vString.Swap("world"))
}

func TestValueSwap_Concurrent(t *testing.T) {
	var v atomic.Value[uint64]
	var count uint64
	var g sync.WaitGroup
	var m, n uint64 = 10000, 10000
	if testing.Short() {
		m = 1000
		n = 1000
	}
	for i := uint64(0); i < m*n; i += n {
		i := i
		g.Add(1)
		go func() {
			var c uint64
			for newV := i; newV < i+n; newV++ {
				c += v.Swap(newV)
			}
			std_atomic.AddUint64(&count, c)
			g.Done()
		}()
	}
	g.Wait()
	assert.Equal(t, (m*n-1)*(m*n)/2, count+v.Load())
}

func TestValue_CompareAndSwap(t *testing.T) {
	// Int
	var vInt atomic.Value[int]
	assert.False(t, vInt.CompareAndSwap(1, 2))
	assert.Equal(t, 0, vInt.Load())
	assert.True(t, vInt.CompareAndSwap(0, 1))
	assert.Equal(t, 1, vInt.Load())
	assert.True(t, vInt.CompareAndSwap(1, 2))
	assert.Equal(t, 2, vInt.Load())
	assert.False(t, vInt.CompareAndSwap(3, 4))
	assert.Equal(t, 2, vInt.Load())

	// String
	var vString atomic.Value[string]
	assert.False(t, vString.CompareAndSwap("a", "b"))
	assert.Equal(t, "", vString.Load())
	assert.True(t, vString.CompareAndSwap("", ""))
	assert.Equal(t, "", vString.Load())
	assert.True(t, vString.CompareAndSwap("", "hello"))
	assert.Equal(t, "hello", vString.Load())
	assert.False(t, vString.CompareAndSwap("hi", "world"))
	assert.Equal(t, "hello", vString.Load())
	assert.True(t, vString.CompareAndSwap("hello", "world"))
	assert.Equal(t, "world", vString.Load())
}

func TestValueCompareAndSwap_Concurrent(t *testing.T) {
	var v atomic.Value[int]
	var w sync.WaitGroup
	v.Store(0)
	m, n := 1000, 100
	if testing.Short() {
		m = 100
		n = 100
	}
	for i := 0; i < m; i++ {
		i := i
		w.Add(1)
		go func() {
			for j := i; j < m*n; runtime.Gosched() {
				if v.CompareAndSwap(j, j+1) {
					j += m
				}
			}
			w.Done()
		}()
	}
	w.Wait()
	assert.Equal(t, m*n, v.Load())
}
