package priorityqueue

import (
	"errors"
	"fmt"
	"sync"
)

// Type of heaps
const (
	MaxHeap = 1
	MinHeap = 2
)

// Element represents an element in the priority queue
type Element interface {
	Priority() int
}

// PriorityQueue represents a generic priority queue
type PriorityQueue struct {
	elements []Element
	heapType int
	sync.Mutex
}

// NewPriorityQueue creates a new priority queue
func NewPriorityQueue(heapType int) *PriorityQueue {
	pQueue := new(PriorityQueue)
	pQueue.elements = make([]Element, 1)
	pQueue.heapType = heapType
	return pQueue
}

// Push inserts an element into the priority queue
func (p *PriorityQueue) Push(e Element) error {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	// Insert element into end of queue
	idx := len(p.elements)
	p.elements = append(p.elements, e)

	// Percolate up
	switch p.heapType {
	case MaxHeap:
		for idx != 1 && p.elements[idx].Priority() > p.elements[idx/2].Priority() {
			p.elements[idx], p.elements[idx/2] = p.elements[idx/2], p.elements[idx]
			idx = idx / 2
		}
	case MinHeap:
		for idx != 1 && p.elements[idx].Priority() < p.elements[idx/2].Priority() {
			p.elements[idx], p.elements[idx/2] = p.elements[idx/2], p.elements[idx]
			idx = idx / 2
		}
	}
	return nil
}

// Pop removes the last element in the priority queue
func (p *PriorityQueue) Pop() (Element, error) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	if len(p.elements) <= 0 {
		return nil, errors.New("priority queue is empty")
	}
	// Swap first and last element
	lastIdx := len(p.elements) - 1

	p.elements[1], p.elements[lastIdx] = p.elements[lastIdx], p.elements[1]

	idx := 1
	// Percolate down
	switch p.heapType {
	case MaxHeap:
		for idx < lastIdx {
			if lastIdx <= 2*idx || lastIdx <= 2*idx+1 {
				break
			}

			// Swap current element with the larger of its children
			idxOfLarger := 2*idx + 1
			if p.elements[2*idx].Priority() > p.elements[2*idx+1].Priority() {
				idxOfLarger = 2 * idx
			}
			if p.elements[idx].Priority() > p.elements[idxOfLarger].Priority() {
				break
			}
			p.elements[idx], p.elements[idxOfLarger] = p.elements[idxOfLarger], p.elements[idx]
			idx = idxOfLarger
		}
	case MinHeap:
		for idx < lastIdx {
			if lastIdx <= 2*idx || lastIdx <= 2*idx+1 {
				break
			}

			// Swap current element with the larger of its children
			idxOfLarger := 2*idx + 1
			if p.elements[2*idx].Priority() < p.elements[2*idx+1].Priority() {
				idxOfLarger = 2 * idx
			}
			if p.elements[idx].Priority() < p.elements[idxOfLarger].Priority() {
				break
			}
			p.elements[idx], p.elements[idxOfLarger] = p.elements[idxOfLarger], p.elements[idx]
			idx = idxOfLarger
		}
	}

	elem := p.elements[lastIdx]
	p.elements = p.elements[:lastIdx]
	return elem, nil
}

func (p *PriorityQueue) String() string {
	typeString := "Max"
	if p.heapType == MinHeap {
		typeString = "Min"
	}
	var queue []Element
	if len(p.elements) > 0 {
		queue = p.elements[1:]
	}
	return fmt.Sprintf("Queue: '%v'\tType:%s", queue, typeString)
}
