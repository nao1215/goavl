package fileutils

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// IsFile reports whether the path exists and is a file.
func IsFile(path string) bool {
	stat, err := os.Stat(path)
	return (err == nil) && (!stat.IsDir())
}

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

// ExtractGoFile extract go file in filepath list.
func ExtractGoFile(files []string) []string {
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
	files = ExtractGoFile(files)
	extractFiles := []string{}

	for _, filepath := range files {
		if IsDesignFile(filepath) {
			extractFiles = append(extractFiles, filepath)
		}
	}
	return extractFiles
}

// IsDesignFile return whether file is goa-design file.
func IsDesignFile(file string) bool {
	if !IsFile(file) {
		return false
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, 0)
	if err != nil {
		return false
	}

	// design package or not
	if f.Name.Name != "design" {
		return false
	}

	// TODO:
	// I want to judge whether file is goa-design file or not by the import path.
	// However, import path is renamed original path to "." in goa-design file.
	// Therefore, it is not possible to determine whether it is a goa-design file
	// with the correct import path ("github.com/shogo82148/goa-v1/design").
	for _, v := range f.Imports {
		if v.Name.Name == "." {
			return true
		}
	}
	return false
}
