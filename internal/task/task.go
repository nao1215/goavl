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
	// InspectionID is　task-specific number
	InspectionID string
	// Check define lint task from one perspective.
	Check check
}

// Setup returns a slice that sets the linter task.
func Setup() []Task {
	tasks := []Task{}

	tasks = append(tasks, NewResourceNameCheckerTask())
	tasks = append(tasks, NewActionNameCheckerTask())
	tasks = append(tasks, NewRoutingNameCheckerTask())
	tasks = append(tasks, NewAttributeNameCheckerTask())
	tasks = append(tasks, NewAttributeNoExampleCheckerTask())
	tasks = append(tasks, NewNoDescriptionCheckerTask())
	tasks = append(tasks, NewAttributeSyntaxTask())
	tasks = append(tasks, NewDefaultSyntaxTask())
	tasks = append(tasks, NewEnumSyntaxTask())
	tasks = append(tasks, NewExampleSyntaxTask())
	tasks = append(tasks, NewFormatSyntaxTask())
	tasks = append(tasks, NewHeaderSyntaxTask())
	tasks = append(tasks, NewMaxLengthSyntaxTask())
	tasks = append(tasks, NewMaximumSyntaxTask())
	tasks = append(tasks, NewMinLengthSyntaxTask())
	tasks = append(tasks, NewMinimumSyntaxTask())
	tasks = append(tasks, NewMultipartFormSyntaxTask())
	tasks = append(tasks, NewNoExampleSyntaxCheckerTask())
	tasks = append(tasks, NewParamSyntaxCheckerTask())
	tasks = append(tasks, NewPatternSyntaxCheckerTask())
	tasks = append(tasks, NewReadOnlySyntaxCheckerTask())
	tasks = append(tasks, NewRequiredSyntaxCheckerTask())
	tasks = append(tasks, NewTypeNameSyntaxCheckerTask())
	tasks = append(tasks, NewURLSyntaxCheckerTask())
	tasks = append(tasks, NewViewSyntaxTask())
	return tasks
}

// ExcludeTask exclude task from task list
func ExcludeTask(tasks []Task, excludeIDs []string) []Task {
	tmp := []Task{}

	for _, v := range excludeIDs {
		for _, task := range tasks {
			if task.InspectionID != v {
				tmp = append(tmp, task)
			}
		}
		tasks = tmp
		tmp = []Task{}
	}
	return tasks
}

// TODO:
// I wanted to define the NewXxxTask() function in each task file.
// However, I couldn't solve "import cycle not allowed" easily.
// So, as a workaround, I define the function in this file.

// NewAttributeSyntaxTask return task that check Attribute() function syntax
func NewAttributeSyntaxTask() Task {
	task := Task{
		Name:         "Attribute() syntax check",
		InspectionID: "FC001",
		Check:        syntax.AttributeSyntaxChecker,
	}
	return task
}

// NewDefaultSyntaxTask return task that check Default() function syntax
func NewDefaultSyntaxTask() Task {
	task := Task{
		Name:         "Default() syntax check",
		InspectionID: "FC002",
		Check:        syntax.DefaultSyntaxChecker,
	}
	return task
}

// NewEnumSyntaxTask return task that check Enum() function syntax
func NewEnumSyntaxTask() Task {
	task := Task{
		Name:         "Enum() syntax check",
		InspectionID: "FC003",
		Check:        syntax.EnumSyntaxChecker,
	}
	return task
}

// NewExampleSyntaxTask return task that check Example() function syntax
func NewExampleSyntaxTask() Task {
	task := Task{
		Name:         "Example() syntax check",
		InspectionID: "FC004",
		Check:        syntax.ExampleSyntaxChecker,
	}
	return task
}

// NewFormatSyntaxTask return task that check Format() function syntax
func NewFormatSyntaxTask() Task {
	task := Task{
		Name:         "Format() syntax check",
		InspectionID: "FC005",
		Check:        syntax.FormatSyntaxChecker,
	}
	return task
}

// NewHeaderSyntaxTask return task that check Format() function syntax
func NewHeaderSyntaxTask() Task {
	task := Task{
		Name:         "Header() syntax check",
		InspectionID: "FC006",
		Check:        syntax.HeaderSyntaxChecker,
	}
	return task
}

// NewMaxLengthSyntaxTask return task that check MaxLength() function syntax
func NewMaxLengthSyntaxTask() Task {
	task := Task{
		Name:         "MaxLength() syntax check",
		InspectionID: "FC007",
		Check:        syntax.MaxLengthChecker,
	}
	return task
}

// NewMaximumSyntaxTask return task that check Maximum() function syntax
func NewMaximumSyntaxTask() Task {
	task := Task{
		Name:         "Maximum() syntax check",
		InspectionID: "FC008",
		Check:        syntax.MaximumSyntaxChecker,
	}
	return task
}

// NewMemberSyntaxTask return task that check Maximum() function syntax
func NewMemberSyntaxTask() Task {
	task := Task{
		Name:         "Member() syntax check",
		InspectionID: "FC009",
		Check:        syntax.MemberSyntaxChecker,
	}
	return task
}

