package secrand

import (
	"math/rand"
	"sync"
)

var _rngPool = sync.Pool{
	New: func() any {
		return rand.New(NewRNG())
	},
}

func Get() *rand.Rand {
	return _rngPool.Get().(*rand.Rand)
}

func Put(rng *rand.Rand) {
	_rngPool.Put(rng)
}
