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
FROM zricethezav/gitleaks:v8.21.1@sha256:9f4bdc62e5f4e4ae915341f3d957b7b5fff099a37ab2f44ffa08fe5b04a95a6d AS gitleaks`,
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
