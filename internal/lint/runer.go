package lint

type check func()

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

	tasks = append(tasks, NewSyntaxViewTask())
	return tasks
}

// Run execute all linter-tasks.
func Run() {
	tasks := setup()
	for _, v := range tasks {
		v.Check()
	}
}
