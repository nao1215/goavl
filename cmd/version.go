package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is goalinter-v1 version.
	Version = "0.0.1"
)

func getVersion() string {
	return fmt.Sprintf("goalinter-v1 version " + Version + " (under Apache License version 2.0)")
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getVersion())
	},
	Short: "Show version info",
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
