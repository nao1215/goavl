package cmd

import (
	"github.com/nao1215/goavl/internal/lint"
	"github.com/spf13/cobra"
)

var astCmd = &cobra.Command{
	Use: "ast",
	Run: func(cmd *cobra.Command, args []string) {
		lint.PrintAST(args)
	},
	Short: "DEBUG: Print abstract syntax tree of the go file specified in the arguments",
}

func init() {
	rootCmd.AddCommand(astCmd)
}
