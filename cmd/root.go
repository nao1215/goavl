package cmd

import (
	"log"
	"os"

	"github.com/nao1215/goavl/internal/lint"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goavl",
	Short: "goavl is linter for goa-v1 (not original one, it's forked project)",
	Run: func(cmd *cobra.Command, args []string) {
		lint.Run()
		os.Exit(0)
	},
}

// Execute run command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
