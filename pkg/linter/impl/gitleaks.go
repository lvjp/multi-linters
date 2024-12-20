package impl

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/lvjp/multi-linters/pkg/linter"
	"github.com/lvjp/multi-linters/pkg/linter/registry"
)

func init() {
	registry.Register(&gitLeaks{
		baseLinter: baseLinter{
			descriptor: linter.Descriptor{
				ID: "gitleaks",
				Project: linter.Project{
					Name:       "GitLeaks",
					License:    linter.LicenseMIT,
					Logo:       "https://gitleaks.io/logo.png",
					Repository: "https://github.com/gitleaks/gitleaks",
					Website:    "https://gitleaks.io/",
				},
				Mode: linter.ModeAllFiles,
				Dockerfile: linter.Dockerfile{
					Build: `# renovate: datasource=docker
FROM zricethezav/gitleaks:v8.22.0@sha256:9c008fc701d3a0a2a15b77a7677b940e2e5942e8af8cc2a0b6e078c4365c3c86 AS gitleaks`,
					Install: `COPY --link --from=gitleaks /usr/bin/gitleaks /usr/bin/`,
				},
			},
			fileMatcher: NewBasicFileMatcher(NewAcceptAllFileMatcherCriteria()),
		},
	})
}

type gitLeaks struct {
	baseLinter
}

func (l *gitLeaks) Execute() (errCount int) {
	for _, file := range l.fileMatcher.Files() {
		cmd := exec.Command("gitleaks", "detect", "--no-banner", "--no-git", "--redact", file)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = filepath.Dir(file)

		if err := cmd.Run(); err != nil {
			errCount++
		}
	}

	return
}
