package main

import "testing"

func BenchmarkQueue_EnqueueDequeue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := NewQueue(16)
		for j := 0; j < 1024; j++ {
			q.Enqueue(j)
		}
		for j := 0; j < 1024; j++ {
			if _, ok := q.Dequeue(); !ok {
				b.Fatal("unexpected empty queue")
			}
		}
	}
}

