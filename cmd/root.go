package cmd

import (
	"fmt"
	"os"

	"github.com/nao1215/goalinter-v1/internal/lint"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goalinter-v1",
	Short: "goalinter-v1 is linter for goa-v1 (not original one, it's forked project)",
	Run: func(cmd *cobra.Command, args []string) {
		lint.Run()
		os.Exit(0)
	},
}

func exitError(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

// Execute run command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}
