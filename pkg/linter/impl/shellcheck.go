package impl

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lvjp/multi-linters/pkg/linter"
	"github.com/lvjp/multi-linters/pkg/linter/registry"
)

func init() {
	registry.Register(&shellcheck{
		baseLinter: baseLinter{
			descriptor: linter.Descriptor{
				ID: "shellcheck",
				Project: linter.Project{
					Name:       "ShellCheck",
					License:    linter.LicenseGPL3_0,
					Logo:       "https://raw.githubusercontent.com/koalaman/shellcheck/refs/heads/master/doc/shellcheck_logo.svg",
					Repository: "https://github.com/koalaman/shellcheck",
					Website:    "https://www.shellcheck.net/",
				},
				Mode: linter.ModeElligibleFiles,
				Dockerfile: linter.Dockerfile{
					Build:   "FROM koalaman/shellcheck:v0.10.0@sha256:2097951f02e735b613f4a34de20c40f937a6c8f18ecb170612c88c34517221fb AS shellcheck",
					Install: "COPY --link --from=shellcheck /bin/shellcheck /usr/bin/shellcheck",
				},
			},
			fileMatcher: NewBasicFileMatcher(NewFileMatcherAggregator(
				NewExtensionFileMatcher(".sh", ".bash", ".dash", ".ksh"),
			)),
		},
	})
}

type shellcheck struct {
	baseLinter
}

func (l *shellcheck) Execute() error {
	for _, file := range l.fileMatcher.Files() {
		cmd := exec.Command("shellcheck", file)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("shellcheck error on file %s", file)
		}
	}

	return nil
}
