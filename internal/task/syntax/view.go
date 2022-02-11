package syntax

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// ViewSyntaxChecker check View() function syntax
func ViewSyntaxChecker(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		switch d := decl.(type) {
		// There are many more other types but we only focus on FuncDecl here.
		case *ast.FuncDecl:
			fmt.Println("function declaration:", d.Name)
			fmt.Println("function declaration:", d.Body.List)
		}
	}
}
