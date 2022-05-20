package priority_queue_test

import (
	"container/heap"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndkimhao/go-xtd/priority_queue"
	"github.com/ndkimhao/go-xtd/xrand"
)

const bigN = 1000000

// check go time for O( 2 * bigN * ( Log(bigN) + allocation_time ) )
func Benchmark12Millions(b *testing.B) {
	for t := 0; t < b.N; t++ {
		tmp := make([]int, bigN)
		for i := 0; i < 2*bigN*6; i++ {
			tmp[0] = i
		}
	}
}

// benchmark push 1.000.000 int nums only
func benchmarkPushPQ(queue priority_queue.PriorityQueue[int], nums []int) {
	err := queue.PushMany(nums...)
	if err != nil {
		err = nil
	}
}

// benchmark push 1.000.000 int nums then pop all of it
func benchmarkPushPopPQ(queue priority_queue.PriorityQueue[int], nums []int) {
	err := queue.PushMany(nums...)
	if err != nil {
		err = nil
	}
	for !queue.Empty() {
		err = queue.Pop()
	}
}

func BenchmarkPushXPQ(b *testing.B) {
	nums := xrand.Perm(bigN)
	comp := func(a, b int) bool {
		return a < b
	}

	for n := 0; n < b.N; n++ {
		queue, _ := priority_queue.NewXPriorityQueue[int](comp)
		benchmarkPushPQ(queue, nums)
	}
}

func BenchmarkPushPopXPQ(b *testing.B) {
	nums := xrand.Perm(bigN)
	comp := func(a, b int) bool {
		return a < b
	}

	for n := 0; n < b.N; n++ {
		queue, _ := priority_queue.NewXPriorityQueue[int](comp)
		benchmarkPushPopPQ(queue, nums)
	}
}

func BenchmarkPushXHeap(b *testing.B) {
	nums := xrand.Perm(bigN)
	comp := func(a, b int) bool {
		return a < b
	}

	for n := 0; n < b.N; n++ {
		queue := &priority_queue.XHeapSlice[int]{comp, nil} //benchmarkPushPQ(queue, nums)
		for _, x := range nums {
			heap.Push(queue, x)
		}
	}
}

func BenchmarkPushPopXHeap(b *testing.B) {
	nums := xrand.Perm(bigN)
	comp := func(a, b int) bool {
		return a < b
	}

	for n := 0; n < b.N; n++ {
		queue := &priority_queue.XHeapSlice[int]{comp, nil}
		//benchmarkPushPopPQ(queue, nums)
		for _, x := range nums {
			heap.Push(queue, x)
		}
		for queue.Len() != 0 {
			heap.Pop(queue)
		}
	}
}

// test pushing, noted : Compare method must be Compare(a, b int) { return a < b }
func testPush(t *testing.T, queue priority_queue.PriorityQueue[int]) error {
	nums := []int{3, 7, 2, 9, 1, 6, 4, 5, 8}

	err := queue.PushMany(nums[0:3]...)
	if !assert.Empty(t, err) {
		return err
	}

	for i := 3; i < 7; i++ {
		err := queue.Push(i)
		if !assert.Empty(t, err) {
			return err
		}
	}

	err = queue.PushMany(nums[7:]...)
	if !assert.Empty(t, err) {
		return err
	}

	if !assert.Equal(t, 9, queue.Size()) {
		return errors.New(fmt.Sprintf("size is wrong after pushed: %v", queue.Size()))
	}

	return nil
}

// test popping and pushing, noted : Compare method must be Compare(a, b int) { return a < b }
func testPop(t *testing.T, queue priority_queue.PriorityQueue[int]) error {
	nums := []int{3, 7, 2, 9, 1, 6, 4, 5, 8}

	err := queue.PushMany(nums...)
	if !assert.Empty(t, err) {
		return err
	}

	prev, err := queue.Front()
	if !assert.Empty(t, err) {
		return err
	}
	err = queue.Pop()
	if !assert.Empty(t, err) {
		return err
	}
	//fmt.Println(prev)
	//fmt.Println(queue.GetSlice())

	for !queue.Empty() {
		cur, err := queue.Front()
		if !assert.Empty(t, err) {
			return err
		}
		err = queue.Pop()
		if !assert.Empty(t, err) {
			return err
		}
		//fmt.Println(cur)
		//fmt.Println(queue.GetSlice())
		//fmt.Println(queue.Empty())
		if prev > cur {
			t.Fatalf("heap not good, %v is before %v", prev, cur)
		}
		prev = cur
	}

	//fmt.Println("Done")
	return nil
}

