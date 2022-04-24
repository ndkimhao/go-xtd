package main

import (
	"bufio"
	"encoding/binary"
	"os"

	"github.com/ndkimhao/go-xtd/xhash"
)

func main() {
	h := xhash.NewFast64()
	w := bufio.NewWriter(os.Stdout)
	var b [8]byte
	for i := 0; ; i++ {
		v := h.Sum64()
		//h.WriteUint64(uint64(i))
		h.WriteUint64(0)
		binary.LittleEndian.PutUint64(b[:], v)
		_, _ = w.Write(b[:])
	}
}
