package binarysearchtree

import (
	"fmt"
	"time"
)

// Node of a binary search tree
type Node struct {
	Element
	LeftChild  *Node
	RightChild *Node
	Parent     *Node
	IsNil      bool

	// Red Black Tree node properties
	color int
}

func (n *Node) String() string {
	s := ""
	if n.IsNil {
		s = "<nil> <nil> <nil> "
	} else {
		s = fmt.Sprint(n.Value(), " ")
		if !n.LeftChild.IsNil {
			s += fmt.Sprint(n.LeftChild.Value(), " ")
		} else {
			s += "<nil> "
		}
		if !n.RightChild.IsNil {
			s += fmt.Sprint(n.RightChild.Value(), " ")
		} else {
			s += "<nil> "
		}
	}
	if !n.Parent.IsNil {
		s += fmt.Sprint(n.Parent.Value(), " ")
	}

	switch n.color {
	case Black:
		s += "Black "
	case Red:
		s += "Red "
	case DoubleBlack:
		s += "Double black"
	}

	return s
}

// NewNode creates a node object with default values
func NewNode(e Element) *Node {
	newNode := NewNilNode()
	newNode.Element = e
	newNode.color = Red
	newNode.IsNil = false
	return newNode
}

// NewNilNode creates a new nil node
func NewNilNode() *Node {
	newNode := &Node{IsNil: true}
	newNode.LeftChild = &Node{IsNil: true, Parent: newNode}
	newNode.RightChild = &Node{IsNil: true, Parent: newNode}
	newNode.Parent = &Node{IsNil: true}

	return newNode
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
	if root.IsNil {
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
func Find(root *Node, e Element) (*Node, bool) {
	if root.IsNil {
		return root, false
	}

	// Look for the element on the left or right subtree depending on the value of the
	// node we are looking for
	//
	if root.Value() > e.Value() {
		return Find(root.LeftChild, e)
	}
	if root.Value() < e.Value() {
		return Find(root.RightChild, e)
	}
	return root, true
}

// Remove deletes an element on a binary search tree
func Remove(root *Node, e Element) (*Node, *Node) {
	if root.IsNil {
		return root, root
	}

	var newNode *Node

	// Recursively look for the node to be deleted by following the pattern for find
	//
	if root.Value() > e.Value() {
		root.LeftChild, newNode = Remove(root.LeftChild, e)
		root.LeftChild.Parent = root
	} else if root.Value() < e.Value() {
		root.RightChild, newNode = Remove(root.RightChild, e)
		root.RightChild.Parent = root
	} else {

		// Case with no child or one child
		//
		if !root.LeftChild.IsNil && !root.RightChild.IsNil {
			// Case with 2 children
			//
			if time.Now().Unix()%2 == 0 {
				del := root.LeftChild
				for !del.RightChild.IsNil {
					del = del.RightChild
				}

				root.Element = del.Element

				root.LeftChild, newNode = Remove(root.LeftChild, del.Element)
				root.LeftChild.Parent = root
			} else {
				del := root.RightChild
				for !del.LeftChild.IsNil {
					del = del.LeftChild
				}

				root.Element = del.Element

				root.RightChild, newNode = Remove(root.RightChild, del.Element)
				root.RightChild.Parent = root
			}
		} else {
			if root.LeftChild.IsNil {
				newNode = root.RightChild
			} else if root.RightChild.IsNil {
				newNode = root.LeftChild
			}

			// Recoloring of node for rbt deletion
			//
			if root.color == Black &&
				(newNode.IsNil || newNode.color == Black) {
				newNode.color = DoubleBlack
			}
			return newNode, newNode
		}
	}

	return root, newNode
}

// LeftRotate rotates a node in a binary search tree
func LeftRotate(x *Node) *Node {
	if x.IsNil {
		return x
	}

	xRight := x.RightChild
	// Get right Child
	//
	if xRight.IsNil {
		return xRight
	}

	// Assign parent of x as parent of right child
	//
	xParent := x.Parent
	xRight.Parent = xParent
	if !xParent.IsNil {
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

	// Set right child of x to the left child of xright
	//
	x.RightChild = xRightLeft
	if !xRightLeft.IsNil {
		xRightLeft.Parent = x
	}

	return xRight
}

// RightRotate rotates a node in a binary search tree
func RightRotate(x *Node) *Node {
	if x.IsNil {
		return x
	}

	// Get left Child
	//
	xLeft := x.LeftChild
	if xLeft.IsNil {
		return xLeft
	}

	// Assign parent of x as parent of left child
	//
	xParent := x.Parent
	xLeft.Parent = xParent
	if !xParent.IsNil {
		if xParent.LeftChild == x {
			xParent.LeftChild = xLeft
		} else if xParent.RightChild == x {
			xParent.RightChild = xLeft
		}
	}

	xLeftRight := xLeft.RightChild

	// Set xleft as the parent of x
	//
	x.Parent = xLeft
	xLeft.RightChild = x

	// Set left child of x to right child of xleft
	//
	x.LeftChild = xLeftRight
	if !xLeftRight.IsNil {
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