// NewMinLengthSyntaxTask return task that check MinLength() function syntax
func NewMinLengthSyntaxTask() Task {
	task := Task{
		Name:         "MinLength() syntax check",
		InspectionID: "FC010",
		Check:        syntax.MinLengthSyntaxChecker,
	}
	return task
}

// NewMinimumSyntaxTask return task that check Minimum() function syntax
func NewMinimumSyntaxTask() Task {
	task := Task{
		Name:         "Minimum() syntax check",
		InspectionID: "FC011",
		Check:        syntax.MinimumSyntaxChecker,
	}
	return task
}

// NewMultipartFormSyntaxTask return task that check MultipartForm() function syntax
func NewMultipartFormSyntaxTask() Task {
	task := Task{
		Name:         "MultipartForm() syntax check",
		InspectionID: "FC012",
		Check:        syntax.MultipartFormSyntaxChecker,
	}
	return task
}

// NewNoExampleSyntaxCheckerTask return task that check NoExample() function syntax
func NewNoExampleSyntaxCheckerTask() Task {
	task := Task{
		Name:         "NoExample() syntax check",
		InspectionID: "FC013",
		Check:        syntax.NoExampleSyntaxChecker,
	}
	return task
}

// NewParamSyntaxCheckerTask return task that check Param() function syntax
func NewParamSyntaxCheckerTask() Task {
	task := Task{
		Name:         "Param() syntax check",
		InspectionID: "FC014",
		Check:        syntax.ParamSyntaxChecker,
	}
	return task
}

// NewPatternSyntaxCheckerTask return task that check Pattern() function syntax
func NewPatternSyntaxCheckerTask() Task {
	task := Task{
		Name:         "Pattern() syntax check",
		InspectionID: "FC015",
		Check:        syntax.PatternSyntaxChecker,
	}
	return task
}

// NewReadOnlySyntaxCheckerTask return task that check ReadOnly() function syntax
func NewReadOnlySyntaxCheckerTask() Task {
	task := Task{
		Name:         "ReadOnly() syntax check",
		InspectionID: "FC016",
		Check:        syntax.ReadOnlySyntaxChecker,
	}
	return task
}

// NewRequiredSyntaxCheckerTask return task that check Required() function syntax
func NewRequiredSyntaxCheckerTask() Task {
	task := Task{
		Name:         "Required() syntax check",
		InspectionID: "FC017",
		Check:        syntax.RequiredSyntaxChecker,
	}
	return task
}

// NewTypeNameSyntaxCheckerTask return task that check TypeName() function syntax
func NewTypeNameSyntaxCheckerTask() Task {
	task := Task{
		Name:         "TypeName() syntax check",
		InspectionID: "FC018",
		Check:        syntax.TypeNameSyntaxChecker,
	}
	return task
}

// NewURLSyntaxCheckerTask return task that check URL() function syntax
func NewURLSyntaxCheckerTask() Task {
	task := Task{
		Name:         "URL() syntax check",
		InspectionID: "FC019",
		Check:        syntax.URLSyntaxChecker,
	}
	return task
}

// NewViewSyntaxTask return task that check View() function syntax
func NewViewSyntaxTask() Task {
	task := Task{
		Name:         "View() syntax check",
		InspectionID: "FC020",
		Check:        syntax.ViewSyntaxChecker,
	}
	return task
}

// NewResourceNameCheckerTask return task that check Resource() argument name.
func NewResourceNameCheckerTask() Task {
	task := Task{
		Name:         "Resource() argument name checker",
		InspectionID: "NC001",
		Check:        name.ResourceNameChecker,
	}
	return task
}

// NewActionNameCheckerTask return task that check Action() argument name.
func NewActionNameCheckerTask() Task {
	task := Task{
		Name:         "Action() argument name checker",
		InspectionID: "NC002",
		Check:        name.ActionNameChecker,
	}
	return task
}

// NewRoutingNameCheckerTask return task that check Routing() argument name.
func NewRoutingNameCheckerTask() Task {
	task := Task{
		Name:         "Routing() argument name checker",
		InspectionID: "NC003",
		Check:        name.RoutingNameChecker,
	}
	return task
}

// NewAttributeNameCheckerTask return task that check Attribute argument name.
func NewAttributeNameCheckerTask() Task {
	task := Task{
		Name:         "Attribute() variable and argument name checker",
		InspectionID: "NC004",
		Check:        name.AttributeNameChecker,
	}
	return task
}

// NewAttributeNoExampleCheckerTask return task that check whether the example of Attribute() is written.
func NewAttributeNoExampleCheckerTask() Task {
	task := Task{
		Name:         "Checker whether the example of Attribute() is written",
		InspectionID: "UF001",
		Check:        syntax.AttributeNoExampleChecker,
	}
	return task
}

// NewNoDescriptionCheckerTask return task that check whether description exist.
func NewNoDescriptionCheckerTask() Task {
	task := Task{
		Name:         "Check whether Description() is written",
		InspectionID: "UF002",
		Check:        syntax.NoDescriptionChecker,
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
