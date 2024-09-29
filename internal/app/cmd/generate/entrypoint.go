package generate

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"

	"github.com/lvjp/multi-linters/pkg/descriptors"
)

//go:embed template.gotmpl
var templateText string

func Entrypoint() error {
	descriptors, err := descriptors.Load()
	if err != nil {
		return err
	}

	tmpl, err := template.New("docker").Parse(string(templateText))
	if err != nil {
		return fmt.Errorf("cannot parse template: %w", err)
	}

	if err := tmpl.Execute(os.Stdout, descriptors); err != nil {
		return fmt.Errorf("template rendering error: %w", err)
	}

	return nil
}
