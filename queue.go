package datastructures

// Queue is a linear data stucture that implements first in, first out
type Queue []interface{}

// Enqueue adds an element to the stack
//
func (q Queue) Enqueue(i ...interface{}) {
	q = append(q, i...)
}

// Dequeue returns the first element in the list
//
func (q Queue) Dequeue() interface{} {
	if len(q) == 0 {
		return nil
	}
	first := q[0]
	q = q[1:]

	return first
}
