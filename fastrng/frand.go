package fastrng // import "lukechampine.com/frand"

import (
	"encoding/binary"

	"github.com/aead/chacha20/chacha"
)

// An RNG is a cryptographically-strong RNG constructed from the ChaCha stream
// cipher.
type RNG struct {
	buf    []byte
	n      int
	rounds int
}

var _nonce [chacha.NonceSize]byte

// Read fills b with random data. It always returns len(b), nil.
//
// For performance reasons, calling Read once on a "large" buffer (larger than
// the RNG's internal buffer) will produce different output than calling Read
// multiple times on smaller buffers. If deterministic output is required,
// clients should call Read in a loop; when copying to an io.Writer, use
// io.CopyBuffer instead of io.Copy. Callers should also be aware that b is
// xored with random data, not directly overwritten; this means that the new
// contents of b depend on its previous contents.
func (r *RNG) Read(b []byte) (int, error) {
	if len(b) <= len(r.buf[r.n:]) {
		// can fill b entirely from buffer
		r.n += copy(b, r.buf[r.n:])
	} else if len(b) <= len(r.buf[r.n:])+len(r.buf[chacha.KeySize:]) {
		// b is larger than current buffer, but can be filled after a reseed
		n := copy(b, r.buf[r.n:])
		chacha.XORKeyStream(r.buf, r.buf, _nonce[:], r.buf[:chacha.KeySize], r.rounds)
		r.n = chacha.KeySize + copy(b[n:], r.buf[chacha.KeySize:])
	} else {
		// filling b would require multiple reseeds; instead, generate a
		// temporary key, then write directly into b using that key
		var tmpKey [chacha.KeySize]byte
		_, _ = r.Read(tmpKey[:])
		chacha.XORKeyStream(b, b, _nonce[:], tmpKey[:], r.rounds)
	}
	return len(b), nil
}

func (s *RNG) Seed(seed int64) {
	panic("RNG.Seed: reseed is not supported")
}

func (r *RNG) Int63() int64 {
	return int64(r.Uint64() & ^uint64(1<<63))
}

func (r *RNG) Uint64() uint64 {
	if 8 <= len(r.buf)-r.n {
		b := r.buf[r.n : r.n+8]
		r.n += 8
		v := binary.LittleEndian.Uint64(b)
		return v
	}
	var b [8]byte
	_, _ = r.Read(b[:])
	return binary.LittleEndian.Uint64(b[:])
}

// New returns a new RNG instance with default settings
func New(seed [32]byte) *RNG {
	return NewCustom(seed, 1024, 12)
}

// NewCustom returns a new RNG instance seeded with the provided entropy and
// using the specified buffer size and number of ChaCha rounds. It panics if
// bufsize < 32 or rounds != 8, 12 or 20.
func NewCustom(seed [32]byte, bufsize int, rounds int) *RNG {
	if bufsize < chacha.KeySize {
		panic("frand: bufsize must be at least 32")
	} else if !(rounds == 8 || rounds == 12 || rounds == 20) {
		panic("frand: rounds must be 8, 12, or 20")
	}
	buf := make([]byte, chacha.KeySize+bufsize)
	chacha.XORKeyStream(buf, buf, _nonce[:], seed[:], rounds)
	return &RNG{
		buf:    buf,
		n:      chacha.KeySize,
		rounds: rounds,
	}
}