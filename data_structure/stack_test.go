package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack(10)
	if s == nil {
		t.Fatal("NewStack() returned nil")
	}
	if s.stack == nil {
		t.Fatal("NewStack().stack should be initialized")
	}
	if len(s.stack) != 0 {
		t.Fatalf("len(NewStack().stack) = %d, want 0", len(s.stack))
	}
}

func TestStack_pop(t *testing.T) {
	type fields struct {
		stack []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		wantOK bool
		wantSt []int
	}{
		{name: "fail(empty)", fields: fields{make([]int, 0)}, want: 0, wantOK: false, wantSt: []int{}},
		{name: "success", fields: fields{[]int{1, 2, 3}}, want: 3, wantOK: true, wantSt: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				stack: tt.fields.stack,
			}
			got, gotOK := s.pop()
			if got != tt.want {
				t.Errorf("pop() = %v, want %v", got, tt.want)
			}
			if gotOK != tt.wantOK {
				t.Errorf("pop() ok = %v, want %v", gotOK, tt.wantOK)
			}
			if !reflect.DeepEqual(s.stack, tt.wantSt) {
				t.Errorf("stack after pop = %v, want %v", s.stack, tt.wantSt)
			}
		})
	}
}

func TestStack_String(t *testing.T) {
	type fields struct {
		stack []int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "empty", fields: fields{stack: []int{}}, want: "stack : []"},
		{name: "non-empty", fields: fields{stack: []int{1, 2, 3}}, want: "stack : [1 2 3]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				stack: tt.fields.stack,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStack_print(t *testing.T) {
	type fields struct {
		stack []int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "empty", fields: fields{stack: []int{}}, want: "stack : []\n"},
		{name: "non-empty", fields: fields{stack: []int{1, 2, 3}}, want: "stack : [1 2 3]\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				stack: tt.fields.stack,
			}
			oldStdout := os.Stdout
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe() error = %v", err)
			}
			os.Stdout = w

			s.print()

			if err := w.Close(); err != nil {
				t.Fatalf("w.Close() error = %v", err)
			}
			os.Stdout = oldStdout

			var buf bytes.Buffer
			if _, err := io.Copy(&buf, r); err != nil {
				t.Fatalf("io.Copy() error = %v", err)
			}
			if err := r.Close(); err != nil {
				t.Fatalf("r.Close() error = %v", err)
			}

			if got := buf.String(); got != tt.want {
				t.Errorf("print() output = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStack_push(t *testing.T) {
	type fields struct {
		stack []int
	}
	type args struct {
		e int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{name: "push to empty", fields: fields{stack: []int{}}, args: args{e: 1}, want: []int{1}},
		{name: "push to non-empty", fields: fields{stack: []int{1, 2}}, args: args{e: 3}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				stack: tt.fields.stack,
			}
			s.push(tt.args.e)
			if !reflect.DeepEqual(s.stack, tt.want) {
				t.Errorf("stack after push = %v, want %v", s.stack, tt.want)
			}
		})
	}
}
