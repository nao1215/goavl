package syntax

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/fatih/color"
	"github.com/nao1215/goavl/internal/utils/goautils"
	"github.com/nao1215/goavl/internal/utils/ioutils"
	"github.com/nao1215/goavl/internal/utils/strutils"
)

// DefaultSyntaxChecker check Default() function syntax
func DefaultSyntaxChecker(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		checkDefault(filepath, fset, decl)
	}
}

func checkDefault(filepath string, fset *token.FileSet, decl ast.Decl) {
	okFuncs := []string{
		"Attribute", "Default",
	}

	functions := goautils.CheckTargetFunctionList()
	for _, v := range okFuncs {
		functions = strutils.Remove(functions, v)
	}

	switch d := decl.(type) {
	case *ast.GenDecl:
		ast.Inspect(d, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.CallExpr:
				for _, function := range functions {
					if node.Fun.(*ast.Ident).Name == function {
						if goautils.NotWarnSyntaxCheck(node.Args, function) {
							return true
						}

						ast.Inspect(node, func(n ast.Node) bool {
							switch n := n.(type) {
							case *ast.Ident:
								if n.Name == "Default" {
									fmt.Fprintf(os.Stderr,
										"[%s] %s:%-4d %s() has Default(). Default() can be used in Attribute()\n",
										color.YellowString("WARN"),
										filepath,
										fset.Position(n.NamePos).Line,
										function)
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
}
