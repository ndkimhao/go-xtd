package xmap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/hasher"
	"github.com/ndkimhao/go-xtd/xfn"
	"github.com/ndkimhao/go-xtd/xmap"
)

type MyKey struct {
	a string
	b int
}

func TestNewCustom(t *testing.T) {
	m := xmap.NewCustom[MyKey, int](hasher.HashOp[MyKey](), xfn.Equal[MyKey])
	assert.True(t, m.Set(MyKey{a: "1"}, 1))
	assert.False(t, m.Set(MyKey{a: "1"}, 1))
	v, ok := m.Get(MyKey{a: "1"})
	assert.True(t, ok)
	assert.Equal(t, v, 1)
}
