package syntax

import (
	"go/parser"
	"go/token"

	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// UseTraitSyntaxChecker check UseTrait() function syntax
func UseTraitSyntaxChecker(filepath, inspectionID string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		syntaxCheck(filepath, inspectionID, fset, decl, "UseTrait",
			[]string{"Resource", "Action", "Type", "MediaType", "Attribute"})
	}
}