// test whether pop return err when we over pop
func testPopErr(t *testing.T, queue priority_queue.PriorityQueue[int]) error {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	err := queue.PushMany(nums...)
	if !assert.Empty(t, err) {
		return err
	}

	prev, err := queue.Front()
	if !assert.Empty(t, err) {
		return err
	}
	err = queue.Pop()
	if !assert.Empty(t, err) {
		return err
	}
	//fmt.Println(prev)
	//fmt.Println(queue.GetSlice())

	for !queue.Empty() {
		cur, err := queue.Front()
		if !assert.Empty(t, err) {
			return err
		}
		err = queue.Pop()
		if !assert.Empty(t, err) {
			return err
		}
		//fmt.Println(cur)
		//fmt.Println(queue.GetSlice())
		//fmt.Println(queue.Empty())
		if prev > cur {
			t.Fatalf("heap not good, %v is before %v", prev, cur)
		}
		prev = cur
	}

	err = queue.Pop()
	if !assert.NotEmpty(t, err) {
		return err
	}
	//fmt.Println("Done")
	return nil
}

// test queue using generics, noted : Compare method must be Compare(a, b string) { return a < b }
func testGenericString(t *testing.T, queue priority_queue.PriorityQueue[string]) error {

	text := []string{"Ma", "Moai", "Du", "Hehe", "Hello", "Cai", "Lol"}

	for i := 0; i < 3; i++ {
		err := queue.Push(text[i])
		if !assert.Empty(t, err) {
			return err
		}
	}

	err := queue.PushMany(text[3:]...)
	if !assert.Empty(t, err) {
		return err
	}

	prev, err := queue.Front()
	if !assert.Empty(t, err) || !assert.Empty(t, queue.Pop()) {
		return errors.New("get/pop error")
	}

	for !queue.Empty() {
		cur, err := queue.Front()
		if !assert.Empty(t, err) || !assert.Empty(t, queue.Pop()) {
			return errors.New("get/pop error")
		}

		if prev > cur {
			t.Fatalf("heap not good, %v is before %v", prev, cur)
		}
		prev = cur
	}

	return nil
}

// test push and pop interchangeable, noted : Compare method must be Compare(a, b int) { return a < b }
func testPushPopPushPop(t *testing.T, queue priority_queue.PriorityQueue[int]) error {
	err := queue.PushMany(1, 2, 3, 4, 5, 6)
	if !assert.Empty(t, err) {
		return err
	}

	prev, err := queue.Front()
	if !assert.Empty(t, err) || !assert.Empty(t, queue.Pop()) {
		return errors.New("get/pop error")
	}

	for i := 0; i < 2; i++ {
		cur, err := queue.Front()
		if !assert.Empty(t, err) || !assert.Empty(t, queue.Pop()) {
			return errors.New("get/pop error")
		}

		if prev > cur {
			t.Fatalf("heap not good, %v is before %v", prev, cur)
		}
	}

	//fmt.Println(queue.GetSlice())
	for i := 7; i < 10; i++ {
		err := queue.Push(i)
		if !assert.Empty(t, err) {
			return err
		}
	}

	//fmt.Println(queue.GetSlice())

	for !queue.Empty() {
		cur, err := queue.Front()
		if !assert.Empty(t, err) || !assert.Empty(t, queue.Pop()) {
			return errors.New("get/pop error")
		}
		//fmt.Println(cur)
		if cur >= 1 && cur <= 3 {
			t.Fatalf("heap not good, item is not poped: %v", cur)
		}
	}

	return nil
}

// test push 1.000.000 int nums way 1
func pushHeavy1(t *testing.T, queue priority_queue.PriorityQueue[int], data []int) error {
	err := queue.PushMany(data...)
	assert.Empty(t, err)
	return err
}

// test push 1.000.000 int nums way 2
func pushHeavy2(t *testing.T, queue priority_queue.PriorityQueue[int], data []int) error {
	for _, x := range data {
		err := queue.Push(x)
		if !assert.Empty(t, err) {
			return err
		}
	}

	return nil
}

// test push 1.000.000 int nums way 3
func pushHeavy3(t *testing.T, queue priority_queue.PriorityQueue[int], data []int) error {
	size := len(data)

	err := queue.PushMany(data[:size/3]...)
	assert.Empty(t, err)

	for i := size / 3; i < size*2/3; i++ {
		err := queue.Push(data[i])
		if !assert.Empty(t, err) {
			return err
		}
	}

	err = queue.PushMany(data[size*2/3:]...)
	assert.Empty(t, err)

	return err
}

// test push and pop 1.000.000 nums
func testSimpleHeavy(t *testing.T, queue priority_queue.PriorityQueue[int]) error {
	size := bigN
	nums := xrand.Perm(size)

	err := queue.PushMany(nums...)
	assert.Empty(t, err)

	for ; size > 0; size-- {
		err := queue.Pop()
		if !assert.Empty(t, err) {
			return err
		}
	}

	if !assert.Equal(t, 0, queue.Size()) {
		return errors.New("not empty")
	}

	return nil
}

