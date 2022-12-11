package one

import (
	danger "github.com/moolmanruan/danger-go"
	"strings"
)

func DoTheThing(info danger.Info) {
	danger.Post("Files:\n" + strings.Join(info.Git().Files, "\n"))
}
