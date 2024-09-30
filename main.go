package main

import (
	"github.com/lvjp/multi-linters/cmd"
	_ "github.com/lvjp/multi-linters/pkg/linter/impl"
)

func main() {
	cmd.Execute()
}
