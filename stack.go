package datastructures

// Stack is a linear data structure that implements First-in, last-out
type Stack []interface{}

// Push adds an element to the stack
//
func (s Stack) Push(i ...interface{}) {
	s = append(s, i...)
}

// Pop removes the latest element of the stack
func (s Stack) Pop() interface{} {
	if len(s) == 0 {
		return nil
	}
	last := s[len(s)-1]
	s = s[:len(s)-1]

	return last
}
