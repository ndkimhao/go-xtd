package stream_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/ds/xslice"
	"github.com/ndkimhao/go-xtd/stream"
)

func TestGenerate(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		s1 := rand.New(rand.NewSource(1))
		v1 := xslice.New[int]()
		stream.Generate(s1.Int).Skip(5).Limit(10).Collect(v1.Append)

		s2 := rand.New(rand.NewSource(1))
		v2 := xslice.New[int]()
		for i := 0; i < 5; i++ {
			s2.Int()
		}
		for i := 0; i < 10; i++ {
			v2.Append(s2.Int())
		}

		assert.Equal(t, v2, v1)
	})
}
