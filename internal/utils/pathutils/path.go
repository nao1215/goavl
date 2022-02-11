package pathutils

import (
	"os"

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
