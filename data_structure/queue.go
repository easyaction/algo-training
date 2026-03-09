package main

import "fmt"

type Queue struct {
	arr  []int
	head int
}

func NewQueue(size int) *Queue {
	if size < 0 {
		size = 0
	}
	return &Queue{
		arr: make([]int, 0, size),
	}
}

func (q *Queue) Dequeue() (int, bool) {
	if q.head >= len(q.arr) {
		return 0, false
	}
	v := q.arr[q.head]
	q.head++
	return v, true
}

func (q *Queue) Enqueue(n int) {
	q.arr = append(q.arr, n)
}

func (q *Queue) values() []int {
	values := make([]int, len(q.arr)-q.head)
	copy(values, q.arr[q.head:])
	return values
}

func (q *Queue) String() string {
	return fmt.Sprintf("queue : %v", q.values())
}
