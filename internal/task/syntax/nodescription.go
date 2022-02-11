package syntax

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/fatih/color"
	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// AttributeNoDescriptionChecker check Attribute () for which no example is given
func AttributeNoDescriptionChecker(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		checkAttributeNoDescription(filepath, fset, decl)
	}
}

func checkAttributeNoDescription(filepath string, fset *token.FileSet, decl ast.Decl) {
	switch d := decl.(type) {
	case *ast.GenDecl:
		ast.Inspect(d, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.ExprStmt:
				if node.X.(*ast.CallExpr).Fun.(*ast.Ident).Name == "Attribute" {
					if len(node.X.(*ast.CallExpr).Args) == 1 {
						// If Attribute() has only one Argument, e.g. Attribute("test")
						return true
					}
					hasDescription := false
					ast.Inspect(node.X.(*ast.CallExpr), func(n ast.Node) bool {
						switch n := n.(type) {
						case *ast.Ident:
							if n.Name == "Description" {
								hasDescription = true
							}
						}
						return true
					})
					if !hasDescription {
						fmt.Fprintf(os.Stderr,
							"[%s] %s:%-4d Not exist Description() in Attribute().\n",
							color.YellowString("WARN"),
							filepath,
							fset.Position(node.X.(*ast.CallExpr).Fun.(*ast.Ident).NamePos).Line)
					}
				}
			}
			return true
		})
	}
}
