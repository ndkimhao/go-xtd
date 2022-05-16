package xsync_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/xsync"
)

type inner struct {
	a int
}

func TestSynchronized_Lock(t *testing.T) {
	var s = xsync.Synchronized[inner]{
		Value: inner{a: 123},
	}
	x := s.Lock()
	assert.Equal(t, 123, x.a)
	assert.False(t, s.Mutex.TryLock())
	s.Unlock()
	assert.True(t, s.Mutex.TryLock())
	s.Unlock()
}
