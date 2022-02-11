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
		checkRoutingArgName(filepath, fset, decl)
	}
}

func checkRoutingArgName(filepath string, fset *token.FileSet, decl ast.Decl) {
	switch d := decl.(type) {
	case *ast.GenDecl:
		ast.Inspect(d, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.CallExpr:
				if node.Fun.(*ast.Ident).Name == "Routing" {
					for _, arg := range node.Args {
						switch cl := arg.(type) {
						case *ast.CallExpr:
							for _, bl := range cl.Args {
								switch bl := bl.(type) {
								case *ast.BasicLit:
									firstArg := strings.Replace(bl.Value, "\"", "", -1)
									if firstArg != "" && !strutils.IsChainCaseForRouting(firstArg) {
										fmt.Fprintf(os.Stderr,
											"[%s] %s:%-4d Routing(%s(\"%s\")) is not chain case ('%s')\n",
											color.YellowString("WARN"),
											filepath,
											fset.Position(node.Fun.(*ast.Ident).NamePos).Line,
											cl.Fun.(*ast.Ident).Name,
											firstArg,
											strutils.ToChainCaseForRouting(firstArg))
									}
								}
							}
						}
					}
				}
			}
			return true
		})
	}
}
