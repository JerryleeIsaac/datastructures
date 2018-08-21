package binarysearchtree

import (
	"fmt"
	"sync"
)

// Colors of a node of a red black tree
const (
	Uncolored   = 0
	Black       = 1
	Red         = 2
	DoubleBlack = 3
)

// Cases for binary search tree insertion
const (
	Left  = 1
	Right = 2
)

// RedBlackTree is a kind of balanced binary search tree that satisfies
// red black property
type RedBlackTree struct {
	root *Node
	sync.Mutex
}

// NewRedBlackTree creates a new red black tree instance
func NewRedBlackTree() *RedBlackTree {
	newRBT := &RedBlackTree{}
	newRBT.root = NewNilNode()
	return newRBT
}

// Insert adds an element to the red black tree
func (r *RedBlackTree) Insert(e Element) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	var newNode *Node
	r.root, newNode = Insert(r.root, e)

	r.fixInsertViolation(newNode)
}

// Remove deletes an element in the red black tree
func (r *RedBlackTree) Remove(e Element) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	var brokenNode *Node
	r.root, brokenNode = Remove(r.root, e)

	fmt.Println(brokenNode)
	r.fixRemoveViolation(brokenNode)
}

func (r *RedBlackTree) fixInsertViolation(node *Node) {
	if node.IsNil {
		return
	}
	parent := node.Parent

	// If root node, color it black
	//
	if parent.IsNil {
		node.color = Black
		r.root = node
		return
	}

	// If node is black, stop recursion
	//
	if node.color == Black {
		return
	}

	// No violation if parent's color is black
	//
	if parent.color == Black {
		return
	}

	grandParent := parent.Parent

	// This should never happen by logic
	//
	if grandParent.IsNil {
		return
	}

	// Get uncle to determine which case this is
	//
	insertCase := Right
	uncle := grandParent.LeftChild
	if parent == grandParent.LeftChild {
		uncle = grandParent.RightChild
		insertCase = Left
	}

	if uncle.IsNil || uncle.color == Black {
		if insertCase == Left {
			// Case 1: Left right case
			//
			if parent.RightChild == node {
				parent = LeftRotate(parent)
			}
			// Case 2: Left left case
			//
			grandParent.color, parent.color = parent.color, grandParent.color
			grandParent = RightRotate(grandParent)
		} else if insertCase == Right {
			// Case 3: Right left case
			//
			if parent.LeftChild == node {
				parent = RightRotate(parent)
			}
			// Case 4 Right right case
			//
			grandParent.color, parent.color = parent.color, grandParent.color
			grandParent = LeftRotate(grandParent)
		}
	} else if uncle.color == Red {
		// Recolor parent and uncle
		//
		uncle.color = Black
		parent.color = Black
		grandParent.color = Red
	}
	r.fixInsertViolation(grandParent)
}

func (r *RedBlackTree) fixRemoveViolation(node *Node) {
	if node.color == Black {
		return
	}
	// Simple case broken node is colored red
	//
	if node.color == Red {
		node.color = Black
		return
	}

	// Get sibling
	parent := node.Parent
	// This should never happen
	//
	if parent.IsNil {
		return
	}
	sibling := parent.LeftChild
	removeCase := Left
	if node == sibling {
		removeCase = Right
		sibling = parent.RightChild
	}
	// This should never happen again
	//
	fmt.Println("Node:", node, "\t\tSibling:", sibling, "\t\tparent:", parent)
	if sibling.IsNil {
		return
	}

	if sibling.color == Black {
		// Cases when sibling is black
		//
		if (sibling.LeftChild.IsNil || sibling.LeftChild.color == Black) &&
			(sibling.RightChild.IsNil || sibling.RightChild.color == Black) {
			fmt.Println("Both black case")
			// Case when both nephews are black
			//
			sibling.color = Red
			if parent.color == Black {
				parent.color = DoubleBlack
			}
		} else if removeCase == Left {
			if sibling.LeftChild.IsNil || sibling.LeftChild.color == Black {
				fmt.Println("Left right case")
				// Left Right Case
				//
				sibling.RightChild.color, sibling.color = sibling.color, sibling.RightChild.color
				sibling = LeftRotate(sibling)
			}

			fmt.Println("right right case")
			// Right right case
			//
			parent.color, sibling.color = sibling.color, parent.color
			parent = RightRotate(parent)
			parent.LeftChild.color = Black
		} else if removeCase == Right {
			if sibling.LeftChild.IsNil || sibling.LeftChild.color == Black {
				fmt.Println("right left case")
				// Right Left Case
				//
				sibling.LeftChild.color, sibling.color = sibling.color, sibling.LeftChild.color
				sibling = RightRotate(sibling)
			}

			fmt.Println("right right case")
			// Right right case
			//
			parent.color, sibling.color = sibling.color, parent.color
			parent = LeftRotate(parent)
			parent.RightChild.color = Black
		}
		r.fixRemoveViolation(parent)
	} else if sibling.color == Red {
		// Cases when sibling is red
		//
		parent.color, sibling.color = sibling.color, parent.color
		if removeCase == Left {
			fmt.Println("left case")
			// Left case
			//
			parent = RightRotate(parent)
		} else if removeCase == Right {
			fmt.Println("right case")
			// Right case
			//
			parent = LeftRotate(parent)
		}
		r.fixRemoveViolation(node)
	}
	node.color = Black
	return
}

func (r *RedBlackTree) String() string {
	if r.root.IsNil {
		return "<nil>"
	}

	return r.print2d(r.root, 0)
}

func (r *RedBlackTree) print2d(node *Node, spaces int) string {
	if node.IsNil {
		return ""
	}

	spaces += 10
	s := "\n"
	s += r.print2d(node.RightChild, spaces)

	for i := 0; i < spaces; i++ {
		s += " "
	}
	color := "r"
	if node.color == Black {
		color = "b"
	} else if node.color == DoubleBlack {
		color = "bb"
	}
	s += fmt.Sprintf("%d%s\n", node.Value(), color)

	s += r.print2d(node.LeftChild, spaces)

	return s
}
