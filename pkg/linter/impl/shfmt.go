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
					Build:   "FROM mvdan/shfmt:3.9.0@sha256:cb4ca87cc18e52f184a7ba1ae1ef7350b79a2c216ace78a0d24b473e87f0b8f5 AS shfmt",
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
