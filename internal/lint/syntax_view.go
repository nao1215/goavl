package lint

import (
	"fmt"
)

// NewSyntaxViewTask return task that check View() function syntax
func NewSyntaxViewTask() Task {
	task := Task{
		Name: "View() syntax check",
		Check: func() {
			fmt.Println("View() syntax check")
		},
	}
	return task
}
