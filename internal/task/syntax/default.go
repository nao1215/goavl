package syntax

import (
	"go/parser"
	"go/token"

	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// DefaultSyntaxChecker check Default() function syntax
func DefaultSyntaxChecker(filepath, inspectionID string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		// The documentation states that "Default" cannot be set in "Param()". However, this explanation is incorrect.
		syntaxCheck(filepath, inspectionID, fset, decl, "Default", []string{"Attribute", "Param"})
	}
}
