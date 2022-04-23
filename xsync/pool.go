package xsync

import (
	"sync"
)

type Pool[T any] struct {
	sync.Pool
}

func (p *Pool[T]) Put(x *T) {
	p.Pool.Put(x)
}

func (p *Pool[T]) Get() *T {
	return p.Pool.Get().(*T)
}
