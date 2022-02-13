package syntax

import (
	"go/parser"
	"go/token"

	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// MinimumSyntaxChecker check Minimum() function syntax
func MinimumSyntaxChecker(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		syntaxCheck(filepath, fset, decl, "Minimum", []string{
			"Attribute", "Header", "Param", "HashOf", "ArrayOf",
		})
	}
}
