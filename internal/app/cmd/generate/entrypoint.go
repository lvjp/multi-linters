package generate

import (
	_ "embed"
	"fmt"
	"maps"
	"os"
	"slices"
	"text/template"

	"github.com/lvjp/multi-linters/pkg/linter"
	"github.com/lvjp/multi-linters/pkg/linter/registry"
)

//go:embed template.gotmpl
var templateText string

func Entrypoint() error {
	tmpl, err := template.New("docker").Parse(templateText)
	if err != nil {
		return fmt.Errorf("cannot parse template: %w", err)
	}

	data := map[string]any{
		"Linters":  registry.Linters(),
		"Packages": listPackages(),
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		return fmt.Errorf("template rendering error: %w", err)
	}

	return nil
}

func listPackages() []linter.Package {
	packages := map[string]linter.Package{
		"git": {Name: "git", Version: "2.47.1-r0"},
	}

	for _, l := range registry.Linters() {
		for _, pkg := range l.Descriptor().Dockerfile.Packages {
			if _, exists := packages[pkg.Name]; exists {
				if packages[pkg.Name].Version != pkg.Version {
					panic(fmt.Sprintf(
						"Multiple versions [%s] [%s] of %s detected with %s",
						packages[pkg.Name].Version,
						pkg.Version,
						pkg.Name,
						l.Descriptor().ID,
					))
				}
			} else {
				packages[pkg.Name] = pkg
			}
		}
	}

	list := make([]linter.Package, len(packages))

	for i, name := range slices.Sorted(maps.Keys(packages)) {
		list[i] = packages[name]
	}

	return list
}
