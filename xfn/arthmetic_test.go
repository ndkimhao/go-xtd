package xfn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/stream"
	"github.com/ndkimhao/go-xtd/xfn"
)

func TestPlus(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []int{2, 4, 6},
			stream.RangeN(3).Map(xfn.Plus(1)).Map(xfn.Mult(2)).Slice())
	})
}

func TestPlus2(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, 12,
			stream.RangeN(3).Map(xfn.Plus(1)).Map(xfn.Mult(2)).Reduce(0, xfn.PlusOp[int]))
	})
}
