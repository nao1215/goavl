package require

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/fatih/color"
	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// ArrayRequireChecker check variable name and argument name.
func ArrayRequireChecker(filepath, inspectionID string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		ioutils.Die(err.Error())
	}

	for _, decl := range f.Decls {
		checkArrayWithRequire(filepath, inspectionID, fset, decl)
	}
}

type arrayOf struct {
	name     string
	line     int
	required bool
}

func checkArrayWithRequire(filepath, inspectionID string, fset *token.FileSet, decl ast.Decl) {
	// ArrayOfを管理するマップ。
	// 文字列はArrayOfを使っているレスポンス名、boolはRequiredの指定があったか
	results := []arrayOf{}

	switch d := decl.(type) {
	case *ast.GenDecl:
		ast.Inspect(d, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.ExprStmt:
				if node.X.(*ast.CallExpr).Fun.(*ast.Ident).Name == "Attribute" {
					for i, arg := range node.X.(*ast.CallExpr).Args {
						switch bl := arg.(type) {
						case *ast.CallExpr:
							if bl.Fun.(*ast.Ident).Name == "ArrayOf" {
								// 型情報は第2引数であり、第1引数が名称
								name := node.X.(*ast.CallExpr).Args[i-1].(*ast.BasicLit).Value
								results = append(results, arrayOf{
									name:     name,
									line:     fset.Position(node.X.(*ast.CallExpr).Fun.(*ast.Ident).NamePos).Line,
									required: false,
								})
							}
						}
					}
				}
				if node.X.(*ast.CallExpr).Fun.(*ast.Ident).Name == "Required" {
					for _, arg := range node.X.(*ast.CallExpr).Args {
						switch bl := arg.(type) {
						case *ast.BasicLit:
							tmp := []arrayOf{}
							for _, v := range results {
								if v.name == bl.Value {
									v.required = true
								}
								tmp = append(tmp, v)
							}
							results = tmp
						}
					}
				}
			}
			return true
		})
	}

	for _, v := range results {
		if !v.required {
			fmt.Fprintf(os.Stderr,
				"[%s] %s:%-4d %s is Arrayof() without Required(). it is nullable.\n",
				color.YellowString(inspectionID),
				filepath,
				v.line,
				v.name)
		}
	}
}
