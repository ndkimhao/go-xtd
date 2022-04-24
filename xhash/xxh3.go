package xhash

import (
	"github.com/zeebo/xxh3"
)

type Uint128 [2]uint64

func Bytes(b []byte) uint64 {
	return xxh3.Hash(b)
}

func String(s string) uint64 {
	return xxh3.HashString(s)
}

func BytesSeed(b []byte, seed uint64) uint64 {
	return xxh3.HashSeed(b, seed)

}

func StringSeed(s string, seed uint64) uint64 {
	return xxh3.HashStringSeed(s, seed)
}

func conv128(h xxh3.Uint128) Uint128 {
	return [2]uint64{h.Lo, h.Hi}
}

func Bytes128(b []byte) Uint128 {
	return conv128(xxh3.Hash128(b))
}

func BytesSeed128(b []byte, seed uint64) Uint128 {
	return conv128(xxh3.Hash128Seed(b, seed))
}

func String128(s string) Uint128 {
	return conv128(xxh3.HashString128(s))
}

func StringSeed128(s string, seed uint64) Uint128 {
	return conv128(xxh3.HashString128Seed(s, seed))
}
