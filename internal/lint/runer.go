package lint

import (
	"github.com/nao1215/goavl/internal/task"
	"github.com/nao1215/goavl/internal/utils/fileutils"
	"github.com/nao1215/goavl/internal/utils/ioutils"
	"github.com/nao1215/goavl/internal/utils/pathutils"
)

// Run execute all linter-tasks.
func Run() {
	tasks := task.Setup()
	files, err := fileutils.Walk(pathutils.CWD())
	if err != nil {
		ioutils.Die(err.Error())
	}

	for _, f := range fileutils.ExtractDesignPackageFile(files) {
		for _, v := range tasks {
			v.Check(f)
		}
	}
}

// PrintAST print abstract syntax tree of the go file.
func PrintAST(files []string) {
	if len(files) == 0 {
		ioutils.Die("specified no go file")
	}

	files = fileutils.ExtractGoFile(files)
	if len(files) == 0 {
		ioutils.Die("you can only specify go file")
	}

	task := task.NewPrintASTTask()
	for _, f := range files {
		if !fileutils.IsFile(f) {
			ioutils.Warn("no such file or directory exists: " + f)
			continue
		}
		task.Check(f)
	}
}
