package syntax

import (
	"go/parser"
	"go/token"

	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// ReadOnlySyntaxChecker check ReadOnly() function syntax
func ReadOnlySyntaxChecker(filepath, inspectionID string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		syntaxCheck(filepath, inspectionID, fset, decl, "ReadOnly",
			[]string{"Attribute"})
	}
}
