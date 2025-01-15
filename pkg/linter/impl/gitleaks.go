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
FROM zricethezav/gitleaks:v8.23.1@sha256:a695fbcee9f56b8f3d8a121caa078c2d8adbb31989cd28b213b90e336be2dbb9 AS gitleaks`,
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
