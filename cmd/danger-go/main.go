package main

import (
	"fmt"
	"log"
	"os"

	"github.com/moolmanruan/danger-go/cmd/danger-go/runner"
	dangerJs "github.com/moolmanruan/danger-go/danger-js"
)

const version = "v0.0.3"

// main entrypoint of the danger-go command
func main() {
	// require process (always included) and command/flag
	if len(os.Args) < 2 {
		fmt.Println(usage)
	}

	if argsContain("-V", "--version") {
		fmt.Printf("danger-go %s\n", version)
		return
	}
	if argsContain("-h", "--help") {
		fmt.Println(usage)
		return
	}

	command := os.Args[1]
	switch command {
	case "ci", "local", "pr":
		var rest []string
		if len(os.Args) > 2 {
			rest = os.Args[2:]
		}
		err := dangerJs.Process(command, rest)
		if err != nil {
			log.Fatalf(err.Error())
		}
	case "runner":
		runner.Run()
	default:
		log.Fatalf("invalid sub-command `%s`\n\n%s", command, usage)
	}
}

// argsContain returns true if any of the provided `args` are in the list passed
// in to the command
func argsContain(args ...string) bool {
	for _, a := range os.Args {
		for _, v := range args {
			if a == v {
				return true
			}
		}
	}
	return false
}

const usage = `Usage: danger-go [options] [command]

Options:
  -V, --version  Output the version number
  -h, --help     Output usage information

Commands:
  init           Helps you get started with DSL
  ci             Runs DSL on CI
  pr             Runs your local Dangerfile against an existing GitHub DSL. Will not post on the DSL
  runner         Runs a dangerfile against a DSL passed in via STDIN [You probably don't need this]
  local          Runs danger standalone on a repo, useful for git hooks`
