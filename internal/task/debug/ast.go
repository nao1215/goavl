package debug

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// PrintAST print abstract syntax tree of the go file.
func PrintAST(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	ast.Print(fset, f)
}
