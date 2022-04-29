package xfn_test

import (
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ndkimhao/go-xtd/xfn"
)

func TestBind_Panics(t *testing.T) {
	t.Run("invalid src", func(t *testing.T) {
		assert.PanicsWithValue(t, "invalid type: src [string] is not a function type", func() {
			xfn.Bind[func()]("abc")
		})
	})
	t.Run("invalid dest", func(t *testing.T) {
		assert.PanicsWithValue(t, "invalid type: dest [int] is not a function type", func() {
			xfn.Bind[int](func() {})
		})
	})
	t.Run("mismatch outputs", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"incompatible output: "+
				"src [func(int, string) bool] has 1 outputs but dest [func(int)] has 0",
			func() {
				xfn.Bind[func(int)](func(int, string) bool { return true })
			})
	})
	t.Run("incompatible outputs", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"incompatible output: "+
				"2nd output of src [func(int, string) (int, bool)] is [bool] "+
				"but 2nd output of dest [func(int) (int, string)] is [string]",
			func() {
				xfn.Bind[func(int) (int, string)](func(int, string) (int, bool) { return 1, true })
			})
	})
	t.Run("too many args", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"too many arguments: "+
				"src [func(int, string)] has 2 inputs but provided 3 bind arguments",
			func() {
				xfn.Bind[func()](func(int, string) {}, 1, 2, 3)
			})
	})
	t.Run("invalid placeholder", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"invalid placeholder: "+
				"dest [func(string)] has 1 inputs but provided Placeholder(2) at 1st bind argument",
			func() {
				xfn.Bind[func(string)](func(int, string) {}, xfn.P2)
			})
		assert.PanicsWithValue(t,
			"invalid placeholder: "+
				"dest [func(int, int, string)] has 3 inputs but provided Placeholder(0) at 4th bind argument",
			func() {
				xfn.Bind[func(int, int, string)](func(string, string, string, string) {}, "a", "b", "c", xfn.Placeholder(0))
			})
	})
	t.Run("placeholder type mismatch", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"type mismatch: src [func(int, time.Time)] 1st input is [int] "+
				"but provided [string] via Placeholder(1) at 1st bind argument", func() {
				xfn.Bind[func(string)](func(int, time.Time) {}, xfn.P1)
			})
	})
	t.Run("type mismatch", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"type mismatch: src [func(int, string)] 1st input is [int] "+
				"but provided [string] at 1st bind argument", func() {
				xfn.Bind[func(string)](func(int, string) {}, "abc")
			})
	})
	t.Run("more inputs mismatch", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"not enough inputs: dest [func(time.Time, string, int)] "+
				"needs 2 more inputs [float64, int32]", func() {
				xfn.Bind[func(time.Time, string, int)](func(string, int, float64, int32) {}, xfn.P2)
			})
	})
	t.Run("type mismatch remaining", func(t *testing.T) {
		assert.PanicsWithValue(t,
			"type mismatch: src [func(string, int, float64)] 3rd input is [float64] "+
				"but provided [float32] via 4th dest input [func(time.Time, string, int, float32)]", func() {
				xfn.Bind[func(time.Time, string, int, float32)](func(string, int, float64) {}, xfn.P(2))
			})
	})
	t.Run("invalid P constructor", func(t *testing.T) {
		assert.PanicsWithValue(t, "invalid index", func() {
			xfn.P(0)
		})
	})
}

type Mocked struct {
	mock.Mock
}

func (m *Mocked) Call(a ...any) (bool, error) {
	args := m.Called(a...)
	return args.Bool(0), args.Error(1)

}

func TestBind(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		now := time.Now().UnixNano()
		m := &Mocked{}
		m.On("Call", time.Duration(now), "Hi", 123, 456, 4.5).Once().Return(true, io.EOF)
		f := xfn.Bind[func(int, string, float64) (bool, error)](
			func(n time.Duration, a string, b int, c int, d float64) (bool, error) {
				return m.Call(n, a, b, c, d)
			}, now, xfn.P2, xfn.P1, 456)
		ok, err := f(123, "Hi", 4.5)
		assert.True(t, ok)
		assert.Equal(t, io.EOF, err)
	})
}
