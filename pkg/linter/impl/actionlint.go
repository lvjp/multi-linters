package impl

import (
	"os"
	"os/exec"
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
					Build: `# renovate: datasource=docker
FROM rhysd/actionlint:1.7.4@sha256:82244e1db1c60d82c7792180a48dd0bcb838370bb589d53ff132503fc9485868 AS actionlint`,
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

		if err := cmd.Run(); err != nil {
			errCount++
		}
	}

	return
}
