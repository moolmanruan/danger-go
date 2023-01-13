package main

import (
	"fmt"

	"github.com/moolmanruan/danger-go"
	"github.com/moolmanruan/danger-go/danger-js"
)

// Run is invoked by danger-go
func Run(d *danger.T, pr dangerJs.DSL) {
	d.Message(fmt.Sprintf("%d new files added!", len(pr.Git.CreateFiles)), "", 0)
	d.Message(fmt.Sprintf("%d files modified!", len(pr.Git.ModifiedFiles)), "", 0)
}
