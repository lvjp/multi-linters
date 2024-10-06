package impl

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lvjp/multi-linters/pkg/linter"
	"github.com/lvjp/multi-linters/pkg/linter/registry"
)

func init() {
	registry.Register(&actionlint{
		baseLinter: baseLinter{
			descriptor: linter.Descriptor{
				ID: "actionlint",
				Project: linter.Project{
					Name:       "actionlint",
					License:    linter.LicenseMIT,
					Repository: "https://github.com/rhysd/actionlint",
					Website:    "https://rhysd.github.io/actionlint/",
				},
				Mode: linter.ModeElligibleFiles,
				Dockerfile: linter.Dockerfile{
					Build:   `FROM rhysd/actionlint:1.7.3@sha256:7617f05bd698cd2f1c3aedc05bc733ccec92cca0738f3e8722c32c5b42c70ae6 AS actionlint`,
					Install: `COPY --link --from=actionlint /usr/local/bin/actionlint /usr/bin/actionlint`,
				},
			},
			fileMatcher: NewBasicFileMatcher(NewAndFileMatcherCriteria(
				func(path string) bool {
					return strings.HasPrefix(path, ".github/workflows/")
				},
				NewExtensionFileMatcher(".yml", ".yaml"),
			)),
		},
	})
}

type actionlint struct {
	baseLinter
}

func (l *actionlint) Execute() (errCount int) {
	for _, file := range l.fileMatcher.Files() {
		cmd := exec.Command("actionlint", file)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = filepath.Dir(file)

		if err := cmd.Run(); err != nil {
			errCount++
		}
	}

	return
}
