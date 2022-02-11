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

// AttributeSyntaxChecker check View() function syntax
func AttributeSyntaxChecker(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		checkAttribute(filepath, fset, decl)
	}
}

func checkAttribute(filepath string, fset *token.FileSet, decl ast.Decl) {
	// Attribute can be used in: View, Type, Attribute, Attributes
	ngFuncs := []string{
		"API",
		"Resource",
		"Action",
		//"MediaType", ドキュメントには書いていないが使える
	}

	switch d := decl.(type) {
	case *ast.GenDecl:
		ast.Inspect(d, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.CallExpr:
				for _, function := range ngFuncs {
					if node.Fun.(*ast.Ident).Name == function {
						if len(node.Args) == 1 && function != "Attributes" {
							// If function has only one Argument, e.g. Attribute("test")
							return true
						}

						ast.Inspect(node, func(n ast.Node) bool {
							switch n := n.(type) {
							case *ast.Ident:
								if n.Name == "Attribute" {
									fmt.Fprintf(os.Stderr,
										"[%s] %s:%-4d %s() has Attribute(). Attribute() can be used in View(), Type(), Attribute(), Attributes(), MediaType()\n",
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
