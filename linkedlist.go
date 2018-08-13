package datastructures

// LinkedList is a generic list implementation in go lang
//
type LinkedList struct {
	first  *listNode
	last   *listNode
	Length int
}

type listNode struct {
	next  *listNode
	prev  *listNode
	Value interface{}
}

// Insert adds a new element in the linked list
//
func (l *LinkedList) Insert(val interface{}) {
	node := &listNode{Value: val}

	if l.first == nil {
		l.first = node
		l.last = node
		return
	}

	l.last.next = node
	node.prev = l.last

	l.last = node
}

// Remove removes a given element in the linked list
//
func (l *LinkedList) Remove(val interface{}) bool {
	// If empty then nothing to remove
	//
	if l.first == nil {
		return false
	}

	// If only one element then remove first and last
	//
	if l.first != nil && l.first == l.last && l.first.Value == val {
		l.first = nil
		l.last = nil
		return true
	}

	// If first element then re-assign first element
	if l.first.Value == val {
		newFirst := l.first.next
		newFirst.prev = nil
		l.first = newFirst
		return true
	}

	// If last element then re-assign last element
	if l.last.Value == val {
		newLast := l.last.prev
		newLast.next = nil
		return true
	}

	// Look for the value in the middle
	//
	for tmp := l.first.next; tmp != l.last; tmp = tmp.next {
		if tmp.Value == val {
			tmp.prev.next = tmp.next
			tmp.next.prev = tmp.prev
			return true
		}
	}

	return false
}

// At gets the element at the given index
func (l *LinkedList) At(index int) (interface{}, bool) {
	// If empty then no element
	//
	if l.first == nil {
		return nil, false
	}

	// If index is larger than size then return false
	//
	if index >= l.Length {
		return nil, false
	}

	tmp := l.first
	for ctr := 0; ctr < index; ctr++ {
		tmp = tmp.next
	}

	return tmp.Value, true
}
