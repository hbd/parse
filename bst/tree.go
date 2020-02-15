// Package bst implements a binary search tree.

package bst

import (
	"fmt"
	"strconv"
	"strings"
)

// Tree represents a BST, starting at the top node.
// It's an "alias" to the first node.
type Tree struct {
	Node
}

// Add a value (node) to the tree.
func (t *Tree) Add(v int) {
	fmt.Printf("Adding %d to tree.\n", v)

	if t.Value <= v {
		if t.Left == nil {
			t.Left = &Node{Value: v}
		} else {
			t.Left.Add(v)
		}
	} else {
		if t.Right == nil {
			t.Right = &Node{Value: v}
		} else {
			t.Right.Add(v)
		}
	}
}

// Count returns the number of nodes, including the top node, in a tree.
func (t Tree) Count() int {
	leftCount := 0
	if t.Left != nil {
		leftCount = t.Left.Count()
	}
	rightCount := 0
	if t.Right != nil {
		rightCount = t.Right.Count()
	}
	return 1 + leftCount + rightCount
}

// Total returns the of all of the tree's nodes' values.
func (t Tree) Total() int {
	leftTotal := 0
	if t.Left != nil {
		leftTotal = t.Left.Total()
	}
	rightTotal := 0
	if t.Right != nil {
		rightTotal = t.Right.Total()
	}
	return t.Value + leftTotal + rightTotal
}

func (t Tree) Rows() int {
	leftRow, rightRow := 0, 0
	if t.Left != nil {
		leftRow = t.Left.Rows()
	}
	if t.Right != nil {
		rightRow = t.Right.Rows()
	}
	return greatest(leftRow, rightRow) + 1
}

func (t Tree) Cols() int {
	leftCol, rightCol := 0, 0
	if t.Left != nil {
		leftCol = t.Left.LeftCol()
	}
	if t.Right != nil {
		rightCol = t.Right.RightCol()
	}
	return leftCol + rightCol + 1
}

func (t Tree) String() string {
	var bldr strings.Builder

	// Write cols header.
	bldr.WriteString("  ")
	for idx := 0; idx < cols; idx++ {
		bldr.WriteString(" ")
		bldr.WriteString(strconv.Itoa(idx))
	}
	bldr.WriteString("\n")

	// var leftNodes, rightNodes string
	// if t.Left != nil {
	// 	leftNodes = t.Left.String()
	// }
	// if t.Right != nil {
	// 	rightNodes = t.Right.String()
	// }
	// return leftNodes + " " + strconv.Itoa(t.Value) + " " + rightNodes

	return bldr.String()
}
