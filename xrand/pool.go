package xrand

import (
	"math/rand"
	"sync"
)

var (
	_rngPool = sync.Pool{New: newRng}

	_masterRngLock sync.Mutex
	_masterRng     = rand.NewSource(Crypto().Int63())
)

func newRng() any {
	_masterRngLock.Lock()
	defer _masterRngLock.Unlock()
	return rand.New(rand.NewSource(_masterRng.Int63()))
}

func Get() *rand.Rand {
	return _rngPool.Get().(*rand.Rand)
}

func Put(rng *rand.Rand) {
	_rngPool.Put(rng)
}
