package impl

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/lvjp/multi-linters/pkg/linter"
	"github.com/lvjp/multi-linters/pkg/linter/registry"
)

func init() {
	registry.Register(&golangciLint{
		baseLinter: baseLinter{
			descriptor: linter.Descriptor{
				ID: "golangci-lint",
				Project: linter.Project{
					Name:       "golangci-lint",
					License:    linter.LicenseGPL3_0,
					Logo:       "https://github.com/golangci/golangci-lint/blob/master/assets/go.png",
					Repository: "https://github.com/golangci/golangci-lint",
					Website:    "https://golangci-lint.run/",
				},
				Mode: linter.ModeSubFolder,
				Dockerfile: linter.Dockerfile{
					Build: `FROM golang:1.23.2-alpine@sha256:9dd2625a1ff2859b8d8b01d8f7822c0f528942fe56cfe7a1e7c38d3b8d72d679 AS golang
FROM golangci/golangci-lint:v1.61.0-alpine@sha256:61e2d68adc792393fcb600340fe5c28059638d813869d5b4c9502392a2fb4c96 AS golangci-lint`,
					Install: `COPY --from=golang /usr/local/go/go.env /usr/lib/go/
COPY --from=golang /usr/local/go/bin/ /usr/lib/go/bin/
COPY --from=golang /usr/local/go/lib/ /usr/lib/go/lib/
COPY --from=golang /usr/local/go/pkg/ /usr/lib/go/pkg/
COPY --from=golang /usr/local/go/src/ /usr/lib/go/src/
COPY --link --from=golangci-lint /usr/bin/golangci-lint /usr/bin/golangci-lint
ENV PATH="${PATH}:/usr/lib/go/bin"`,
					Packages: []linter.Package{
						{Name: "git", Version: "2.45.2-r0"},
						{Name: "mercurial", Version: "6.7.4-r1"},
						{Name: "subversion", Version: "1.14.3-r2"},
						{Name: "fossil", Version: "2.24-r1"},
					},
				},
			},
			fileMatcher: NewBasicFileMatcher(NewFileNameFileMatcher("go.mod")),
		},
	})
}

type golangciLint struct {
	baseLinter
}

func (l *golangciLint) Execute() error {
	for _, file := range l.fileMatcher.Files() {
		cmd := exec.Command("golangci-lint", "run")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = filepath.Dir(file)

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("golangci-lint error on file %s: %w", file, err)
		}
	}

	return nil
}
