package matrix

import (
	"strconv"
	"strings"
)

type Matrix [][]int

func (matrix Matrix) String() string {
	var bldr strings.Builder

	// Number of cols = length of longest row.
	cols := 0
	for _, row := range matrix {
		cols = greatest(cols, len(row))
	}

	// Write cols header.
	bldr.WriteString("  ")
	for idx := 0; idx < cols; idx++ {
		bldr.WriteString(" ")
		bldr.WriteString(strconv.Itoa(idx))
	}
	bldr.WriteString("\n")

	bldr.WriteString("   ")
	for idx := 0; idx < cols+int(cols/2); idx++ {
		bldr.WriteString("-")
	}
	bldr.WriteString("-\n")

	// Write the rows.
	for rowidx, row := range matrix {
		bldr.WriteString(strconv.Itoa(rowidx))
		bldr.WriteString(":")
		for _, col := range row {
			bldr.WriteString(" ")
			bldr.WriteString(strconv.Itoa(col))
		}
		bldr.WriteString("\n")
	}
	return bldr.String()
}

func greatest(a, b int) int {
	if a > b {
		return a
	}
	return b
}
