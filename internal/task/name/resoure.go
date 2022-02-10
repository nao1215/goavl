package name

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/nao1215/goavl/internal/utils/strutils"
)

// ResourceNameChecker check variable name and argument name.
func ResourceNameChecker(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			ast.Inspect(d, func(node ast.Node) bool {
				switch node := node.(type) {
				case *ast.CallExpr:
					if node.Fun.(*ast.Ident).Name == "Resource" {
						firstArg := strings.Replace(node.Args[0].(*ast.BasicLit).Value, "\"", "", -1)
						if !strutils.IsSnakeCase(firstArg) {
							fmt.Fprintf(os.Stderr, "%s:%d Resource(\"%s\") is not snake case ('%s')\n",
								filepath,
								fset.Position(node.Fun.(*ast.Ident).NamePos).Line,
								firstArg,
								strutils.ToSnakeCase(firstArg))
						}
					}
				}
				return true
			})
		}
	}
}
