package main

import (
	"fmt"
)

type Stack struct {
	stack []int
}

func NewStack(size int) *Stack {
	if size < 0 {
		size = 0
	}
	return &Stack{stack: make([]int, 0, size)}
}

func (s *Stack) pop() (int, bool) {
	if len(s.stack) == 0 {
		return 0, false
	}
	e := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return e, true
}

func (s *Stack) push(e int) {
	s.stack = append(s.stack, e)
}

func (s *Stack) String() string {
	return fmt.Sprintf("stack : %v", s.stack)
}

func (s *Stack) print() {
	fmt.Println(s.String())
}
