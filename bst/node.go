package bst

import (
	"fmt"
	"strconv"
)

// Node contains a left and a right node.
// Either can be nil if it does not exist.
type Node struct {
	Left  *Node
	Right *Node
	Value int
}

// Add a value (node) to the node.
func (t *Node) Add(v int) {
	fmt.Printf("Adding %d to node.\n", v)

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

func (n Node) Count() int {
	leftCount := 0
	if n.Left != nil {
		leftCount = n.Left.Count()
	}
	rightCount := 0
	if n.Right != nil {
		rightCount = n.Right.Count()
	}
	return 1 + leftCount + rightCount
}

// Total the of all of the nodes' values.
func (t Node) Total() int {
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

func (n Node) Rows() int {
	leftRow, rightRow := 0, 0
	if n.Left != nil {
		leftRow = n.Left.Rows()
	}
	if n.Right != nil {
		rightRow = n.Right.Rows()
	}
	return greatest(leftRow, rightRow) + 1
}

func (t Node) String() string {
	var leftNodes, rightNodes string
	if t.Left != nil {
		leftNodes = t.Left.String()
	}
	if t.Right != nil {
		rightNodes = t.Right.String()
	}

	return leftNodes + " " + strconv.Itoa(t.Value) + " " + rightNodes
}

func (n Node) LeftCol() int {
	if n.Left != nil {
		return 1 + n.Left.LeftCol()
	}
	return 0
}

func (n Node) RightCol() int {
	if n.Right != nil {
		return 1 + n.Right.RightCol()
	}
	return 0
}
