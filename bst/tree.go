// Package bst implements a binary search tree.

package bst

import (
	"fmt"
	"parse/matrix"
)

// Tree represents a BST, starting at the top node.
// It's an "alias" to the first node.
type Tree struct {
	Node
}

// Add a value (node) to the tree.
func (t *Tree) Add(v int) {
	fmt.Printf("Adding %d to tree with root %d.\n", v, t.Value)

	if t.Value > v {
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

// Rows returns the number of "rows" in the tree.
// For example, if the BST looks like:
//    3
//  1   2
// Then it has 2 rows.
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

func (t Tree) Print() {
	fmt.Printf("top:\n%d\n", t.Value)
	if t.Left != nil {
		t.Left.printNode(1)
	}
	if t.Right != nil {
		t.Right.printNode(1)
	}
}

func (t Tree) String() string {
	top := fmt.Sprintf("%s%d\n", withSpaces(t.Rows()-1), t.Value)
	var nextRow string
	if t.Left != nil {
		t.Left.string(&nextRow)
		fmt.Printf("left returned with value %d\n", t.Left.Value)
	}
	if t.Right != nil {
		t.Right.string(&nextRow)
		fmt.Printf("right returned with value %d\n", t.Right.Value)
	}
	return top + nextRow
}

// ToMatrix converts the BST to a matrix.
// It uses BFS to iterate across the tree and store nodes at their level (row).
// For example, the top is stored at level 0 (row 0), and the left and right nodes of the top are stored at level 1.
// If the BST looks like:
//    3
//  1   2
// The matrix looks like:
// Level 0: 3
// Level 1: 1, 2
func (t Tree) ToMatrix() matrix.Matrix {
	// Initialize the matrix with known dimensions.
	rows := t.Rows()
	mtx := matrix.Matrix{}
	mtx.Rows = make([][]int, rows)

	// Initialize the first row.
	mtx.Rows[0] = make([]int, 1)
	mtx.Rows[0][0] = t.Value

	// Create the following rows. Start on second row by passing 1.
	if t.Left != nil {
		t.Left.matrixRows(mtx.Rows, 1)
	}
	if t.Right != nil {
		t.Right.matrixRows(mtx.Rows, 1)
	}

	return mtx
}

// Describe prints a verbose human-readable description of a tree to stdout.
func (t Tree) Describe() {
	fmt.Println("Describing a tree:")
	fmt.Printf("The root node has the value %d.\n", t.Value)
	fmt.Printf("%s\n", t.Node.description(t.Value))
	fmt.Printf("%s\n", t.Node.ascii(t.Value))
	if t.Left != nil {
		t.Left.Describe("left")
	}
	if t.Right != nil {
		t.Right.Describe("right")
	}
}

/*
	fmt.Printf("left:\n")
	if t.Left != nil {
		a := t.Left.String()
		_ = a
	}
	fmt.Printf("%d ", t.Value)
	fmt.Printf("right:\n")
	if t.Right != nil {
		a := t.Right.String()
		_ = a
	}

	return ""


	// rows, cols := t.Rows(), t.Cols()

	// m := t.ToMatrix()

	// var leftNodes, rightNodes string
	// if t.Left != nil {
	// 	leftNodes = t.Left.String()
	// }
	// if t.Right != nil {
	// 	rightNodes = t.Right.String()
	// }
	// return leftNodes + " " + strconv.Itoa(t.Value) + " " + rightNodes

	// return m.String()

*/
