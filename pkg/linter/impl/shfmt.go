package impl

import (
	"os"
	"os/exec"

	"github.com/lvjp/multi-linters/pkg/linter"
	"github.com/lvjp/multi-linters/pkg/linter/registry"
)

func init() {
	registry.Register(&shfmt{
		baseLinter: baseLinter{
			descriptor: linter.Descriptor{
				ID: "shfmt",
				Project: linter.Project{
					Name:       "shfmt",
					License:    linter.LicenseBSD,
					Repository: "https://github.com/mvdan/sh",
				},
				Mode: linter.ModeElligibleFiles,
				Dockerfile: linter.Dockerfile{
					Build: `# renovate: datasource=docker
FROM mvdan/shfmt:v3.10.0@sha256:d19cc37644449fe9a488f234d2c0cf0b770eaf6a5a40e30103e8099013ef8f9e AS shfmt`,
					Install: "COPY --link --from=shfmt /bin/shfmt /usr/bin/",
				},
			},
			fileMatcher: NewBasicFileMatcher(NewFileMatcherAggregator(
				NewExtensionFileMatcher(".sh", ".bash", ".dash", ".ksh"),
			)),
		},
	})
}

type shfmt struct {
	baseLinter
}

func (l *shfmt) Execute() (errCount int) {
	for _, file := range l.fileMatcher.Files() {
		cmd := exec.Command("shfmt", file)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			errCount++
		}
	}

	return
}
