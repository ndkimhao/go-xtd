package priority_queue

import (
	"strings"
	"testing"
)

func TestPush(t *testing.T) {
	queue := New[int](
		func(a, b int) bool {
			return a < b
		},
		3, 7, 2,
	)

	queue.Push(9)
	queue.Push(1)
	queue.Push(6)
	queue.Push(4)

	queue.PushMany(5, 8)

	if len(queue.GetSlice()) != 9 {
		t.Fatalf("queue not inputing correctly:\n pushed %v got %v", 9, len(queue.GetSlice()))
	}
}

func TestPop(t *testing.T) {
	queue := New[int](
		func(a, b int) bool {
			return a < b
		},
		3, 7, 2,
	)

	queue.Push(9)
	queue.Push(1)
	queue.Push(6)
	queue.Push(4)

	queue.PushMany(5, 8)

	//fmt.Println(queue.GetSlice())

	prev := queue.Pop()
	//fmt.Println(prev)
	//fmt.Println(queue.GetSlice())

	for !queue.Empty() {
		cur := queue.Pop()
		//fmt.Println(cur)
		//fmt.Println(queue.GetSlice())
		//fmt.Println(queue.Empty())
		if prev > cur {
			t.Fatalf("heap not good, %v is before %v", prev, cur)
		}
		prev = cur
	}

	//fmt.Println("Done")
}

func TestGenericString(t *testing.T) {
	queue := New[string](
		func(a, b string) bool {
			if strings.Compare(a, b) < 0 {
				return true
			}

			return false
		},
	)

	queue.Push("Ma")
	queue.Push("Moai")
	queue.Push("Du")

	queue.PushMany("Hehe", "Hello", "Cai", "Lol")

	prev := queue.Pop()

	for !queue.Empty() {
		cur := queue.Pop()
		if prev > cur {
			t.Fatalf("heap not good, %v is before %v", prev, cur)
		}
		prev = cur
	}
}

func TestPushPopPushPop(t *testing.T) {
	queue := New[int](
		func(a, b int) bool {
			return a > b
		},
		1, 2, 3, 4, 5, 6,
	)

	//fmt.Println(queue.GetSlice())

	for i := 0; i < 3; i++ {
		queue.Pop()
	}

	//fmt.Println(queue.GetSlice())

	queue.Push(7)
	queue.Push(8)
	queue.Push(9)

	//fmt.Println(queue.GetSlice())

	for !queue.Empty() {
		cur := queue.Pop()
		//fmt.Println(cur)
		if cur >= 4 && cur <= 6 {
			t.Fatalf("heap not good, item is not poped: %v", cur)
		}
	}
}
