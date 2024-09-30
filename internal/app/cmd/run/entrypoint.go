package run

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/lvjp/multi-linters/pkg/linter"
	"github.com/lvjp/multi-linters/pkg/linter/registry"

	"github.com/jedib0t/go-pretty/v6/table"
)

func Entrypoint() error {
	if err := computeActivation(); err != nil {
		return err
	}

	fmt.Println()

	if err := runLinters(); err != nil {
		return err
	}

	return nil
}

func computeActivation() error {
	files, err := listFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		for _, l := range registry.Linters() {
			if l.Descriptor().Mode != linter.ModeProject {
				l.FileMatcher().OnFile(file)
			}
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAutoIndex(true)
	t.SetTitle("Linters activation")
	t.AppendHeader(table.Row{"ID", "Mode", "Activated"})

	for _, l := range registry.Linters() {
		descriptor := l.Descriptor()
		t.AppendRow(table.Row{
			descriptor.ID,
			descriptor.Mode,
			l.Activated(),
		})
	}

	t.Render()
	return nil
}

func runLinters() error {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAutoIndex(true)
	t.SetTitle("Execution summary")
	t.AppendHeader(table.Row{"ID", "Mode", "Status", "Elapsed time"})

	var lastError error
	for _, linter := range registry.Linters() {
		if !linter.Activated() {
			continue
		}

		descriptor := linter.Descriptor()
		fmt.Println("### Launch", descriptor.ID)
		begin := time.Now()
		err := linter.Execute()
		elapsed := time.Since(begin).Truncate(time.Millisecond)
		fmt.Println("Elapsed time:", elapsed)

		var status string
		if err != nil {
			status = "failed"
			fmt.Println("Execution error:", err)
			lastError = fmt.Errorf("failed to execute %s linter: %w", descriptor.ID, err)
		} else {
			status = "succeeded"
			fmt.Println("Execution succeeded !")
		}

		t.AppendRow(table.Row{descriptor.ID, descriptor.Mode, status, elapsed})
		fmt.Println()
	}

	t.Render()

	return lastError
}

func listFiles() ([]string, error) {
	cmd := exec.Command("git", "ls-files", "-z")
	cmd.Stderr = os.Stderr

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("cannot run `git ls-files -z`: %w", err)
	}

	buf := bufio.NewReader(bytes.NewReader(out.Bytes()))

	var files []string

	for {
		line, err := buf.ReadBytes(0)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		files = append(files, string(line[0:len(line)-1]))
	}

	return files, nil
}
