package impl

import (
	"path/filepath"
	"strings"

	"github.com/lvjp/multi-linters/pkg/linter"
)

type baseLinter struct {
	descriptor  linter.Descriptor
	fileMatcher linter.FileMatcher
}

func (l *baseLinter) Descriptor() linter.Descriptor {
	return l.descriptor
}

func (l *baseLinter) FileMatcher() linter.FileMatcher {
	return l.fileMatcher
}

func (l *baseLinter) Activated() bool {
	return l.fileMatcher.Matched()
}

type FileMatcherCriteria func(path string) bool

type basicFileMatcher struct {
	criteria FileMatcherCriteria
	matched  []string
}

func NewBasicFileMatcher(criteria FileMatcherCriteria) linter.FileMatcher {
	return &basicFileMatcher{
		criteria: criteria,
	}
}

func (fm *basicFileMatcher) OnFile(path string) {
	if fm.criteria(path) {
		fm.matched = append(fm.matched, path)
	}
}

func (fm *basicFileMatcher) Matched() bool {
	return len(fm.matched) > 0
}

func (fm *basicFileMatcher) Files() []string {
	return fm.matched
}

func NewFileMatcherAggregator(criterias ...FileMatcherCriteria) FileMatcherCriteria {
	return func(path string) bool {
		for _, c := range criterias {
			if c(path) {
				return true
			}
		}

		return false
	}
}

func NewAndFileMatcherCriteria(criterias ...FileMatcherCriteria) FileMatcherCriteria {
	return func(path string) bool {
		for _, c := range criterias {
			if !c(path) {
				return false
			}
		}

		return true
	}
}

func NewAcceptAllFileMatcherCriteria() FileMatcherCriteria {
	return func(_ string) bool {
		return true
	}
}

func NewExtensionFileMatcher(extensions ...string) FileMatcherCriteria {
	return func(path string) bool {
		base := filepath.Base(path)

		for _, ext := range extensions {
			if strings.HasSuffix(base, ext) {
				return true
			}
		}

		return false
	}
}

func NewFileNameFileMatcher(names ...string) FileMatcherCriteria {
	return func(path string) bool {
		base := filepath.Base(path)

		for _, name := range names {
			if base == name {
				return true
			}
		}

		return false
	}
}
