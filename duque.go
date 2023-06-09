package deque

import "fmt"

type Deque[T interface{}] struct {
	arr  []T
	head int
	tail int
	size int
}

const minCap = 16

// 环形队列空一个位置用来判断队列满的情况

func NewDeque[T any](cap int) *Deque[T] {
	q := &Deque[T]{}

	// 初始空间分配
	if cap > minCap {
		q.arr = make([]T, cap, cap)
	} else {
		q.arr = make([]T, minCap, minCap)
	}

	return q
}

func (q *Deque[T]) Size() int {
	return q.size
}

func (q *Deque[T]) Cap() int {
	return len(q.arr)
}

func (q *Deque[T]) PutFront(v T) {
	if q.isFull() {
		q.resize()
	}
	q.head = q.prev(q.head)
	q.arr[q.head] = v
	q.size++
}

func (q *Deque[T]) PutBack(v T) {
	if q.isFull() {
		q.resize()
	}

	q.arr[q.tail] = v
	q.tail = q.next(q.tail)
	q.size++
}

func (q *Deque[T]) PopFront() {
	if q.size == 0 {
		return
	}
	q.head = q.next(q.head)
	q.size--
}

func (q *Deque[T]) PopBack() {
	if q.size == 0 {
		return
	}
	q.tail = q.prev(q.tail)
	q.size--
}

func (q *Deque[T]) InsertAt(index int, value T) {
	if q.isFull() {
		q.resize()
	}

	if index < 0 || index > q.size {
		panic("out of range")
	}

	if index == 0 {
		q.PutFront(value)
		return
	}

	if index == q.size {
		q.PutBack(value)
		return
	}

	// 保证移动的元素不会超过所有元素是一半
	if index*2 > q.size {
		// 实现head不变，index之后的全部移动
		for i := q.size; i >= index; i-- {
			real := q.realIndex(i)
			next := q.next(real)
			q.arr[next] = q.arr[real]
		}
		q.tail = q.next(q.tail)
	} else {
		// 实现tail不变，index之前的全部移动
		for i := 0; i <= index; i++ {
			real := q.realIndex(i)
			prev := q.prev(real)
			q.arr[prev] = q.arr[real]
		}
		q.head = q.prev(q.head)
	}

	real := q.realIndex(index)
	q.arr[real] = value
	q.size++
}

func (q *Deque[T]) DeleteAt(index int) {
	if q.isFull() {
		q.resize()
	}

	if index < 0 || index >= q.size {
		panic("out of range")
	}

	if index == 0 {
		q.PopFront()
		return
	}

	if index == q.size-1 {
		q.PopBack()
		return
	}

	if index*2 > q.size {
		// 实现head不变，index之后的全部移动
		for i := index; i < q.size; i++ {
			real := q.realIndex(i)
			next := q.next(real)
			q.arr[real] = q.arr[next]
		}
		q.tail = q.prev(q.tail)
	} else {
		// 实现tail不变，index之前的全部移动
		for i := index; i >= 0; i-- {
			real := q.realIndex(i)
			prev := q.prev(real)
			q.arr[real] = q.arr[prev]
		}
		q.head = q.next(q.head)
	}

	q.size--
}

func (q *Deque[T]) Front() T {
	return q.arr[q.head]
}

func (q *Deque[T]) Back() T {
	return q.arr[q.prev(q.tail)]
}

func (q *Deque[T]) At(index int) T {
	return q.arr[(q.head+index)%len(q.arr)]
}

func (q *Deque[T]) isFull() bool {
	return q.size == len(q.arr)
}

func (q *Deque[T]) resize() {
	fmt.Println("resize")
	newArr := make([]T, q.size<<1)
	// 数据移动

	if q.tail > q.head {
		copy(newArr, q.arr[q.head:q.tail])
	} else {
		copy(newArr, q.arr[q.head:])
		copy(newArr[q.size-q.head:], q.arr[:q.tail])
	}
	q.head = 0
	q.tail = q.size
	q.arr = newArr
}

func (q *Deque[T]) prev(index int) int {
	return (index - 1 + len(q.arr)) % len(q.arr)
}

func (q *Deque[T]) next(index int) int {
	return (index + 1) % len(q.arr)
}

func (q *Deque[T]) realIndex(index int) int {
	return (q.head + index) % len(q.arr)
}

func (q *Deque[T]) TestPrint() {
	fmt.Println(q.arr)
}
