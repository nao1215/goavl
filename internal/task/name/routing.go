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

// RoutingNameChecker check variable name and argument name.
func RoutingNameChecker(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			ast.Inspect(d, func(node ast.Node) bool {
				switch node := node.(type) {
				case *ast.CallExpr:
					if node.Fun.(*ast.Ident).Name == "Routing" {
						firstArg := strings.Replace(node.Args[0].(*ast.CallExpr).Args[0].(*ast.BasicLit).Value,
							"\"", "", -1)
						if !strutils.IsChainCaseForRouting(firstArg) {
							fmt.Fprintf(os.Stderr,
								"[%s] %s:%d Routing(%s(\"%s\")) is not chain case ('%s')\n",
								color.YellowString("WARN"),
								filepath,
								fset.Position(node.Fun.(*ast.Ident).NamePos).Line,
								node.Args[0].(*ast.CallExpr).Fun.(*ast.Ident).Name,
								firstArg,
								strutils.ToChainCaseForRouting(firstArg))
						}
					}
				}
				return true
			})
		}
	}
}
