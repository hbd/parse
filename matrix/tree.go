package matrix

import "fmt"

// PrintTree prints a matrix as a tree structure.
// For example, if the following matrix is given:
// 0: 2
// 1: 1 3
// The following tree is printed:
//   2
// 1   3
// This might be impossible to do correctly, since the matrix has no understanding
// the relationships between the numbers. It can't determine which node to print
// a value under.
func (mtx Matrix) PrintTree() {
	dist := len(mtx.Rows)
	between := len(mtx.Rows)
	for _, row := range mtx.Rows {
		for _, val := range row {
			fmt.Printf("%s", numSpaces(dist))
			fmt.Printf("%d", val)
			fmt.Printf("%s", numSpaces(between))
		}
		fmt.Printf("\n")
		dist /= 2
	}

	// Reverse.
	// dist := len(mtx.Rows)
	// for idx := len(mtx.Rows) - 1; idx >= 0; idx-- {
	// 	for _, val := range mtx.Rows[idx] {
	// 		fmt.Printf("%s", numSpaces(dist))
	// 		fmt.Printf("%d", val)
	// 	}
	// 	dist /= 2
	// 	fmt.Printf("\n")
	// }
}

// numSpaces returns a string containing the given number of space.
func numSpaces(num int) string {
	str := ""
	for num > 0 {
		str += " "
		num--
	}
	return str
}

// How to print a matrix as a tree:
// The top node must be in the middle of the bottom nodes.
// To determine where to print the top node,
// divide the total number of rows by half.
// Each following row follows a similar pattern
// by remaining half the distance of the previous row.
//
// 0: 2
// 1: 1 3
// 2: 100

// Rows: 3
//    2             3 spaces from front, 3 from end
//  1   3           1 spaces from front, 1 from end
//        100       0 spaces from front, 0 from end
//

// Spaces between
//
