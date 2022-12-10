package runner

import (
	"bufio"
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const dangerURLPrefix = "danger://dsl/"

// Run reads the danger DSL URL from stdin, invokes the Go dangerfile as a
// plugin, and then writes the results JSON to stdout.
func Run() {
	in := readAll()
	in = strings.TrimSpace(in)

	if !strings.HasPrefix(in, dangerURLPrefix) {
		log.Fatalf("did not receive a Danger URL")
	}

	jsonPath := strings.Replace(in, dangerURLPrefix, "", 1)
	_, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("failed to read JSON file at %s", jsonPath)
	}

	// TODO: Invoke the plugin built from the dangerfile.go here...

	resp := DangerResults{
		Fails:     []Violation{},
		Messages:  []Violation{},
		Warnings:  []Violation{},
		Markdowns: []Violation{},
	}
	resp.Warnings = append(resp.Warnings, Violation{Message: "foo bar baz!"})
	respBB, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("marshalling response: %s", err.Error())
	}
	fmt.Print(string(respBB))
}

// readAll reads everything on stdin until io.EOF and returns the result
func readAll() string {
	reader := bufio.NewReader(os.Stdin)
	var bb bytes.Buffer
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("reading stdin: %s", err.Error())
		}
		bb.Write(line)
	}
	return bb.String()
}
