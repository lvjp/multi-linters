package cmd

import (
	"fmt"

	"github.com/lvjp/multi-linters/internal/app/buildinfo"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print multi-linters version information",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(buildinfo.VersionString())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
