package lint

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

// NewNamingTask return task that check variable name and argument name.
func NewNamingTask() Task {
	task := Task{
		Name:  "Variable name and Argument name checker",
		Check: NamingChecker,
	}
	return task
}

// NamingChecker check variable name and argument name.
func NamingChecker(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch s := spec.(type) {
				case *ast.ValueSpec:
					for _, v := range s.Values {
						switch fun := v.(type) {
						case *ast.CallExpr:
							if fun.Fun.(*ast.Ident).Name == "Resource" {
								fmt.Println(fun.Args[0].(*ast.BasicLit).Value)
							}
						}
					}
				}
			}
		}
	}
}
