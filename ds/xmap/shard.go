package xmap

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"

	"github.com/ndkimhao/go-xtd/hasher"
	"github.com/ndkimhao/go-xtd/xfn"
	"github.com/ndkimhao/go-xtd/xrand"
)

type Sharded[K comparable, V any] struct {
	shards []Shard[K, V]
	hasher xfn.BiFunction[K, uint64, uint64]
	seed   uint64
}

type Shard[K comparable, V any] struct {
	items map[K]V
	sync.RWMutex
}

var DefaultShards = runtime.NumCPU() * 2

func NewSharded[K comparable, V any]() Sharded[K, V] {
	return NewShardedCustom[K, V](DefaultShards, hasher.Of[K]().HashSeed)
}

func NewShardedCustom[K comparable, V any](shards int, hasher xfn.BiFunction[K, uint64, uint64]) Sharded[K, V] {
	if shards <= 0 {
		panic(fmt.Sprint("invalid number of shards: ", shards))
	}
	m := Sharded[K, V]{
		shards: make([]Shard[K, V], shards),
		hasher: hasher,
		seed:   xrand.Uint64(),
	}
	for i := range m.shards {
		m.shards[i].items = map[K]V{}
	}
	return m
}

func (m Sharded[K, V]) GetShard(key K) *Shard[K, V] {
	if len(m.shards) == 1 {
		return &m.shards[0]
	}
	hash := m.hasher(key, m.seed)
	return &m.shards[hash%uint64(len(m.shards))]
}

func (m Sharded[K, V]) Set(key K, value V) {
	shard := m.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

func (m Sharded[K, V]) TrySet(key K, value V) (ok bool) {
	shard := m.GetShard(key)
	shard.Lock()
	_, found := shard.items[key]
	if !found {
		shard.items[key] = value
	}
	shard.Unlock()
	return !found
}

func (m Sharded[K, V]) Get(key K) V {
	value, found := m.TryGet(key)
	if !found {
		panic(fmt.Sprint("key not found: ", key))
	}
	return value
}

func (m Sharded[K, V]) TryGet(key K) (value V, ok bool) {
	shard := m.GetShard(key)
	shard.RLock()
	val, found := shard.items[key]
	shard.RUnlock()
	return val, found
}

func (m Sharded[K, V]) GetOrSetFactory(key K, factory func() V) (value V, found bool) {
	if oldValue, ok := m.TryGet(key); ok {
		return oldValue, true
	}
	shard := m.GetShard(key)
	shard.Lock()
	defer shard.Unlock()
	if oldValue, ok := shard.items[key]; ok {
		return oldValue, true
	}
	newValue := factory()
	shard.items[key] = newValue
	return newValue, false
}

func (m Sharded[K, V]) GetOrSetDefault(key K, defaultValue V) (value V, found bool) {
	if oldValue, ok := m.TryGet(key); ok {
		return oldValue, true
	}
	shard := m.GetShard(key)
	shard.Lock()
	defer shard.Unlock()
	if oldValue, ok := shard.items[key]; ok {
		return oldValue, true
	}
	shard.items[key] = defaultValue
	return defaultValue, false
}

func (m Sharded[K, V]) Len() int {
	count := 0
	for i := 0; i < len(m.shards); i++ {
		shard := &m.shards[i] // avoid copying Shard, which contains a sync.RWMutex
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

func (m Sharded[K, V]) Empty() bool {
	for i := 0; i < len(m.shards); i++ {
		shard := &m.shards[i] // avoid copying Shard, which contains a sync.RWMutex
		shard.RLock()
		shardLen := len(shard.items)
		shard.RUnlock()
		if shardLen > 0 {
			return false
		}
	}
	return true
}

func (m Sharded[K, V]) Has(key K) bool {
	shard := m.GetShard(key)
	shard.RLock()
	_, found := shard.items[key]
	shard.RUnlock()
	return found
}

func (m Sharded[K, V]) Delete(key K) {
	shard := m.GetShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

type IterateFn[K comparable, V any] func(key K, value V)
type IterateShardFn[K comparable, V any] func(shard map[K]V)

func (m Sharded[K, V]) Iterate(fn IterateFn[K, V]) {
	for i := range m.shards {
		func() {
			shard := &m.shards[i]
			shard.RLock()
			defer shard.RUnlock()
			for key, value := range shard.items {
				fn(key, value)
			}
		}()
	}
}

func (m Sharded[K, V]) IterateShards(fn IterateShardFn[K, V]) {
	for i := range m.shards {
		func() {
			shard := &m.shards[i]
			shard.Lock()
			defer shard.Unlock()
			fn(shard.items)
		}()
	}
}

func (m Sharded[K, V]) IterateShardsReadOnly(fn IterateShardFn[K, V]) {
	for i := range m.shards {
		func() {
			shard := &m.shards[i]
			shard.RLock()
			defer shard.RUnlock()
			fn(shard.items)
		}()
	}
}

func (m Sharded[K, V]) Keys() []K {
	var keys []K
	m.IterateShardsReadOnly(func(shard map[K]V) {
		for key := range shard {
			keys = append(keys, key)
		}
	})
	return keys
}

func (m Sharded[K, V]) Values() []V {
	var values []V
	m.IterateShardsReadOnly(func(shard map[K]V) {
		for _, value := range shard {
			values = append(values, value)
		}
	})
	return values
}

func (m Sharded[K, V]) Entries() []Entry[K, V] {
	var res []Entry[K, V]
	m.IterateShardsReadOnly(func(shard map[K]V) {
		for key, value := range shard {
			res = append(res, Entry[K, V]{
				Key:   key,
				Value: value,
			})
		}
	})
	return res
}

func (m Sharded[K, V]) ToMap() map[K]V {
	ret := make(map[K]V)
	m.IterateShardsReadOnly(func(shard map[K]V) {
		for key, value := range shard {
			ret[key] = value
		}
	})
	return ret
}

func (m Sharded[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.ToMap())
}
