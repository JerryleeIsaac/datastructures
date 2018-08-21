package binarysearchtree

import "fmt"

// Node of a binary search tree
type Node struct {
	Element
	LeftChild  *Node
	RightChild *Node
	Parent     *Node

	// Red Black Tree node properties
	color int
}

func (n Node) String() string {
	s := fmt.Sprint(n.Value(), " ")

	if n.LeftChild != nil {
		s += fmt.Sprint(n.LeftChild.Value(), " ")
	} else {
		s += "<nil> "
	}
	if n.RightChild != nil {
		s += fmt.Sprint(n.RightChild.Value(), " ")
	} else {
		s += "<nil> "
	}

	if n.Parent != nil {
		s += fmt.Sprint(n.Parent.Value(), " ")
	} else {
		s += "<nil> "
	}

	switch n.color {
	case Black:
		s += "Black "
	case Red:
		s += "Red "
	}

	return s
}

// NewNode creates a node object with default values
func NewNode(e Element) *Node {
	return &Node{Element: e, color: Red}
}

// Element is an element stored in a node of a binary search tree
type Element interface {
	// Value returns the value contained in an element
	Value() int
}

// Insert adds a new node to a binary search tree
func Insert(root *Node, e Element) (*Node, *Node) {
	var newNode *Node
	// Create new node if root is nil
	//
	if root == nil {
		newNode = NewNode(e)
		return newNode, newNode
	}

	// Insert on left or right subtree depending on the value of the new node
	//
	if root.Value() >= e.Value() {
		root.LeftChild, newNode = Insert(root.LeftChild, e)
		root.LeftChild.Parent = root
	} else {
		root.RightChild, newNode = Insert(root.RightChild, e)
		root.RightChild.Parent = root
	}
	return root, newNode
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
func Remove(root *Node, e Element) *Node {
	if root == nil {
		return nil
	}

	// Recursively look for the node to be deleted by following the pattern for find
	//
	if root.Value() > e.Value() {
		root.LeftChild = Remove(root.LeftChild, e)
	} else if root.Value() < e.Value() {
		root.RightChild = Remove(root.RightChild, e)
	} else {

		// Case with no child or one child
		//
		if root.LeftChild == nil {
			return root.RightChild
		} else if root.RightChild == nil {
			return root.LeftChild
		}

		// Case with 2 children
		//
		del := root.RightChild
		for del.LeftChild != nil {
			del = del.LeftChild
		}

		root.Element = del.Element

		root.RightChild = Remove(root.RightChild, del.Element)
	}

	return root
}

// LeftRotate rotates a node in a binary search tree
func LeftRotate(x *Node) *Node {
	if x == nil {
		return nil
	}

	// Get right Child
	//
	if x.RightChild == nil {
		return nil
	}
	xRight := x.RightChild

	// Assign parent of x as parent of right child
	//
	xParent := x.Parent
	xRight.Parent = xParent
	if xParent != nil {
		if xParent.LeftChild == x {
			xParent.LeftChild = xRight
		} else if xParent.RightChild == x {
			xParent.RightChild = xRight
		}
	}

	xRightLeft := xRight.LeftChild

	// Set xright as the parent of x
	//
	x.Parent = xRight
	xRight.LeftChild = x

	// Set left child of xright as right child of x
	//
	x.RightChild = xRightLeft
	if xRightLeft != nil {
		xRightLeft.Parent = x
	}

	return xRight
}

// RightRotate rotates a node in a binary search tree
func RightRotate(x *Node) *Node {
	if x == nil {
		return nil
	}

	// Get left Child
	//
	if x.LeftChild == nil {
		return nil
	}
	xLeft := x.LeftChild

	// Assign parent of x as parent of right child
	//
	xParent := x.Parent
	xLeft.Parent = xParent
	if xParent != nil {
		if xParent.LeftChild == x {
			xParent.LeftChild = xLeft
		} else if xParent.RightChild == x {
			xParent.RightChild = xLeft
		}
	}

	xLeftRight := xLeft.RightChild

	// Set xright as the parent of x
	//
	x.Parent = xLeft
	xLeft.RightChild = x

	// Set left child of xright as right child of x
	//
	x.LeftChild = xLeftRight
	if xLeftRight != nil {
		xLeftRight.Parent = x
	}

	return xLeft
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
