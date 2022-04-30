package xmap

import (
	"fmt"

	"github.com/ndkimhao/go-xtd/generics"
	"github.com/ndkimhao/go-xtd/xtd"
)

type ArrayMap[K comparable, V any] struct {
	_   xtd.NoCopy
	pos map[K]int
	kv  []Entry[K, V]
}

func New[K comparable, V any]() *ArrayMap[K, V] {
	return &ArrayMap[K, V]{}
}

func (p *ArrayMap[K, V]) lazeInit() {
	if p.pos == nil {
		p.pos = map[K]int{}
	}
}

func (p *ArrayMap[K, V]) Has(key K) bool {
	_, found := p.pos[key]
	return found
}

func (p *ArrayMap[K, V]) Get(key K) V {
	value, found := p.TryGet(key)
	if !found {
		panic(fmt.Sprint("key not found: ", key))
	}
	return value
}

func (p *ArrayMap[K, V]) TryGet(key K) (value V, ok bool) {
	if i, found := p.pos[key]; found {
		return p.kv[i].Value, true
	}
	return generics.ZeroOf[V](), false
}

func (p *ArrayMap[K, V]) Set(key K, value V) {
	p.lazeInit()
	i := len(p.kv)
	p.kv = append(p.kv, Entry[K, V]{Key: key, Value: value})
	p.pos[key] = i
}

func (p *ArrayMap[K, V]) Delete(key K) (deleted bool) {
	i, found := p.pos[key]
	if !found {
		return false
	}
	last := len(p.kv) - 1
	if i != last {
		p.kv[i] = p.kv[last]
		p.pos[p.kv[i].Key] = last
	}
	delete(p.pos, key)
	p.kv = p.kv[:last]
	return true
}

func (p *ArrayMap[K, V]) Len() int {
	return len(p.kv)
}

func (p *ArrayMap[K, V]) Empty() bool {
	return len(p.kv) == 0
}

func (p *ArrayMap[K, V]) At(i int) Entry[K, V] {
	return p.kv[i]
}

func (p *ArrayMap[K, V]) ValueAt(i int) V {
	return p.kv[i].Value
}

func (p *ArrayMap[K, V]) KeyAt(i int) K {
	return p.kv[i].Key
}

func (p *ArrayMap[K, V]) Clear() {
	Clear(p.pos)
	p.kv = nil
}
