package lint

import (
	"log"
	"os"

	"github.com/nao1215/goavl/internal/task"
	"github.com/nao1215/goavl/internal/utils/fileutils"
)

// Run execute all linter-tasks.
func Run() {
	tasks := task.Setup()
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := fileutils.Walk(cwd)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range fileutils.ExtractDesignPackageFile(files) {
		for _, v := range tasks {
			v.Check(f)
		}
	}
}
