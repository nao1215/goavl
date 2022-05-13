package syntax

import (
	"go/parser"
	"go/token"

	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// RequiredSyntaxChecker check Required() function syntax
func RequiredSyntaxChecker(filepath, inspectionID string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		// The documentation says "Required can be used in: Attributes, Headers, Payload, Type, Params".
		// However, goa can use Required within MediaType. If you use Required within MediaType,
		// the generated files will also change
		syntaxCheck(filepath, inspectionID, fset, decl, "Required",
			[]string{"Attributes", "Headers", "Payload", "Type", "Params", "MediaType"})
	}
}
