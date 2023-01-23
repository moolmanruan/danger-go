package gofmt

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/moolmanruan/danger-go"
	"github.com/moolmanruan/danger-go/danger-js"
)

func Run(d *danger.T, pr dangerJs.DSL) {
	ff := make([]dangerJs.FilePath, 0)
	for _, f := range append(pr.Git.CreateFiles, pr.Git.ModifiedFiles...) {
		if strings.HasSuffix(f, ".go") {
			ff = append(ff, f)
		}
	}
	args := append([]string{"-l"}, ff...)
	cmd := exec.Command("gofmt", args...)
	bb, err := cmd.CombinedOutput()
	if err != nil {
		d.Warn(fmt.Sprintf("Failed to run gofmt plugin:\n%s", bytes.TrimSpace(bb)), "", 0)
		return
	}
	if len(bb) > 0 {
		d.Fail(fmt.Sprintf("These files need to be formatted with gofmt:\n%s", bytes.TrimSpace(bb)), "", 0)
	}
}
