package xfn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/stream"
	"github.com/ndkimhao/go-xtd/xfn"
)

func TestEqual(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		assert.Equal(t, []int{2},
			stream.RangeN(4).Filter(xfn.Equal(2)).Slice())
		assert.Equal(t, []int{0, 1, 3},
			stream.RangeN(4).Filter(xfn.Equal(2).Negate()).Slice())
	})
}

func TestEqualAny(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		assert.Equal(t, []int{1, 2},
			stream.RangeN(4).Filter(xfn.EqualAny(-1, 2, 1)).Slice())
	})
}
