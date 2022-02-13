package syntax

import (
	"go/parser"
	"go/token"

	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// MultipartFormSyntaxChecker check MultipartForm() function syntax
func MultipartFormSyntaxChecker(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		syntaxCheck(filepath, fset, decl, "MultipartForm", []string{"Action"})
	}
}
