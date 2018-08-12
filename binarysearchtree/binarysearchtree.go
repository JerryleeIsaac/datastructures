package binarysearchtree

// Node of a binary search tree
type Node struct {
	Element
	LeftChild  *Node
	RightChild *Node
	Parent     *Node

	Height int
	// Red Black Tree node properties
	color int
}

// Element is an element stored in a node of a binary search tree
type Element interface {
	// Value returns the value contained in an element
	Value() int
}

// Insert adds a new node to a binary search tree
func Insert(root *Node, newNode *Node) *Node {
	// Create new node if root is nil
	//
	if root == nil {
		return newNode
	}

	// Insert on left or right subtree depending on the value of the new node
	//
	if root.Value() >= newNode.Value() {
		root.LeftChild = Insert(root.LeftChild, newNode)
		root.LeftChild.Parent = root
	} else {
		root.RightChild = Insert(root.RightChild, newNode)
		root.RightChild.Parent = root
	}
	root.Height = max(height(root.LeftChild), height(root.RightChild)) + 1
	return root
}

// Find looks for the node containing the element in the binary search tree
func Find(root *Node, node *Node) (*Node, bool) {
	if root == nil {
		return nil, false
	}

	// Look for the element on the left or right subtree depending on the value of the
	// node we are looking for
	//
	if root.Value() > node.Value() {
		return Find(root.LeftChild, node)
	}
	if root.Value() < node.Value() {
		return Find(root.RightChild, node)
	}
	return root, true
}

// Remove deletes an element on a binary search tree
func Remove(root *Node, node *Node) (*Node, bool) {
	if root == nil {
		return nil, false
	}

	// Recursively look for the node to be deleted by following the pattern for find
	//
	if root.Value() > node.Value() {
		return Remove(root.LeftChild, node)
	}
	if root.Value() < node.Value() {
		return Remove(root.RightChild, node)
	}

	var temp *Node
	if root.LeftChild != nil && root.LeftChild.Value() == node.Value() {
		// Case A: Node to be deleted is on the left child
		//
		temp = root.LeftChild
		if temp.LeftChild == nil && temp.RightChild == nil {
			// Case A1: Node to be deleted has no children
			//
			root.LeftChild = nil
		} else if temp.LeftChild != nil && temp.RightChild != nil {
			// Case A2: Node to be deleted has two children
			//
			del := temp.RightChild
			for del.RightChild != nil {
				del = del.RightChild
			}
			del.Parent.RightChild = nil
			temp.Element = del.Element
		} else {
			// Case A3: Node to be deleted has only one children
			//
			if temp.LeftChild != nil {
				root.LeftChild = temp.LeftChild
				temp.LeftChild.Parent = root
			} else if temp.RightChild != nil {
				root.LeftChild = temp.RightChild
				temp.RightChild.Parent = root
			}
		}
	} else if root.RightChild != nil && root.RightChild.Value() == node.Value() {
		// Case B: Node to be deleted is on the right child
		//
		temp = root.RightChild
		if temp.LeftChild == nil && temp.RightChild == nil {
			// Case B1: Node to be deleted has no children
			//
			root.RightChild = nil
		} else if temp.LeftChild != nil && temp.RightChild != nil {
			// Case B2: Node to be deleted has two children
			//
			del := temp.LeftChild
			for del.LeftChild != nil {
				del = del.LeftChild
			}
			del.Parent.LeftChild = nil
			temp.Element = del.Element
		} else {
			// Case B3: Node to be deleted has only one children
			//
			if temp.LeftChild != nil {
				root.RightChild = temp.LeftChild
				temp.LeftChild.Parent = root
			} else if temp.RightChild != nil {
				root.RightChild = temp.RightChild
				temp.RightChild.Parent = root
			}
		}
	}
	root.Height = max(height(root.LeftChild), height(root.RightChild)) + 1
	return temp, true
}

// LeftRotate rotates a node in a binary search tree
func LeftRotate(node *Node) *Node {
	right := node.RightChild
	rightLeft := right.LeftChild

	right.LeftChild = node
	node.Parent = right

	node.RightChild = rightLeft
	rightLeft.Parent = node

	return right
}

// RightRotate rotates a node in a binary search tree
func RightRotate(node *Node) *Node {
	left := node.LeftChild
	leftRight := left.RightChild

	left.RightChild = node
	node.Parent = left

	node.LeftChild = leftRight
	leftRight.Parent = node

	return left
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
