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

// ViewSyntaxChecker check View() function syntax
func ViewSyntaxChecker(filepath string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}
	for _, decl := range f.Decls {
		checkView(filepath, fset, decl)
	}
}

func checkView(filepath string, fset *token.FileSet, decl ast.Decl) {
	// View can be used in: MediaType, Response
	ngFuncs := []string{
		"API",
		"Resource",
		"Action",
		"Attribute",
		"Attributes",
		"Type",
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

						hasView := false
						ast.Inspect(node, func(n ast.Node) bool {
							switch n := n.(type) {
							case *ast.Ident:
								if n.Name == "View" {
									hasView = true
								}
							}
							return true
						})
						if hasView {
							fmt.Fprintf(os.Stderr,
								"[%s] %s:%-4d %s() has View(). View() can be used in MediaType() or Response()\n",
								color.YellowString("WARN"),
								filepath,
								fset.Position(node.Fun.(*ast.Ident).NamePos).Line,
								function)
						}
					}
				}
			}
			return true
		})
	}
}
