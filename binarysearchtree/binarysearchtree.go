package binarysearchtree

// BinarySearchTree is an interface containing generic methods of a binary search tree
type BinarySearchTree interface {
	Element
	LeftChild() BinarySearchTree
	RightChild() BinarySearchTree
	Parent() BinarySearchTree
	SetLeftChild(BinarySearchTree)
	SetRightChild(BinarySearchTree)
	SetParent(BinarySearchTree)
	SetElement(e Element)
	GetElement() Element
}

// Element is an element stored in a node of a binary search tree
type Element interface {
	// Value returns the value contained in an element
	Value() int
}

// Insert adds a new node to a binary search tree
func Insert(root BinarySearchTree, newNode BinarySearchTree) BinarySearchTree {
	// Create new node if root is nil
	//
	if root == nil {
		return newNode
	}

	// Insert on left or right subtree depending on the value of the new node
	//
	if root.Value() >= newNode.Value() {
		root.SetLeftChild(Insert(root.LeftChild(), newNode))
		root.LeftChild().SetParent(newNode)
	} else {
		root.SetRightChild(Insert(root.RightChild(), newNode))
		root.RightChild().SetParent(newNode)
	}
	return root
}

// Find looks for the node containing the element in the binary search tree
func Find(root BinarySearchTree, node BinarySearchTree) (BinarySearchTree, bool) {
	if root == nil {
		return nil, false
	}

	// Look for the element on the left or right subtree depending on the value of the
	// node we are looking for
	//
	if root.Value() > node.Value() {
		return Find(root.LeftChild(), node)
	}
	if root.Value() < node.Value() {
		return Find(root.RightChild(), node)
	}
	return root, true
}

// Remove deletes an element on a binary search tree
func Remove(root BinarySearchTree, node BinarySearchTree) (BinarySearchTree, bool) {
	if root == nil {
		return nil, false
	}

	// Recursively look for the node to be deleted by following the pattern for find
	//
	if root.Value() > node.Value() {
		return Remove(root.LeftChild(), node)
	}
	if root.Value() < node.Value() {
		return Remove(root.RightChild(), node)
	}

	var temp BinarySearchTree
	if root.LeftChild() != nil && root.LeftChild().Value() == node.Value() {
		// Case A: Node to be deleted is on the left child
		//
		temp = root.LeftChild()
		if temp.LeftChild() == nil && temp.RightChild() == nil {
			// Case A1: Node to be deleted has no children
			//
			root.SetLeftChild(nil)
		} else if temp.LeftChild() != nil && temp.RightChild() != nil {
			// Case A2: Node to be deleted has two children
			//
			del := temp.RightChild()
			for del.RightChild() != nil {
				del = del.RightChild()
			}
			del.Parent().SetRightChild(nil)
			temp.SetElement(del.GetElement())
		} else {
			// Case A3: Node to be deleted has only one children
			//
			if temp.LeftChild() != nil {
				root.SetLeftChild(temp.LeftChild())
				temp.LeftChild().SetParent(root)
			} else if temp.RightChild() != nil {
				root.SetLeftChild(temp.RightChild())
				temp.RightChild().SetParent(root)
			}
		}
	} else if root.RightChild() != nil && root.RightChild().Value() == node.Value() {
		// Case B: Node to be deleted is on the right child
		//
		temp = root.RightChild()
		if temp.LeftChild() == nil && temp.RightChild() == nil {
			// Case B1: Node to be deleted has no children
			//
			root.SetRightChild(nil)
		} else if temp.LeftChild() != nil && temp.RightChild() != nil {
			// Case B2: Node to be deleted has two children
			//
			del := temp.LeftChild()
			for del.LeftChild() != nil {
				del = del.LeftChild()
			}
			del.Parent().SetLeftChild(nil)
			temp.SetElement(del.GetElement())
		} else {
			// Case B3: Node to be deleted has only one children
			//
			if temp.LeftChild() != nil {
				root.SetRightChild(temp.LeftChild())
				temp.LeftChild().SetParent(root)
			} else if temp.RightChild() != nil {
				root.SetRightChild(temp.RightChild())
				temp.RightChild().SetParent(root)
			}
		}
	}
	return temp, true
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
