package pathutils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/nao1215/goavl/internal/utils/ioutils"
)

// CWD return current working directory.
func CWD() string {
	cwd, err := os.Getwd()
	if err != nil {
		ioutils.Die(err.Error())
	}
	return cwd
}

// RemoveCWDPath remove CWD from path.
func RemoveCWDPath(path string) string {
	return strings.Replace(path, CWD()+string(filepath.Separator), "", 1)
}
