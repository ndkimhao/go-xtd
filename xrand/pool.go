package xrand

import (
	crypto_rand "crypto/rand"
	"math/rand"
	"sync"

	"github.com/ndkimhao/go-xtd/fastrng"
)

var (
	_rngPool = sync.Pool{New: newRng}

	_masterRngLock sync.Mutex
	_masterRng     = func() *fastrng.RNG {
		var b [32]byte
		_, err := crypto_rand.Read(b[:])
		if err != nil {
			panic("cannot read from crypto/rand")
		}
		return fastrng.New(b)
	}()
)

func newRng() any {
	_masterRngLock.Lock()
	defer _masterRngLock.Unlock()
	var b [32]byte
	_, _ = _masterRng.Read(b[:])
	return rand.New(fastrng.New(b))
}

func Get() *rand.Rand {
	return _rngPool.Get().(*rand.Rand)
}

func Put(rng *rand.Rand) {
	_rngPool.Put(rng)
}
