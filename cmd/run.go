package cmd

import (
	"github.com/lvjp/multi-linters/internal/app/cmd/run"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run all linter againt the current folder",
	Run: func(_ *cobra.Command, _ []string) {
		if err := run.Entrypoint(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
