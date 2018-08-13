package binarysearchtree

import (
	"fmt"
)

// Colors of a node of a red black tree
const (
	Uncolored = 0
	Black     = 1
	Red       = 2
)

// Cases for binary search tree insertion
const (
	Left  = 1
	Right = 2
)

// RedBlackTree is a kind of balanced binary search tree that satisfies
// red black property
//
type RedBlackTree struct {
	root *Node
}

// Insert adds an element to the red black tree
func (r *RedBlackTree) Insert(e Element) {
	newNode := &Node{Element: e, color: Red}
	r.root = Insert(r.root, newNode)

	r.fixViolation(newNode)
}

func (r *RedBlackTree) fixViolation(node *Node) {
	if node == nil {
		return
	}
	parent := node.Parent

	// If root node, color it black
	//
	if parent == nil {
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
	if grandParent == nil {
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

	if uncle == nil || uncle.color == Black {
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
	r.fixViolation(grandParent)
}

func (r RedBlackTree) String() string {
	if r.root == nil {
		return ""
	}

	return r.print2d(r.root, 0)
}

func (r RedBlackTree) print2d(node *Node, spaces int) string {
	if node == nil {
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
	}
	s += fmt.Sprintf("%d%s\n", node.Value(), color)

	s += r.print2d(node.LeftChild, spaces)

	return s
}
