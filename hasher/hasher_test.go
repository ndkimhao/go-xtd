package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	a := Hash("abc")
	b := Hash("abcd")
	c := Hash("abc")
	d := Hash(123)
	assert.Equal(t, a, c)
	assert.NotEqual(t, a, b)
	assert.NotEqual(t, a, d)
}
