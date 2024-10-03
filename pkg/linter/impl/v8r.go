package impl

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lvjp/multi-linters/pkg/linter"
	"github.com/lvjp/multi-linters/pkg/linter/registry"
)

func init() {
	registry.Register(&v8r{
		baseLinter: baseLinter{
			descriptor: linter.Descriptor{
				ID: "v8r",
				Project: linter.Project{
					Name:       "v8r",
					License:    linter.LicenseMIT,
					Logo:       "https://chris48s.github.io/v8r/img/logo.png",
					Repository: "https://github.com/chris48s/v8r",
					Website:    "https://chris48s.github.io/v8r/",
				},
				Mode: linter.ModeElligibleFiles,
				Dockerfile: linter.Dockerfile{
					Install: "RUN npm -g install v8r@4.1.0",
					Packages: []linter.Package{
						{Name: "npm", Version: "10.8.0-r0"},
					},
				},
			},
			fileMatcher: NewBasicFileMatcher(NewExtensionFileMatcher(
				".json",
				".yml",
				".yaml",
			)),
		},
	})
}

type v8r struct {
	baseLinter
}

func (l *v8r) Execute() (errCount int) {
	for _, file := range l.fileMatcher.Files() {
		// cmd := exec.Command("v8r", file)
		cmd := exec.Command("v8r", "--ignore-errors", file)
		fmt.Println(cmd.Args)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			// return fmt.Errorf("v8r error on file %s: %w", file, err)
			fmt.Printf("v8r error on file %s: %v\n", file, err)
		}
	}

	return 0
}
