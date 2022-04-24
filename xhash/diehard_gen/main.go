package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"os"

	"github.com/ndkimhao/go-xtd/xhash"
)

var (
	modeUint64     = flag.Bool("uint64", false, "Use xrand.Uint64(i)")
	modeUint64Seed = flag.Bool("uint64_seed", false, "Use xrand.Uint64Seed(0, i)")
	modeFast64     = flag.Bool("fast64", false, "Use xrand.Fast64, WriteUint64(0)")
)

func main() {
	flag.Parse()
	h := xhash.NewFast64()
	w := bufio.NewWriter(os.Stdout)
	var b [8]byte
	write := func(v uint64) {
		binary.LittleEndian.PutUint64(b[:], v)
		_, _ = w.Write(b[:])
	}
	if *modeFast64 {
		for {
			write(h.Sum64())
			h.WriteUint64(0)
		}
	} else if *modeUint64 {
		for i := uint64(0); ; i++ {
			write(xhash.Uint64(i))
		}
	} else if *modeUint64Seed {
		for i := uint64(0); ; i++ {
			write(xhash.Uint64Seed(0, i))
		}
	} else {
		panic("no mode selected")
	}
}
