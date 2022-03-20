package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is goavl version.
	Version = "0.3.2"
)

func getVersion() string {
	return fmt.Sprintf("goavl version " + Version + " (under Apache License version 2.0)")
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
