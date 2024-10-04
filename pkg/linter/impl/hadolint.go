package impl

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lvjp/multi-linters/pkg/linter"
	"github.com/lvjp/multi-linters/pkg/linter/registry"
)

func init() {
	registry.Register(&hadolint{
		baseLinter: baseLinter{
			descriptor: linter.Descriptor{
				ID: "hadolint",
				Project: linter.Project{
					Name:       "hadolint",
					License:    linter.LicenseGPL3_0,
					Logo:       "https://hadolint.github.io/hadolint/img/cat_container.png",
					Repository: "https://github.com/hadolint/hadolint",
					Website:    "https://hadolint.github.io/hadolint",
				},
				Mode: linter.ModeElligibleFiles,
				Dockerfile: linter.Dockerfile{
					Build:   "FROM hadolint/hadolint:v2.12.0-alpine@sha256:3c206a451cec6d486367e758645269fd7d696c5ccb6ff59d8b03b0e45268a199 AS hadolint",
					Install: "COPY --link --from=hadolint /bin/hadolint /usr/bin/hadolint",
				},
			},
			fileMatcher: NewBasicFileMatcher(NewFileMatcherAggregator(
				NewExtensionFileMatcher(".Dockerfile"),
				NewFileNameFileMatcher("Dockerfile"),
			)),
		},
	})
}

type hadolint struct {
	baseLinter
}

func (l *hadolint) Execute() error {
	for _, file := range l.fileMatcher.Files() {
		cmd := exec.Command("hadolint", file)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("hadolint error on file %s", file)
		}
	}

	return nil
}
