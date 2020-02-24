package bst

import (
	"fmt"
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
	fmt.Printf("Adding %d to node with value %d.\n", v, t.Value)

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

func (n Node) string(str *string) string {
	row := fmt.Sprintf("%s%d", withSpaces(n.Rows()-1), n.Value)
	*str += row
	var nextRow string
	if n.Left != nil {
		n.Left.string(&nextRow)
		fmt.Printf("(n) left returned with value %d\n", n.Left.Value)
	}
	if n.Right != nil {
		n.Right.string(&nextRow)
		fmt.Printf("(n) right returned with value %d -- %s\n", n.Right.Value, nextRow)
	}
	fmt.Printf("looking at %s\n", *str+nextRow)
	return *str + "\n" + nextRow
}

// func (t Node) String() string {
// 	var leftNodes, rightNodes string
// 	if t.Left != nil {
// 		leftNodes = t.Left.String()
// 	}
// 	if t.Right != nil {
// 		rightNodes = t.Right.String()
// 	}

// 	return leftNodes + " " + strconv.Itoa(t.Value) + " " + rightNodes
// }

func (n Node) printNode(count int) {
	// Print number of space for this branch.
	for idx := 0; idx < count; idx++ {
		fmt.Printf(" ")
	}
	fmt.Printf("node %d: ", n.Value)
	fmt.Printf("%d ", n.Value)
	fmt.Printf("\n")
	fmt.Printf("left:\n")
	if n.Left != nil {
		n.Left.printNode(count + 1)
	}
	fmt.Printf("right:\n")
	if n.Right != nil {
		n.Right.printNode(count + 1)
	}
}

func (n Node) String() string {
	n.printNode(1)
	return ""
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

// matrixRows converts nodes following a head node to a matrix.
// It does this by taking the given rows and row number, and placing the
// current node value in the given row. It then follows a left-leaning DFS
// to create the following rows.
func (n Node) matrixRows(rows [][]int, row int) {
	rows[row] = append(rows[row], n.Value)
	if n.Left != nil {
		n.Left.matrixRows(rows, row+1)
	}
	if n.Right != nil {
		n.Right.matrixRows(rows, row+1)
	}
}

// description returns a verbose, human readable description of a node.
func (n Node) description(parentVal int) description {
	desc := description{
		numChildren: 0,
		parentVal:   parentVal,
	}

	if n.Left != nil {
		desc.numChildren += 1
		desc.leftNodeVal = pInt(n.Left.Value)
	}
	if n.Right != nil {
		desc.numChildren += 1
		desc.rightNodeVal = pInt(n.Right.Value)
	}
	return desc
}

func (n Node) ascii(parentVal int) string {
	str := fmt.Sprintf("  %d\n", parentVal)
	if n.Left != nil {
		if n.Right != nil {
			str += fmt.Sprintf(" / \\\n")
			str += fmt.Sprintf("%d   %d\n", n.Left.Value, n.Right.Value)
		} else {
			str += fmt.Sprintf(" /\n")
			str += fmt.Sprintf("%d\n", n.Left.Value)
		}
	} else if n.Right != nil {
		str += fmt.Sprintf("   \\\n")
		str += fmt.Sprintf("    %d", n.Right.Value)
	}
	return str
}

func (d description) String() string {
	str := fmt.Sprintf("The node has %d children and its parent value is %d.\n", d.numChildren, d.parentVal)
	if d.leftNodeVal != nil {
		str += fmt.Sprintf("Its left node value is %d.\n", *d.leftNodeVal)
	} else {
		str += fmt.Sprintf("It does not have a left node.\n")
	}

	if d.rightNodeVal != nil {
		str += fmt.Sprintf("Its right node value is %d.\n", *d.rightNodeVal)
	} else {
		str += fmt.Sprintf("It does not have a right node.\n")
	}
	return str
}

// Describe prints a verbose human-readable description of a tree to stdout.
func (n Node) Describe(position string) {
	fmt.Printf("Describing a %s node:\n", position)
	fmt.Printf("The node has the value %d.\n", n.Value)
	fmt.Printf("%s\n", n.description(n.Value))
	fmt.Printf("%s\n", n.ascii(n.Value))
	if n.Left != nil {
		n.Left.Describe("left")
	}
	if n.Right != nil {
		n.Right.Describe("right")
	}
}
