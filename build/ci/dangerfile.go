package main

import (
	"fmt"
	"github.com/moolmanruan/danger-go/cmd/danger-go/runner"
	danger_js "github.com/moolmanruan/danger-go/danger-js"
)

// Run is invoked by danger-go
func Run(pr danger_js.PR) runner.DangerResults {
	resp := runner.DangerResults{
		Fails:     []runner.Violation{},
		Messages:  []runner.Violation{},
		Warnings:  []runner.Violation{},
		Markdowns: []runner.Violation{},
	}
	resp.Messages = append(resp.Messages, runner.Violation{
		Message: fmt.Sprintf("%d new files added!", len(pr.Danger.Git.CreateFiles)),
	})
	return resp
}
