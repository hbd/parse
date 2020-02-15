// Parse parses go files.

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"parse/bst"

	"github.com/sirupsen/logrus"
)

func main() {
	var objects = map[string]*ast.Object{}

	// Parse go files.
	// Parse structs from a go file.
	fset := token.NewFileSet()

	tfs, err := parser.ParseDir(fset, "parser", nil, 0)
	if err != nil {
		logrus.WithError(err).Fatalf("Error parsing dir.")
	}
	for pkgName, pkg := range tfs {
		fmt.Printf("package: %s\n**contents**:\n", pkgName)
		if pkg.Scope == nil {
			fmt.Println("No scope.")
		} else {
			for objName, obj := range pkg.Scope.Objects {
				fmt.Printf("obj name: %s\nobj: %v\n", objName, obj)
			}
		}

		if pkg.Files == nil {
			fmt.Println("No files.")
		} else {
			for fname, f := range pkg.Files {
				fmt.Printf("file name: %s\n", fname)
				if f.Scope == nil {
					fmt.Println("File doesn't have a scope.")
				} else {
					for objName, obj := range f.Scope.Objects {
						fmt.Printf("obj name: %s\nobj kind: %s\n\n", objName, obj.Kind)
					}
					objects = f.Scope.Objects
				}
			}
		}
	}

	_ = objects

	left, right := &bst.Node{Value: 1}, &bst.Node{Value: 3}
	searchTree := bst.Tree{
		Node: bst.Node{
			Left:  left,
			Right: right,
			Value: 2,
		},
	}
	searchTree.Add(0)
	stats := searchTree.Stats()

	fmt.Println("**tree**")
	fmt.Printf("num nodes: %d\ntotal: %d\nrows: %d\ncols: %d\n", stats.Count, stats.Total, stats.Rows, stats.Cols)
	fmt.Printf("%s\n", searchTree)
}
