package utils

import (
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
