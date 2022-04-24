// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xsort_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/algo"
	"github.com/ndkimhao/go-xtd/constraints"
	"github.com/ndkimhao/go-xtd/slice"
	"github.com/ndkimhao/go-xtd/xfn"
	"github.com/ndkimhao/go-xtd/xsort"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func assertSorted[T constraints.Ordered](t *testing.T, s slice.Slice[T]) bool {
	return assert.Truef(t, algo.IsSortedOrdered[T](s.Begin(), s.End()), "got: %v", s)
}

func randomInts() slice.Slice[int] {
	n := 1000000
	if testing.Short() {
		n /= 100
	}
	s := slice.NewLen[int](n)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < len(s); i++ {
		s[i] = rng.Intn(100)
	}
	if algo.IsSortedOrdered[int](s.Begin(), s.End()) {
		panic("terrible rand.rand")
	}
	return s
}

func TestSortIntSlice(t *testing.T) {
	s := slice.Copy(ints[:])
	xsort.SortOrdered(s)
	assertSorted[int](t, s)
}

func TestSortFloat64Slice(t *testing.T) {
	s := slice.Copy(float64s[:])
	xsort.SortOrdered(s)
	assertSorted[float64](t, s)
}

func TestSortStringSlice(t *testing.T) {
	s := slice.Copy(strings[:])
	xsort.SortOrdered(s)
	assertSorted[string](t, s)
}

func TestSortLarge_Random(t *testing.T) {
	s := randomInts()
	xsort.SortOrdered(s)
	assertSorted[int](t, s)
}

func TestSortAnyIntSlice(t *testing.T) {
	s := slice.Copy(ints[:])
	xsort.Sort(s, xfn.Less[int])
	assertSorted[int](t, s)
}

func TestSortAnyFloat64Slice(t *testing.T) {
	s := slice.Copy(float64s[:])
	xsort.Sort(s, xfn.LessFloat[float64])
	t.Logf("got: %v", s)
	assert.Truef(t, algo.IsSorted[float64](s.Begin(), s.End(), xfn.LessFloat[float64]), "got: %v", s)
}

func TestSortAnyStringSlice(t *testing.T) {
	s := slice.Copy(strings[:])
	xsort.Sort(s, xfn.Less[string])
	assertSorted[string](t, s)
}

func TestSortAnyLarge_Random(t *testing.T) {
	s := randomInts()
	xsort.Sort(s, xfn.Less[int])
	assertSorted[int](t, s)
}

func TestSortStableIntSlice(t *testing.T) {
	s := slice.Copy(ints[:])
	xsort.StableOrdered(s)
	assertSorted[int](t, s)
}

func TestSortStableFloat64Slice(t *testing.T) {
	s := slice.Copy(float64s[:])
	xsort.StableOrdered(s)
	assertSorted[float64](t, s)
}

func TestSortStableStringSlice(t *testing.T) {
	s := slice.Copy(strings[:])
	xsort.StableOrdered(s)
	assertSorted[string](t, s)
}

func TestSortStableLarge_Random(t *testing.T) {
	s := randomInts()
	xsort.StableOrdered(s)
	assertSorted[int](t, s)
}

func TestSortStableAnyIntSlice(t *testing.T) {
	s := slice.Copy(ints[:])
	xsort.Stable(s, xfn.Less[int])
	assertSorted[int](t, s)
}

func TestSortStableAnyFloat64Slice(t *testing.T) {
	s := slice.Copy(float64s[:])
	xsort.Stable(s, xfn.LessFloat[float64])
	assert.Truef(t, algo.IsSorted[float64](s.Begin(), s.End(), xfn.LessFloat[float64]), "got: %v", s)
}

func TestSortStableAnyStringSlice(t *testing.T) {
	s := slice.Copy(strings[:])
	xsort.Stable(s, xfn.Less[string])
	assertSorted[string](t, s)
}

func TestSortStableAnyLarge_Random(t *testing.T) {
	s := randomInts()
	xsort.Stable(s, xfn.Less[int])
	assertSorted[int](t, s)
}

func TestReverseSortIntSlice(t *testing.T) {
	a := slice.Copy(ints[:])
	xsort.SortOrdered(a)
	b := slice.Copy(ints[:])
	xsort.Sort(b, xfn.Greater[int])
	assert.Equal(t, b.Reversed(), a)
}

func TestReverseSortStableIntSlice(t *testing.T) {
	a := slice.Copy(ints[:])
	xsort.StableOrdered(a)
	b := slice.Copy(ints[:])
	xsort.Stable(b, xfn.Greater[int])
	assert.Equal(t, b.Reversed(), a)
}

func TestBreakPatterns(t *testing.T) {
	// Special slice used to trigger breakPatterns.
	data := make([]int, 30)
	for i := range data {
		data[i] = 10
	}
	data[(len(data)/4)*1] = 0
	data[(len(data)/4)*2] = 1
	data[(len(data)/4)*3] = 2
	xsort.SortOrdered(data)
	assertSorted[int](t, data)
}

func TestNonDeterministicComparison(t *testing.T) {
	// Ensure that sort.SortOrdered does not panic when Less returns inconsistent results.
	// See https://golang.org/issue/14377.
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	s := slice.NewLen[int](500)
	r := rand.New(rand.NewSource(0))

	for i := 0; i < 10; i++ {
		xsort.Sort(s, func(a, b int) bool { return r.Float32() < 0.5 })
	}
}

type intPair struct{ a, b int }
type intPairs slice.Slice[intPair]

func cmpIntPair(x, y intPair) bool { return x.a < y.a }

// Record initial order in B.
func (d intPairs) initB() {
	for i := range d {
		d[i].b = i
	}
}

// InOrder checks if a-equal elements were not reordered.
func (d intPairs) inOrder() bool {
	lastA, lastB := -1, 0
	for i := 0; i < len(d); i++ {
		if lastA != d[i].a {
			lastA = d[i].a
			lastB = d[i].b
			continue
		}
		if d[i].b <= lastB {
			return false
		}
		lastB = d[i].b
	}
	return true
}

func TestStability(t *testing.T) {
	n, m := 1000000, 1000
	if testing.Short() {
		n, m = 1000, 100
	}
	data := make(intPairs, n)
	ds := slice.Slice[intPair](data)

	// random distribution
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < len(data); i++ {
		data[i].a = rng.Intn(m)
	}
	if algo.IsSorted(ds.Begin(), ds.End(), cmpIntPair) {
		t.Fatalf("terrible rand.rand")
	}
	data.initB()
	xsort.Stable(data, cmpIntPair)
	if !algo.IsSorted(ds.Begin(), ds.End(), cmpIntPair) {
		t.Errorf("Stable didn't sort %d ints", n)
	}
	if !data.inOrder() {
		t.Errorf("Stable wasn't stable on %d ints", n)
	}

	// already sorted
	data.initB()
	xsort.Stable(data, cmpIntPair)
	if !algo.IsSorted(ds.Begin(), ds.End(), cmpIntPair) {
		t.Errorf("Stable shuffled sorted %d ints (order)", n)
	}
	if !data.inOrder() {
		t.Errorf("Stable shuffled sorted %d ints (stability)", n)
	}

	// sorted reversed
	for i := 0; i < len(data); i++ {
		data[i].a = len(data) - i
	}
	data.initB()
	xsort.Stable(data, cmpIntPair)
	if !algo.IsSorted(ds.Begin(), ds.End(), cmpIntPair) {
		t.Errorf("Stable didn't sort %d ints", n)
	}
	if !data.inOrder() {
		t.Errorf("Stable wasn't stable on %d ints", n)
	}
}
