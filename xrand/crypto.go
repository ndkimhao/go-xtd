package xrand

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() uint64 {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot read from crypto/rand")
	}
	return binary.LittleEndian.Uint64(b[:])
}

var _cryptoSource rand.Source64 = cryptoSource{}

func Crypto() rand.Source64 {
	return _cryptoSource
}
