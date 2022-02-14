package cmd

import (
	"github.com/nao1215/goavl/internal/lint"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		lint.PrintCheckTaskList()
	},
	Short: "Print inspection id and inspection details",
}

func init() {
	rootCmd.AddCommand(listCmd)
}
