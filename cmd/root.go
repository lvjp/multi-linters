package cmd

import (
	"github.com/lvjp/multi-linters/internal/app/buildinfo"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "multi-linters",
	Version: buildinfo.VersionString(),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
