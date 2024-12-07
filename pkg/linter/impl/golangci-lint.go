package impl

import (
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
					Build: `# renovate: datasource=docker
FROM golang:1.23.4-alpine@sha256:6c5c9590f169f77c8046e45c611d3b28fe477789acd8d3762d23d4744de69812 AS golang
# renovate: datasource=docker
FROM golangci/golangci-lint:v1.62.2-alpine@sha256:0f3af3929517ed4afa1f1bcba4eae827296017720e08ecd5c68b9f0640bc310d AS golangci-lint`,
					Install: `COPY --from=golang /usr/local/go/go.env /usr/lib/go/
COPY --from=golang /usr/local/go/bin/ /usr/lib/go/bin/
COPY --from=golang /usr/local/go/lib/ /usr/lib/go/lib/
COPY --from=golang /usr/local/go/pkg/ /usr/lib/go/pkg/
COPY --from=golang /usr/local/go/src/ /usr/lib/go/src/
COPY --link --from=golangci-lint /usr/bin/golangci-lint /usr/bin/golangci-lint
ENV PATH="${PATH}:/usr/lib/go/bin"`,
					Packages: []linter.Package{
						{Name: "git", Version: "2.47.1-r0"},
						{Name: "mercurial", Version: "6.9-r0"},
						{Name: "subversion", Version: "1.14.4-r0"},
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

func (l *golangciLint) Execute() (errCount int) {
	for _, file := range l.fileMatcher.Files() {
		cmd := exec.Command("golangci-lint", "run")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = filepath.Dir(file)

		if err := cmd.Run(); err != nil {
			errCount++
		}
	}

	return
}
