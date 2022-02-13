package syntax

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/nao1215/goavl/internal/utils/goautils"
	"github.com/nao1215/goavl/internal/utils/strutils"
)

func syntaxCheck(filepath string, fset *token.FileSet, decl ast.Decl, targetFunc string, okFuncs []string) {
	functions := goautils.CheckTargetFunctionList()
	for _, v := range okFuncs {
		functions = strutils.Remove(functions, v)
	}

	result := map[int]string{}
	switch d := decl.(type) {
	case *ast.GenDecl:
		ast.Inspect(d, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.CallExpr:
				for _, function := range functions {
					if node.Fun.(*ast.Ident).Name == function {
						if goautils.NotWarnSyntaxCheck(node.Args, function) || function == targetFunc {
							return true
						}

						ast.Inspect(node, func(n ast.Node) bool {
							switch n := n.(type) {
							case *ast.Ident:
								if n.Name == targetFunc {
									result[fset.Position(n.NamePos).Line] = fmt.Sprintf(
										"[%s] %s:%-4d %s() has %s(). %s() can be used in %s\n",
										color.YellowString("WARN"),
										filepath,
										fset.Position(n.NamePos).Line,
										function,
										targetFunc,
										targetFunc,
										strings.Join(okFuncs, ", "),
									)
								}
							}
							return true
						})
					}
				}
			}
			return true
		})
	}
	for _, v := range result {
		fmt.Fprint(os.Stderr, v)
	}
}
