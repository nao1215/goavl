package lint

import (
	"log"
	"os"

	"github.com/nao1215/goavl/internal/task"
	"github.com/nao1215/goavl/internal/utils"
)

// Run execute all linter-tasks.
func Run() {
	tasks := task.Setup()
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := utils.Walk(cwd)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range tasks {
		for _, f := range utils.ExtractDesignPackageFile(files) {
			v.Check(f)
		}
	}
}
