package registry

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/lvjp/multi-linters/pkg/linter"
)

var lintersRegistry = make(map[string]linter.Linter)

func Register(linter linter.Linter) {
	if linter == nil {
		panic("Must not provide nil Linter")
	}

	id := linter.Descriptor().ID
	_, registered := lintersRegistry[id]
	if registered {
		panic(fmt.Sprintf("Linter named %s already registered", id))
	}

	lintersRegistry[id] = linter
}

func Linters() []linter.Linter {
	return slices.SortedFunc(
		maps.Values(lintersRegistry),
		func(a, b linter.Linter) int {
			return strings.Compare(a.Descriptor().ID, b.Descriptor().ID)
		},
	)
}
