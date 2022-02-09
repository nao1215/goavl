package lint

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

// NewViewSyntaxTask return task that check View() function syntax
func NewViewSyntaxTask() Task {
	task := Task{
		Name:  "View() syntax check",
		Check: ViewSyntaxChecker,
	}
	return task
}

// ViewSyntaxChecker check View() function syntax
func ViewSyntaxChecker(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, decl := range f.Decls {
		switch d := decl.(type) {
		// There are many more other types but we only focus on FuncDecl here.
		case *ast.FuncDecl:
			fmt.Println("function declaration:", d.Name)
			fmt.Println("function declaration:", d.Body.List)
		}
	}
	ast.Print(fset, f)
}
