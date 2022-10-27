package main

// 双端队列
type Deque interface {
	PushFront(int) bool
	PushTail(int) bool
	PopFront() (int, bool)
	PopTail() (int, bool)
	Cap() int
	Size() int
}

type DequeImpl struct {
	cap, size int
	head, tail int
	data []int
}

func create_deque(cap int) Deque {
	return &DequeImpl{
		cap: cap,
		size: 0,
		head: 0,
		tail: 0,
		data: make([]int, cap),
	}
}

func (q *DequeImpl) PushFront(val int) bool {
	// check queue is full or not
	if q.head == (q.tail + 1) % q.cap {
		return false
	}
	q.head = (q.head - 1 + q.cap) % q.cap
	q.data[q.head] = val
	q.size++
	return true
}
func (q *DequeImpl) PushTail(val int) bool {
	if q.head == (q.tail + 1) % q.cap {
		return false
	}
	q.data[q.tail] = val
	q.tail = (q.tail + 1) % q.cap
	q.size++
	return true

}
func (q *DequeImpl) PopFront() (int, bool) {
	// check queue is empty or not
	if q.head == q.tail {
		return -1, false
	}
	ans := q.data[q.head]
	q.head = (q.head + 1) % q.cap
	q.size--
	return ans, true
}
func (q *DequeImpl) PopTail() (int, bool) {
	if q.head == q.tail {
		return -1, false
	}
	q.tail = (q.tail - 1 + q.cap) % q.cap
	ans := q.data[q.tail]
	q.size--
	return ans, true
}

func (q *DequeImpl) Size() int {
	// or use return q.size 
	if q.head <= q.tail {
		return q.tail - q.head
	} 
	return q.tail + (q.cap - q.head)
}
func (q *DequeImpl) Cap() int {return q.cap}

func main() {

}