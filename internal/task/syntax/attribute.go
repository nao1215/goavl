package syntax

import (
	"go/parser"
	"go/token"

	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// AttributeSyntaxChecker check Attribute() function syntax
func AttributeSyntaxChecker(filepath, inspectionID string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		syntaxCheck(filepath, inspectionID, fset, decl, "Attribute", []string{
			"View", "Type", "Attribute", "Attributes",
			"MediaType", // ドキュメントには書いていないが使える
		})
	}
}
