package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"reflect"
)

func main() {
	// expr := &ast.BinaryExpr{
	// 	X: &ast.BasicLit{
	// 		Value: "3",
	// 		Kind:  token.INT,
	// 	},
	// 	Op: token.ADD,
	// 	Y: &ast.BasicLit{
	// 		Value: "3",
	// 		Kind:  token.INT,
	// 	},
	// }

	var files []*ast.File
	var data []string
	fset := token.NewFileSet()
	fmt.Println(os.Args)
	for _, goFile := range os.Args[1:] {
		fmt.Println(goFile)
		f, err := parser.ParseFile(fset, goFile, nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}
		files = append(files, f)

		file, err := os.ReadFile(goFile)
		if err == nil {
			data = append(data, string(file))
		}
	}

	fmt.Println(files)
	for i, file := range files {
		ast.Inspect(file, func(n ast.Node) bool {
			fmt.Println("------------")
			if n != nil {
				println(n.Pos())
				println("position char:\n", string(data[i][n.Pos() - 1: n.End()]))
			}

			switch a := n.(type) {
			default:
				fmt.Println(reflect.TypeOf(a))
				
			}

			be, ok := n.(*ast.BinaryExpr)
			if !ok {
				return true
			}

			if be.Op != token.ADD {
				return true
			}

			if _, ok := be.X.(*ast.BasicLit); !ok {
				return true
			}

			if _, ok := be.Y.(*ast.BasicLit); !ok {
				return true
			}

			posn := fset.Position(be.Pos())
			fmt.Println("integer addition found: %v", posn)

			return true
		})
	}
}

func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
