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

// ActionNameChecker check variable name and argument name.
func ActionNameChecker(filepath string) {
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
							for _, a := range fun.Args {
								switch funcLit := a.(type) {
								case *ast.FuncLit:
									exprStmt := funcLit.Body.List[0].(*ast.ExprStmt) // TODO: Danger
									if exprStmt.X.(*ast.CallExpr).Fun.(*ast.Ident).Name == "Action" {
										for _, arg := range exprStmt.X.(*ast.CallExpr).Args {
											switch bl := arg.(type) {
											case *ast.BasicLit:
												firstArg := strings.Replace(bl.Value, "\"", "", -1)
												if !strutils.IsSnakeCase(firstArg) {
													fmt.Fprintf(os.Stderr, "%s:%d Action('%s') is not snake case ('%s')\n",
														filepath,
														fset.Position(exprStmt.X.(*ast.CallExpr).Fun.(*ast.Ident).NamePos).Line,
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
			}
		}
	}
}
