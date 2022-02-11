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

// ActionNameChecker check variable name and argument name.
func ActionNameChecker(filepath string) {
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
					if node.Fun.(*ast.Ident).Name == "Action" {
						for _, arg := range node.Args {
							switch bl := arg.(type) {
							case *ast.BasicLit:
								firstArg := strings.Replace(bl.Value, "\"", "", -1)
								if !strutils.IsSnakeCase(firstArg) {
									fmt.Fprintf(os.Stderr,
										"[%s] %s:%-4d Action(\"%s\") is not snake case ('%s')\n",
										color.YellowString("WARN"),
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
}
