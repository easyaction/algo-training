package main

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue(10)
	if q == nil {
		t.Fatal("NewQueue() returned nil")
	}
	if len(q.arr) != 0 {
		t.Fatalf("len(q.arr) = %d, want 0", len(q.arr))
	}
	if cap(q.arr) != 10 {
		t.Fatalf("cap(q.arr) = %d, want 10", cap(q.arr))
	}
}

func TestNewQueue_MinCapacity(t *testing.T) {
	q := NewQueue(0)
	if cap(q.arr) != 0 {
		t.Fatalf("cap(q.arr) = %d, want 0", cap(q.arr))
	}
}

func TestQueue_append(t *testing.T) {
	q := NewQueue(1)
	q.append(1)
	q.append(2)

	want := []int{1, 2}
	if got := q.values(); !reflect.DeepEqual(got, want) {
		t.Fatalf("queue values = %v, want %v", got, want)
	}
}

func TestQueue_popleft(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		wantVal int
		wantOK  bool
		wantQ   []int
	}{
		{name: "empty", initial: []int{}, wantVal: 0, wantOK: false, wantQ: []int{}},
		{name: "non-empty", initial: []int{1, 2, 3}, wantVal: 1, wantOK: true, wantQ: []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue(1)
			for _, v := range tt.initial {
				q.append(v)
			}
			gotVal, gotOK := q.popleft()

			if gotVal != tt.wantVal {
				t.Errorf("popleft() val = %d, want %d", gotVal, tt.wantVal)
			}
			if gotOK != tt.wantOK {
				t.Errorf("popleft() ok = %v, want %v", gotOK, tt.wantOK)
			}
			if gotQ := q.values(); !reflect.DeepEqual(gotQ, tt.wantQ) {
				t.Fatalf("queue after popleft = %v, want %v", gotQ, tt.wantQ)
			}
		})
	}
}

func TestQueue_InterleavedOperations(t *testing.T) {
	q := NewQueue(2)
	q.Enqueue(1)
	q.Enqueue(2)
	if v, ok := q.Dequeue(); !ok || v != 1 {
		t.Fatalf("first Dequeue() = (%d, %v), want (1, true)", v, ok)
	}

	q.Enqueue(3)
	q.Enqueue(4)

	got := []int{}
	for {
		v, ok := q.Dequeue()
		if !ok {
			break
		}
		got = append(got, v)
	}

	want := []int{2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("dequeue order = %v, want %v", got, want)
	}
}

func TestQueue_String(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		want    string
	}{
		{name: "empty", initial: []int{}, want: "queue : []"},
		{name: "non-empty", initial: []int{1, 2, 3}, want: "queue : [1 2 3]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue(1)
			for _, v := range tt.initial {
				q.Enqueue(v)
			}
			if got := q.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}
