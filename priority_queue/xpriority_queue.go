package priority_queue

import (
	"errors"
	"fmt"

	"github.com/ndkimhao/go-xtd/ds/xslice"
	"github.com/ndkimhao/go-xtd/generics"
)

type XPriorityQueue[T any] struct {
	v    []T
	comp func(T, T) bool
}

func (q XPriorityQueue[T]) Compare(a, b T) bool {
	return q.comp(a, b)
}

func (q *XPriorityQueue[T]) Push(value T) error {
	size := len(q.v)

	//q.v.Append(value)
	q.v = append(q.v, value)

	if size >= len(q.v) {
		return errors.New(fmt.Sprintf("Push %v into queue of len %v failed", value, size))
	}

	if len(q.v) == 1 {
		return nil
	}

	var cur = len(q.v)
	var par = cur >> 1

	for par > 0 {
		if q.comp(q.v[cur-1], q.v[par-1]) {
			q.v[cur-1], q.v[par-1] = q.v[par-1], q.v[cur-1]
			cur = par
			par >>= 1
		} else {
			break
		}
	}

	return nil
}

func (q *XPriorityQueue[T]) PushMany(values ...T) error {
	for _, x := range values {
		err := q.Push(x)

		if err != nil {
			return err
		}
	}

	return nil
}

func (q XPriorityQueue[T]) Front() (T, error) {
	if len(q.v) == 0 {
		return generics.ZeroOf[T](), errors.New(fmt.Sprintf("Get front of queue of length %v failed", len(q.v)))
	}

	return q.v[0], nil
}

func (q *XPriorityQueue[T]) Pop() error {
	if q.Empty() {
		return errors.New("try to pop empty queue")
	}

	q.v[0] = q.v[len(q.v)-1]
	q.v = q.v[:len(q.v)-1]

	cur := 1
	child := cur << 1

	for child-1 < len(q.v) {
		var tmp = child
		if child < len(q.v) && q.comp(q.v[child], q.v[child-1]) {
			tmp++
		}

		if q.comp(q.v[tmp-1], q.v[cur-1]) {
			q.v[cur-1], q.v[tmp-1] = q.v[tmp-1], q.v[cur-1]
			cur = tmp
		} else {
			break
		}

		child = cur << 1
	}

	return nil
}

func (q XPriorityQueue[T]) Empty() bool {
	if len(q.v) != 0 {
		return false
	}

	return true
}

func (q XPriorityQueue[T]) Size() int {
	return len(q.v)
}

func (q XPriorityQueue[T]) GetSlice() xslice.Slice[T] {
	return q.v
}
