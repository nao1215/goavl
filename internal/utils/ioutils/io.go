package ioutils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Die print message with command name. After printing, exit 1.
func Die(msg string) {
	fmt.Fprintf(os.Stderr, "[%s] goavl: %s\n", color.RedString("ERROR"), msg)
	os.Exit(1)
}

// Warn print message with command name.
func Warn(msg string) {
	fmt.Fprintf(os.Stderr, "[%s] goavl: %s\n", color.YellowString("WARN"), msg)
}
