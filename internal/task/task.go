package task

import (
	"github.com/nao1215/goavl/internal/task/debug"
	"github.com/nao1215/goavl/internal/task/name"
	"github.com/nao1215/goavl/internal/task/syntax"
)

type check func(filepath string)

// Task define one of the perspectives that Linter checks
type Task struct {
	// Name is linter-task name.
	Name string
	// Check define lint task from one perspective.
	Check check
}

// Setup returns a slice that sets the linter task.
func Setup() []Task {
	tasks := []Task{}

	tasks = append(tasks, NewViewSyntaxTask()) // Not implement
	tasks = append(tasks, NewResourceNameCheckerTask())
	tasks = append(tasks, NewActionNameCheckerTask())
	tasks = append(tasks, NewRoutingNameCheckerTask())
	return tasks
}

// TODO:
// I wanted to define the NewXxxTask() function in each task file.
// However, I couldn't solve "import cycle not allowed" easily.
// So, as a workaround, I define the function in this file.

// NewViewSyntaxTask return task that check View() function syntax
func NewViewSyntaxTask() Task {
	task := Task{
		Name:  "View() syntax check",
		Check: syntax.ViewSyntaxChecker,
	}
	return task
}

// NewResourceNameCheckerTask return task that check Resource() argument name.
func NewResourceNameCheckerTask() Task {
	task := Task{
		Name:  "Resource() argument name checker",
		Check: name.ResourceNameChecker,
	}
	return task
}

// NewActionNameCheckerTask return task that check Action() argument name.
func NewActionNameCheckerTask() Task {
	task := Task{
		Name:  "Action() argument name checker",
		Check: name.ActionNameChecker,
	}
	return task
}

// NewRoutingNameCheckerTask return task that check Routing() argument name.
func NewRoutingNameCheckerTask() Task {
	task := Task{
		Name:  "Routing() argument name checker",
		Check: name.RoutingNameChecker,
	}
	return task
}

// NewPrintASTTask return task that print abstract syntax tree of the go file.
func NewPrintASTTask() Task {
	task := Task{
		Name:  "Print ast tree",
		Check: debug.PrintAST,
	}
	return task
}
