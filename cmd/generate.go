package cmd

import (
	"github.com/lvjp/multi-linters/internal/app/cmd/generate"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate multi-linters Dockerfile",
	Run: func(_ *cobra.Command, _ []string) {
		if err := generate.Entrypoint(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
