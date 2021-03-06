package lint

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/nao1215/goavl/internal/task"
	"github.com/nao1215/goavl/internal/utils/fileutils"
	"github.com/nao1215/goavl/internal/utils/ioutils"
	"github.com/nao1215/goavl/internal/utils/pathutils"
)

// Run execute all linter-tasks.
func Run(files, excludeIDs []string) {
	tasks := task.Setup()
	files, err := fileutils.Walk(pathutils.CWD())
	if err != nil {
		ioutils.Die(err.Error())
	}

	for _, f := range fileutils.ExtractDesignPackageFile(files) {
		f = pathutils.RemoveCWDPath(f)
		for _, v := range tasks {
			v.Check(f, v.InspectionID)
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
		task.Check(f, task.InspectionID)
	}
}

// PrintCheckTaskList print task list.
func PrintCheckTaskList() {
	tasks := task.Setup()
	for _, v := range tasks {
		fmt.Printf("%s: %s\n", color.GreenString(v.InspectionID), v.Name)
	}
}

// CheckOneFile check file that user specify.
func CheckOneFile(file string) {
	if !fileutils.IsFile(file) {
		ioutils.Die("no such file or directory exists: " + file)
	}

	if !fileutils.IsDesignFile(file) {
		ioutils.Die("this file is not goa-design file: " + file)
	}

	tasks := task.Setup()
	for _, v := range tasks {
		v.Check(file, v.InspectionID)
	}
}
