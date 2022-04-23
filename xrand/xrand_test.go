package xrand_test

import (
	"testing"

	"github.com/ndkimhao/go-xtd/xrand"
)

func TestSmoke(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(xrand.Uint64())
	}
}
