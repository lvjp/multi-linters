package buildinfo

import (
	"fmt"
	"runtime/debug"
)

type BuildInfo struct {
	Revision     string
	RevisionTime string
	Modified     bool

	GoVersion string
	GoOS      string
	GoArch    string
}

func (bi *BuildInfo) VersionString() string {
	if bi == nil {
		return "unknown"
	}

	ret := fmt.Sprintf(
		"%s %s %s %s/%s",
		bi.Revision,
		bi.RevisionTime,
		bi.GoVersion,
		bi.GoOS,
		bi.GoArch,
	)

	return ret
}

func VersionString() string {
	raw, ok := debug.ReadBuildInfo()
	if !ok {
		return "-"
	}

	bi := &BuildInfo{
		Revision:     "-",
		RevisionTime: "-",
		Modified:     false,
		GoVersion:    raw.GoVersion,
		GoOS:         "-",
		GoArch:       "-",
	}

	for _, s := range raw.Settings {
		switch s.Key {
		case "vcs.revision":
			bi.Revision = s.Value
		case "vcs.time":
			bi.RevisionTime = s.Value
		case "vcs.modified":
			bi.Modified = s.Value == "true"
		case "GOOS":
			bi.GoOS = s.Value
		case "GOARCH":
			bi.GoArch = s.Value
		}
	}

	str := bi.VersionString()

	if bi.Modified {
		str += " (modified)"
	}

	return str
}
