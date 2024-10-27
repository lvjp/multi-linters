package impl

import (
	"fmt"
	"os"
	"os/exec"

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
FROM zricethezav/gitleaks:v8.21.2@sha256:0e99e8821643ea5b235718642b93bb32486af9c8162c8b8731f7cbdc951a7f46 AS gitleaks`,
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
		if !l.executeFile(file) {
			errCount++
		}
	}

	return
}

func (l *gitLeaks) executeFile(file string) bool {
	fp, err := os.Open(file)
	if err != nil {
		return false
	}

	defer fp.Close()

	fmt.Println(">", file)
	cmd := exec.Command("gitleaks", "stdin", "--no-banner", "--redact", "--verbose", "--report-format=json")
	cmd.Stdin = fp
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run() == nil
}
