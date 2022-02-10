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
			for _, spec := range d.Specs {
				switch s := spec.(type) {
				case *ast.ValueSpec:
					for _, v := range s.Values {
						switch fun := v.(type) {
						case *ast.CallExpr:
							if fun.Fun.(*ast.Ident).Name == "Resource" {
								firstArg := strings.Replace(fun.Args[0].(*ast.BasicLit).Value, "\"", "", -1)
								if !strutils.IsSnakeCase(firstArg) {
									fmt.Fprintf(os.Stderr, "%s:%d '%s' is not snake case ('%s')\n",
										filepath,
										fset.Position(fun.Fun.(*ast.Ident).NamePos).Line,
										firstArg,
										strutils.ToSnakeCase(firstArg))
								}
							}
						}
					}
				}
			}
		}
	}
}
