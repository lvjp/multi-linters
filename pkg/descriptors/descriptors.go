package descriptors

import (
	"embed"
	"fmt"
	"io/fs"

	"gopkg.in/yaml.v3"
)

//go:embed linters/*.yaml
//go:embed schemas/*.json
var descriptorsFS embed.FS

func Load() ([]Linter, error) {
	paths, err := listFiles()
	if err != nil {
		return nil, err
	}

	ids := make(map[string]any, len(paths))
	linters := make([]Linter, 0, len(paths))

	for _, path := range paths {
		raw, err := fs.ReadFile(descriptorsFS, path)
		if err != nil {
			return nil, fmt.Errorf("cannot read file [%s]: %w", path, err)
		}

		var linter Linter
		if err := yaml.Unmarshal(raw, &linter); err != nil {
			return nil, fmt.Errorf("cannot unmarshal file [%s]: %v", path, err)
		}

		if _, exists := ids[linter.ID]; exists {
			return nil, fmt.Errorf("try to register again [%s] with file [%s]", linter.ID, path)
		}

		ids[linter.ID] = linter
		linters = append(linters, linter)
	}

	return linters, nil
}

func listFiles() ([]string, error) {
	var paths []string

	const rootPath = "linters"

	err := fs.WalkDir(descriptorsFS, rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root
		if path == rootPath {
			return nil
		}

		paths = append(paths, path)

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("cannot list descriptors files: %w", err)
	}

	return paths, nil
}

type Linter struct {
	ID      string
	Project struct {
		Name       string
		License    string
		Logo       string
		Repository string
		Website    string
	}
	Capabilities struct {
		Mode  Mode
		Sarif *struct {
			Args []string
		}
	}
	Dockerfile struct {
		Build   string
		Install string
	}
}

type Mode string

const (
	ModeFilePerFile       Mode = "file_per_file"
	ModeRandomGroupOfFile Mode = "random_group_of_files"
	ModeProject           Mode = "project"
)
