package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/slice"
)

func TestIterator(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		s := slice.Of(1, 2, 3)
		it := s.Begin()
		assert.Equal(t, &s[0], it.Ref())
		assert.Equal(t, 1, it.Get())
		it.Set(5)
		assert.Equal(t, 5, s[0])
		assert.Equal(t, &s[1], it.Next().Ref())
		assert.True(t, it.Equal(s.Begin()))
		assert.False(t, it.Equal(s.End()))
		assert.True(t, s.End().Equal(s.Begin().Add(3)))
		assert.Equal(t, &s[2], s.End().Prev().Ref())
		assert.Panics(t, func() { s.End().Ref() })
		assert.Panics(t, func() { s.Begin().Prev() })
		assert.Panics(t, func() { s.End().Next() })
		assert.Panics(t, func() { s.Begin().Add(-1) })
		assert.Panics(t, func() { s.End().Add(1) })
		assert.Equal(t, &s[0], s.Begin().Add(0).Ref())
		assert.Equal(t, &s[2], s.Begin().Add(2).Add(0).Ref())
	})
	t.Run("Empty", func(t *testing.T) {
		s := slice.Of[int]()
		assert.True(t, s.Begin().Equal(s.End()))
		assert.Panics(t, func() { s.Begin().Ref() })
	})
}
