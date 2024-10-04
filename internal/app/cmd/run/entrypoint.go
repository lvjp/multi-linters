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

	"github.com/lvjp/multi-linters/pkg/linter/registry"

	"github.com/jedib0t/go-pretty/v6/table"
)

func Entrypoint() error {
	if err := computeActivation(); err != nil {
		return err
	}

	fmt.Println()

	runLinters()

	return nil
}

func computeActivation() error {
	files, err := listFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		for _, l := range registry.Linters() {
			l.FileMatcher().OnFile(file)
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

func runLinters() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAutoIndex(true)
	t.SetTitle("Execution summary")
	t.AppendHeader(table.Row{"ID", "Mode", "Files", "Errors", "status", "Elapsed time"})

	for _, linter := range registry.Linters() {
		if !linter.Activated() {
			continue
		}

		descriptor := linter.Descriptor()
		fmt.Println("### Launch", descriptor.ID)
		begin := time.Now()
		errCount := linter.Execute()
		elapsed := time.Since(begin).Truncate(time.Millisecond)
		fmt.Println("Elapsed time:", elapsed)

		var status string
		if errCount > 0 {
			status = "❌"
			fmt.Println("❌ Execution error")
		} else {
			status = "✅"
			fmt.Println("✅ Execution succeeded !")
		}

		t.AppendRow(table.Row{
			descriptor.ID,
			descriptor.Mode,
			len(linter.FileMatcher().Files()),
			errCount,
			status,
			elapsed,
		})
		fmt.Println()
	}

	t.Render()
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
