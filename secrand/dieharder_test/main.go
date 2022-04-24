package main

import (
	"encoding/binary"
	"os"
	"strconv"

	"github.com/ndkimhao/go-xtd/secrand"
)

func main() {
	strSeed := "0"
	if len(os.Args) >= 2 {
		strSeed = os.Args[1]
	}
	seed, err := strconv.ParseInt(strSeed, 10, 64)
	if err != nil {
		panic(err)
	}
	raw := len(os.Args) >= 3 && os.Args[2] == "-raw"
	var seed32 [32]byte
	binary.LittleEndian.PutUint64(seed32[:], uint64(seed))
	rng := secrand.NewRNGFromSeed(seed32)
	var buf [499]byte
	var ibuf [8]byte
	for {
		if raw {
			_, _ = rng.Read(buf[:])
			_, _ = os.Stdout.Write(buf[:])
		} else {
			binary.LittleEndian.PutUint64(ibuf[:], rng.Uint64())
			_, _ = os.Stdout.Write(ibuf[:])
		}
	}
}
