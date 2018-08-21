package datastructures

import "sync"

// Queue is a linear data stucture that implements first in, first out
type Queue struct {
	container []interface{}
	sync.Mutex
}

// Enqueue adds an element to the stack
//
func (q *Queue) Enqueue(i ...interface{}) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	q.container = append(q.container, i...)
}

// Dequeue returns the first element in the list
//
func (q *Queue) Dequeue() interface{} {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	if len(q.container) == 0 {
		return nil
	}
	first := q.container[0]
	q.container = q.container[1:]

	return first
}
