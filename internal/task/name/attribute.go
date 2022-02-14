package name

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/nao1215/goavl/internal/utils/ioutils"
	"github.com/nao1215/goavl/internal/utils/strutils"
)

// AttributeNameChecker check argument name.
func AttributeNameChecker(filepath, inspectionID string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		checkAttributeArgName(filepath, inspectionID, fset, decl)
	}
}

func checkAttributeArgName(filepath, inspectionID string, fset *token.FileSet, decl ast.Decl) {
	switch d := decl.(type) {
	case *ast.GenDecl:
		ast.Inspect(d, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.ExprStmt:
				if node.X.(*ast.CallExpr).Fun.(*ast.Ident).Name == "Attribute" {
					for _, arg := range node.X.(*ast.CallExpr).Args {
						switch bl := arg.(type) {
						case *ast.BasicLit:
							firstArg := strings.Replace(bl.Value, "\"", "", -1)
							if !strutils.IsSnakeCase(firstArg) {
								fmt.Fprintf(os.Stderr,
									"[%s] %s:%-4d Attribute(\"%s\") is not snake case ('%s')\n",
									color.YellowString(inspectionID),
									filepath,
									fset.Position(node.X.(*ast.CallExpr).Fun.(*ast.Ident).NamePos).Line,
									firstArg,
									strutils.ToSnakeCase(firstArg))
							}
						}
					}
				}
			}
			return true
		})
	}
}
