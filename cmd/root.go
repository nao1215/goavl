package cmd

import (
	"os"

	"github.com/nao1215/goavl/internal/lint"
	"github.com/nao1215/goavl/internal/utils/ioutils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goavl",
	Short: "goavl is linter for goa-v1 (not original one, it's forked project)",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			ioutils.Die(err.Error())
		}
		exclude, err := cmd.Flags().GetStringSlice("exclude")
		if err != nil {
			ioutils.Die(err.Error())
		}

		if file == "" {
			lint.Run(args, exclude)
		} else {
			lint.CheckOneFile(file)
		}
		os.Exit(0)
	},
}

// Execute run command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ioutils.Die(err.Error())
	}
}

func init() {
	rootCmd.Flags().StringP("file", "f", "", "specify the file to be checked")
	rootCmd.Flags().StringSliceP("exclude", "e", []string{}, "specify inspection ID to be excluded by comma separating")
}
