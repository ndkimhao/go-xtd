package ring

type Ring[T any] struct {
	buf []T
	st  int
	len int
}

func New[T any]() Ring[T] {
	return Ring[T]{}
}

func NewLen[T any](len int) Ring[T] {
	return Ring[T]{
		buf: make([]T, len),
		st:  0,
		len: len,
	}
}

func NewLenCap[T any](len, cap int) Ring[T] {
	return Ring[T]{
		buf: make([]T, len, cap),
		st:  0,
		len: len,
	}
}

func Of[T any](values ...T) Ring[T] {
	return OfSlice(values)
}

func OfSlice[T any](values []T) Ring[T] {
	return Ring[T]{
		buf: values,
		st:  0,
		len: len(values),
	}
}

func Copy[T any](values []T) Ring[T] {
	if len(values) == 0 {
		return New[T]()
	}
	return OfSlice(append([]T(nil), values...))
}

func (r *Ring[T]) checkGrowth() {
	oldCap := len(r.buf)
	if r.len == oldCap {
		var newCap int
		switch {
		case oldCap == 0:
			newCap = 4
		case oldCap < 1024:
			newCap = oldCap * 2
		default:
			newCap = oldCap * 5 / 4 // newCap = oldCap * 1.25
		}
		newBuf := make([]T, newCap)
		r.ToSlice(newBuf)
		r.buf = newBuf
		r.st = 0
	}
}

func (r *Ring[T]) checkShrink() {
	oldCap := len(r.buf)
	if oldCap > 1024 && oldCap/8 > r.len {
		newCap := oldCap / 8
		newBuf := make([]T, newCap)
		r.ToSlice(newBuf)
		r.buf = newBuf
		r.st = 0
	}
}

func (r *Ring[T]) Push(value T) {
	r.Append(value)
}

func (r *Ring[T]) Pop() T {
	if r.Empty() {
		panic("pop from empty queu")
	}
	v := r.First()
	r.DeleteFirst()
	return v
}

func (r *Ring[T]) Append(value T) {
	r.checkGrowth()
	i := r.st + r.len
	if i > len(r.buf) {
		i -= len(r.buf)
	}
	r.buf[i] = value
	r.len++
}

func (r *Ring[T]) Prepend(value T) {
	i := r.st - 1
	if i < 0 {
		i += len(r.buf)
	}
	r.buf[i] = value
	r.st = i
	r.len++
}

func (r *Ring[T]) DeleteLast() {
	if r.len == 0 {
		panic("delete from empty queue")
	}
	i := r.st + r.len - 1
	if i >= len(r.buf) {
		i -= len(r.buf)
	}
	var zero T
	r.buf[i] = zero
	r.len--
}

func (r *Ring[T]) DeleteFirst() {
	if r.len == 0 {
		panic("delete from empty queue")
	}
	var zero T
	r.buf[r.st] = zero
	r.st++
	if r.st == len(r.buf) {
		r.st = 0
	}
}

func (r Ring[T]) Len() int {
	return r.len
}

func (r Ring[T]) Cap() int {
	return len(r.buf)
}

func (r Ring[T]) Empty() bool {
	return r.len == 0
}

func (r Ring[T]) ToSlice(s []T) []T {
	if r.len == 0 {
		return nil
	}
	if s == nil {
		s = make([]T, 0, r.Len())
	} else {
		s = s[:0]
	}
	if r.st+r.len <= len(r.buf) {
		s = append(s, r.buf[r.st:r.st+r.len]...)
	} else {
		s = append(s, r.buf[r.st:]...)
		s = append(s, r.buf[:r.st+r.len-len(r.buf)]...)
	}
	return s
}

func (r Ring[T]) AtRef(i int) *T {
	if i < 0 || r.len <= i {
		panic("index out of bound")
	}
	i += r.st
	if i >= len(r.buf) {
		i -= len(r.buf)
	}
	return &r.buf[i]
}

func (r Ring[T]) At(i int) T {
	return *r.AtRef(i)
}

func (r Ring[T]) Set(i int, value T) {
	*r.AtRef(i) = value
}

func (r Ring[T]) First() T {
	if r.len == 0 {
		panic("get from empty queue")
	}
	return r.buf[r.st]
}

func (r Ring[T]) Last() T {
	if r.len == 0 {
		panic("get from empty queue")
	}
	i := r.st + r.len - 1
	if i >= len(r.buf) {
		i -= len(r.buf)
	}
	return r.buf[i]
}

//func (r Ring[T]) IteratorAt(pos int) Iterator[T] {
//	slen := r.Len()
//	if pos < 0 || slen < pos {
//		panic("out of bound")
//	}
//	var beg *T
//	if slen > 0 {
//		beg = &r[0]
//	}
//	return Iterator[T]{pos: pos, len: slen, beg: beg}
//}
//
//func (r Ring[T]) Begin() Iterator[T] {
//	return r.IteratorAt(0)
//}
//
//func (r Ring[T]) End() Iterator[T] {
//	return r.IteratorAt(r.Len())
//}
//
//func (r Ring[T]) RBegin() iter.ReverseRandom[T, Iterator[T]] {
//	return iter.ReverseRandomIterator[T](r.End())
//}
//
//func (r Ring[T]) REnd() iter.ReverseRandom[T, Iterator[T]] {
//	return iter.ReverseRandomIterator[T](r.Begin())
//}
