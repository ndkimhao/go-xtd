package priority_queue

import (
	"sync"

	"github.com/ndkimhao/go-xtd/vec"
)

type PQueue[T any] struct {
	v    vec.Vector[T]
	comp func(T, T) bool
}

func New[T any](comp func(T, T) bool, datas ...T) *PQueue[T] {
	var q PQueue[T]
	q.comp = comp

	q.PushMany(datas...)
	return &q
}

func Swap[T any](q1, q2 *PQueue[T]) {
	q1.v, q2.v = q2.v[:], q1.v[:]
	q1.comp, q2.comp = q2.comp, q1.comp
}

func (q *PQueue[T]) Push(value T) {
	var mu sync.Mutex

	mu.Lock()

	q.v.Append(value)

	if len(q.v) == 1 {
		return
	}

	var cur = len(q.v)
	var par = cur / 2

	for par-1 >= 0 {
		if q.comp(q.v[cur-1], q.v[par-1]) {
			q.v[cur-1], q.v[par-1] = q.v[par-1], q.v[cur-1]
			cur = par
			par /= 2
		} else {
			break
		}
	}
}

func (q *PQueue[T]) PushMany(values ...T) {
	for _, x := range values {
		q.Push(x)
	}
}

func (q *PQueue[T]) front() *T {
	if q.v == nil || len(q.v) == 0 {
		return nil
	}

	return &q.v[0]
}

func (q *PQueue[T]) Front() T {
	result := q.front()
	if result != nil {
		return *result
	}
	var tmp T
	return tmp
}

func (q *PQueue[T]) Pop() T {
	var result T
	if q.front() != nil {
		result = *q.front()
	} else {
		return result
	}

	q.v[0] = q.v[len(q.v)-1]
	q.v = q.v[:len(q.v)-1]

	cur := 1
	par := cur * 2

	for par-1 < len(q.v) {
		var tmp = par - 1
		if par < len(q.v) && q.comp(q.v[par], q.v[par-1]) {
			tmp = par
		}

		if q.comp(q.v[tmp], q.v[cur-1]) {
			q.v[cur-1], q.v[tmp] = q.v[tmp], q.v[cur-1]
			cur = tmp
		} else {
			break
		}

		par *= 2
	}

	return result
}

func (q *PQueue[T]) Empty() bool {
	if q.front() == nil {
		return true
	}

	return false
}

func (q *PQueue[T]) Size() int {
	if q.front() == nil {
		return 0
	}

	return len(q.v)
}

func (q *PQueue[T]) GetSlice() vec.Vector[T] {
	return q.v
}