// extreme integrity test, noted that Compare function must be Compare(a, b int) {return a < b}
func testIntegrityHeavy(t *testing.T, queue priority_queue.PriorityQueue[int]) error {
	size := bigN
	nums := xrand.Perm(size)

	err := queue.PushMany(nums...)
	assert.Empty(t, err)
	assert.Equal(t, size, queue.Size())

	// heavy push pop routines
	for _, n := range xrand.Perm(1000) {

		// pop first n numbers ~ [0, n) and check result
		for i := 0; i < n; i++ {
			x, err := queue.Front()
			if !assert.Empty(t, err) || !assert.Equal(t, i, x) || !assert.Empty(t, queue.Pop()) {
				return errors.New("integrity Error")
			}
		}

		// push back [0, n/2)
		for i := 0; i < n/2; i++ {
			err := queue.Push(i)
			if !assert.Empty(t, err) {
				return err
			}
		}

		// pop and check first n/2 numbers
		for i := 0; i < n/2; i++ {
			x, err := queue.Front()
			if !assert.Empty(t, err) || !assert.Equal(t, i, x) || !assert.Empty(t, queue.Pop()) {
				return errors.New("integrity Error")
			}
		}

		// push back [0, n)
		for i := 0; i < n; i++ {
			err := queue.Push(i)
			if !assert.Empty(t, err) {
				return err
			}
		}
	}

	// pop and check integrity of priority queue
	for i := 0; i < size; i++ {
		x, err := queue.Front()
		if !assert.Empty(t, err) || !assert.Equal(t, i, x) || !assert.Empty(t, queue.Pop()) {
			return errors.New("integrity Error")
		}
	}

	if !assert.Equal(t, 0, queue.Size()) {
		return errors.New("integrity Error")
	}

	return nil
}

func TestPushXPQ(t *testing.T) {
	comp := func(a, b int) bool {
		return a < b
	}

	queue, err := priority_queue.NewXPriorityQueue[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, testPush(t, queue))
}

func TestPopXPQ(t *testing.T) {
	comp := func(a, b int) bool {
		return a < b
	}

	queue, err := priority_queue.NewXPriorityQueue[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, testPop(t, queue))
}

func TestPopErrXPQ(t *testing.T) {
	comp := func(a, b int) bool {
		return a < b
	}

	queue, err := priority_queue.NewXPriorityQueue[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, testPopErr(t, queue))
}

func TestGenericStringXPQ(t *testing.T) {
	comp := func(a, b string) bool {
		return a < b
	}

	queue, err := priority_queue.NewXPriorityQueue[string](comp)
	assert.Empty(t, err)
	assert.Empty(t, testGenericString(t, queue))
}

func TestPushPopPushPopXPQ(t *testing.T) {
	comp := func(a, b int) bool {
		return a < b
	}

	queue, err := priority_queue.NewXPriorityQueue[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, testPushPopPushPop(t, queue))
}

func TestHeavyPushingXPQ(t *testing.T) {
	size := bigN
	nums := xrand.Perm(size)
	comp := func(a, b int) bool {
		return a < b
	}

	queue, err := priority_queue.NewXPriorityQueue[int](comp, nums...)
	assert.Empty(t, err)

	queue, err = priority_queue.NewXPriorityQueue[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, pushHeavy1(t, queue, nums))

	queue, err = priority_queue.NewXPriorityQueue[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, pushHeavy2(t, queue, nums))

	queue, err = priority_queue.NewXPriorityQueue[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, pushHeavy3(t, queue, nums))
}

func TestSimpleHeavyXPQ(t *testing.T) {
	comp := func(a, b int) bool {
		return a < b
	}
	queue, err := priority_queue.NewXPriorityQueue[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, testSimpleHeavy(t, queue))
}

func TestIntegrityHeavyXPQ(t *testing.T) {
	comp := func(a, b int) bool {
		return a < b
	}
	queue, err := priority_queue.NewXPriorityQueue[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, testIntegrityHeavy(t, queue))
}

func TestPushXHeap(t *testing.T) {
	comp := func(a, b int) bool {
		return a < b
	}

	queue, err := priority_queue.NewXHeap[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, testPush(t, queue))
}

func TestPopXHeap(t *testing.T) {
	comp := func(a, b int) bool {
		return a < b
	}

	queue, err := priority_queue.NewXHeap[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, testPop(t, queue))
}

func TestPopErrXHeap(t *testing.T) {
	comp := func(a, b int) bool {
		return a < b
	}

	queue, err := priority_queue.NewXPriorityQueue[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, testPopErr(t, queue))
}

func TestIntegrityHeavyXHeap(t *testing.T) {
	comp := func(a, b int) bool {
		return a < b
	}
	queue, err := priority_queue.NewXHeap[int](comp)
	assert.Empty(t, err)
	assert.Empty(t, testIntegrityHeavy(t, queue))
}
