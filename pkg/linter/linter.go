package linter

type FileMatcher interface {
	OnFile(path string)
	Matched() bool
	Files() []string
}

type Linter interface {
	Descriptor() Descriptor
	FileMatcher() FileMatcher
	Activated() bool
	Execute() error
}

type Descriptor struct {
	ID         string
	Project    Project
	Mode       Mode
	Dockerfile Dockerfile
}
type Project struct {
	Name       string
	License    License
	Logo       string
	Repository string
	Website    string
}

type Dockerfile struct {
	Build    string
	Install  string
	Packages []Package
}

type Package struct {
	Name    string
	Version string
}

type License string

const (
	LicenseGPL3_0 License = "GPL-3.0-only"
)

type Mode string

const (
	ModeFilePerFile       Mode = "file_per_file"
	ModeSubFolder         Mode = "subfolder"
	ModeRandomGroupOfFile Mode = "random_group_of_files"
	ModeProject           Mode = "project"
)
