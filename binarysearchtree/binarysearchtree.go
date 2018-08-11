package binarysearchtree

// BinarySearchTree is an interface containing generic methods of a binary search tree
type BinarySearchTree interface {
	Element
	LeftChild() BinarySearchTree
	RightChild() BinarySearchTree
	Parent() BinarySearchTree
	SetLeftChild(BinarySearchTree)
	SetRightChild(BinarySearchTree)
	SetElement(e Element)
	GetElement() Element
}

// Element is an element stored in a node of a binary search tree
type Element interface {
	// Value returns the value contained in an element
	Value() int
}

// Insert adds an element to a binary search tree
func Insert(b BinarySearchTree, e Element) {
	if b == nil {
		return
	}
	if b.Value() >= e.Value() {
		if b.LeftChild() == nil {
			b.LeftChild().SetElement(e)
		} else {
			Insert(b.LeftChild(), e)
		}
	} else {
		if b.RightChild() == nil {
			b.RightChild().SetElement(e)
		} else {
			Insert(b.RightChild(), e)
		}
	}

}

// Find looks for the element in the binary search tree
func Find(b BinarySearchTree, e Element) (Element, bool) {
	if b == nil {
		return nil, false
	}
	if b.Value() > e.Value() {
		return Find(b.LeftChild(), e)
	}
	if b.Value() < e.Value() {
		return Find(b.RightChild(), e)
	}
	return b.GetElement(), true
}

// Remove deletes an element on a binary search tree
func Remove(b BinarySearchTree, e Element) (Element, bool) {
	if b == nil {
		return nil, false
	}
	if b.Value() > e.Value() {
		return Remove(b.LeftChild(), e)
	}
	if b.Value() < e.Value() {
		return Remove(b.RightChild(), e)
	}

	var node BinarySearchTree
	var elem Element
	if b.LeftChild() != nil && b.LeftChild().Value() == e.Value() {
		node = b.LeftChild()
		elem = node.GetElement()
		if node.LeftChild() == nil && node.RightChild() == nil {
			b.SetLeftChild(nil)
		} else if node.LeftChild() != nil && node.RightChild() != nil {
			del := node.RightChild()
			for del.RightChild() != nil {
				del = del.RightChild()
			}
			del.Parent().SetRightChild(nil)
			node.SetElement(del.GetElement())
		} else if node.LeftChild() != nil {
			b.SetLeftChild(node.LeftChild())
		} else if node.RightChild() != nil {
			b.SetLeftChild(node.RightChild())
		}
	} else if b.RightChild() != nil && b.RightChild().Value() == e.Value() {
		node = b.RightChild()
		elem = node.GetElement()
		if node.LeftChild() == nil && node.RightChild() == nil {
			b.SetRightChild(nil)
		} else if node.LeftChild() != nil && node.RightChild() != nil {
			del := node.LeftChild()
			for del.LeftChild() != nil {
				del = del.LeftChild()
			}
			del.Parent().SetLeftChild(nil)
			node.SetElement(del.GetElement())
		} else if node.LeftChild() != nil {
			b.SetRightChild(node.LeftChild())
		} else if node.RightChild() != nil {
			b.SetRightChild(node.LeftChild())
		}
	}
	return elem, true
}

func leftRotate(b BinarySearchTree) BinarySearchTree {
	x := b.RightChild()
	z := x.LeftChild()

	x.SetLeftChild(b)
	b.SetRightChild(z)

	return x
}

func rightRotate(b BinarySearchTree) BinarySearchTree {
	x := b.LeftChild()
	z := x.RightChild()

	x.SetRightChild(b)
	b.SetLeftChild(z)

	return x
}
