package datastructures

import (
	"sync"
)

// Stack is a linear data structure that implements First-in, last-out
type Stack struct {
	container []interface{}
	sync.Mutex
}

// Push adds an element to the stack
//
func (s *Stack) Push(i ...interface{}) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.container = append(s.container, i...)
}

// Pop removes the latest element of the stack
func (s *Stack) Pop() interface{} {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	if len(s.container) == 0 {
		return nil
	}
	last := s.container[len(s.container)-1]
	s.container = s.container[:len(s.container)-1]

	return last
}
