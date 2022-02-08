package lint

import (
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
	ast.Print(fset, f)
}
