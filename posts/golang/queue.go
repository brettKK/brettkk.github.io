package main

import "fmt"

type Queue interface {
	Enqueue(int) bool
	Dequeue() (int, bool)
	Size() int
	Cap() int
}

type QueueImpl struct {
	cap int
	size int
	head, tail int
	data []int
}

func (q *QueueImpl) Enqueue(val int) bool {
	// check queue is full or not
	if q.head == (q.tail + 1) % q.cap {
		// 队列满。cap=N的队列只能装N-1个元素。 最后一个元素为空，不然放入元素后，tail和head重合 不能判断是空还是满了。
		return false
	} 
	q.data[q.tail] = val
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
	return true
}

func (q *QueueImpl) Dequeue() (int, bool) {
	// check queue is empty or not
	if q.head == q.tail {
		return -1, false
	}
	ans := q.data[q.head]
	// or use `cap`
	q.head = (q.head + 1) % len(q.data)
	q.size--
	return ans, true
}
func (q *QueueImpl) Size() int {return q.size}
func (q *QueueImpl) Cap() int {return q.cap}
func create_queue(cap int) Queue {
	return &QueueImpl{
		cap: cap,
		size: 0,
		head: 0,
		tail: 0,
		data: make([]int, cap),
	}
}

func main() {
	var cap int = 4
	var queue Queue = create_queue(cap)
	fmt.Printf("enque=%v\n", queue.Enqueue(1))
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(5)
	queue.Enqueue(7)
	fmt.Printf("size=%d, cap=%d\n", queue.Size(), queue.Cap())
	for i := 0; i < cap; i++ {
		val, resBool := queue.Dequeue()
		fmt.Printf("i=%d, val=%d, res=%v, size=%d, cap=%d,\n", i, val, resBool, queue.Size(), queue.Cap())
	}

}