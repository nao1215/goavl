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

// ResourceNameChecker check variable name and argument name.
func ResourceNameChecker(filepath, inspectionID string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}

	for _, decl := range f.Decls {
		checkResourceArgName(filepath, inspectionID, fset, decl)
	}
}

func checkResourceArgName(filepath, inspectionID string, fset *token.FileSet, decl ast.Decl) {
	switch d := decl.(type) {
	case *ast.GenDecl:
		ast.Inspect(d, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.CallExpr:
				if node.Fun.(*ast.Ident).Name == "Resource" {
					for _, bl := range node.Args {
						switch bl := bl.(type) {
						case *ast.BasicLit:
							firstArg := strings.Replace(bl.Value, "\"", "", -1)
							if !strutils.IsSnakeCase(firstArg) {
								fmt.Fprintf(os.Stderr,
									"[%s] %s:%-4d Resource(\"%s\") is not snake case ('%s')\n",
									color.YellowString(inspectionID),
									filepath,
									fset.Position(node.Fun.(*ast.Ident).NamePos).Line,
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
