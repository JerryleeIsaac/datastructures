package binarysearchtree

// Colors of a node of a red black tree
const (
	Black = 0
	Red   = 1
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
	newNode := &Node{Element: e}
	r.root = Insert(r.root, newNode)

	r.fixViolation(newNode)
}

func (r *RedBlackTree) fixViolation(node *Node) {
	if node == nil || node.color == Black {
		return
	}
	parent := node.Parent

	// If root node, color it black
	//
	if parent == nil {
		node.color = Black
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
	insertCase := Left
	uncle := grandParent.LeftChild
	if parent == grandParent.LeftChild {
		uncle = grandParent.RightChild
		insertCase = Right
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
			RightRotate(grandParent)
		} else if insertCase == Right {
			// Case 3: Right left case
			//
			if parent.LeftChild == node {
				parent = RightRotate(parent)
			}
			// Case 4 Right right case
			//
			LeftRotate(grandParent)
		}
		grandParent.color, parent.color = parent.color, grandParent.color
	} else if uncle.color == Red {
		// Recolor parent and uncle
		//
		uncle.color = Black
		parent.color = Black
		grandParent.color = Red
	}
	r.fixViolation(grandParent)
}
