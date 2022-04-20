package xtd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestZeroOf(t *testing.T) {
	tests := []struct {
		name     string
		expected any
		actual   any
	}{
		{name: "int", expected: 0, actual: ZeroOf[int]()},
		{name: "bool", expected: false, actual: ZeroOf[bool]()},
		{name: "uintptr", expected: uintptr(0), actual: ZeroOf[uintptr]()},
		{name: "map[int]string", expected: map[int]string(nil), actual: ZeroOf[map[int]string]()},
		{name: "time.Time", expected: time.Time{}, actual: ZeroOf[time.Time]()},
		{name: "*time.Time", expected: (*time.Time)(nil), actual: ZeroOf[*time.Time]()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.actual)
		})
	}
}
