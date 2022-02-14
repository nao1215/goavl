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

// AttributeNoExampleChecker check Attribute () for which no example is given
func AttributeNoExampleChecker(filepath, inspectionID string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		checkAttributeNoExample(filepath, inspectionID, fset, decl)
	}
}

func checkAttributeNoExample(filepath, inspectionID string, fset *token.FileSet, decl ast.Decl) {
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
					hasExampleOrNoExample := false
					ast.Inspect(node.X.(*ast.CallExpr), func(n ast.Node) bool {
						switch n := n.(type) {
						case *ast.Ident:
							if n.Name == "Example" {
								hasExampleOrNoExample = true
							} else if n.Name == "NoExample" {
								fmt.Fprintf(os.Stderr,
									"[%s] %s:%-4d NoExample() is not user(client) friendly. you use Example()\n",
									color.YellowString(inspectionID),
									filepath,
									fset.Position(n.NamePos).Line)
								hasExampleOrNoExample = true
							}
						}
						return true
					})
					if !hasExampleOrNoExample {
						fmt.Fprintf(os.Stderr,
							"[%s] %s:%-4d Not exist Example() in Attribute().\n",
							color.YellowString(inspectionID),
							filepath,
							fset.Position(node.X.(*ast.CallExpr).Fun.(*ast.Ident).NamePos).Line)
					}
				}
			}
			return true
		})
	}
}
