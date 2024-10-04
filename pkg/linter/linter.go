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
	Execute() (errCount int)
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
	LicenseBSD    License = "BSD-3-Clause"
	LicenseGPL3_0 License = "GPL-3.0-only"
	LicenseMIT    License = "MIT"
)

type Mode string

const (
	ModeElligibleFiles Mode = "elligible_files"
	ModeSubFolder      Mode = "subfolder"
	ModeAllFiles       Mode = "all_files"
)
