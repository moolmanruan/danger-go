package runner

import (
	"bufio"
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	danger_js "github.com/moolmanruan/danger-go/danger-js"
	"io"
	"log"
	"os"
	"os/exec"
	"plugin"
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
	prJSON, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("failed to read JSON file at %s", jsonPath)
	}

	var pr danger_js.PR
	err = json.Unmarshal(prJSON, &pr)
	if err != nil {
		log.Fatalf("failed to unmarshal PR JSON: %s", err.Error())
	}

	//TODO: Pass in file, not directory
	//TODO: Build in temp directory
	libPath, err := buildPlugin("/Users/ruan/danger-go/dangerfile")
	if err != nil {
		log.Fatalf("building plugin from dangerfile: %s", err.Error())
	}
	fn, err := loadPlugin(libPath)
	if err != nil {
		log.Fatalf("loading dangerfile plugin: %s", err.Error())
	}

	resp := fn(pr)
	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("marshalling response: %s", err.Error())
	}
	fmt.Print(string(respJSON))
}

func buildPlugin(dangerFilePath string) (string, error) {
	fmt.Println("Building dangerfile plugin in directory:", dangerFilePath)
	cmd := exec.Command("go", "build", "-buildmode=plugin", dangerFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return "dangerfile.so", nil
}

type MainFunc = func(pr danger_js.PR) DangerResults

func loadPlugin(libPath string) (MainFunc, error) {
	fmt.Println("Loading dangerfile plugin:", libPath)

	p, err := plugin.Open(libPath)
	if err != nil {
		return nil, err
	}

	dangerSymbol, err := p.Lookup("Run")
	if err != nil {
		return nil, err
	}

	dangerFn, ok := dangerSymbol.(MainFunc)
	if !ok {
		return nil, errors.New("failed to cast Danger function")
	}

	return dangerFn, nil
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
