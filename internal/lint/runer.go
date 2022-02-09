package lint

import (
	"log"
	"os"

	"github.com/nao1215/goavl/internal/utils"
)

type check func(filepath string)

// Task define one of the perspectives that Linter checks
type Task struct {
	// Name is linter-task name.
	Name string
	// Check define lint task from one perspective.
	Check check
}

// setup returns a slice that sets the linter task.
func setup() []Task {
	tasks := []Task{}

	tasks = append(tasks, NewViewSyntaxTask())
	tasks = append(tasks, NewNamingTask())
	return tasks
}

// Run execute all linter-tasks.
func Run() {
	tasks := setup()
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
