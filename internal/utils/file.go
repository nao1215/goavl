package utils

import (
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Walk returns  files under the certain directory (target directory)
func Walk(target string) ([]string, error) {
	var files []string
	err := filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// extractGoFile extract go file in filepath list.
func extractGoFile(files []string) []string {
	f := []string{}
	for _, v := range files {
		if strings.HasSuffix(v, ".go") {
			f = append(f, v)
		}
	}
	return f
}

// ExtractDesignPackageFile extract goa-design package.
func ExtractDesignPackageFile(files []string) []string {
	files = extractGoFile(files)
	extractFiles := []string{}

	for _, filepath := range files {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, filepath, nil, 0)
		if err != nil {
			log.Fatal(err)
		}

		// design package or not
		if f.Name.Name != "design" {
			continue
		}

		// TODO:
		// I want to judge whether file is goa-design file or not by the import path.
		// However, import path is renamed original path to "." in goa-design file.
		// Therefore, it is not possible to determine whether it is a goa-design file
		// with the correct import path ("github.com/shogo82148/goa-v1/design").
		for _, v := range f.Imports {
			if v.Name.Name == "." {
				extractFiles = append(extractFiles, filepath)
			}
		}
	}
	return extractFiles
}
