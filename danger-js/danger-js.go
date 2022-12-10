package danger_js

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const dangerJsBinary = "danger"
const dangerGoBinary = "danger-go"

func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	stdout := new(strings.Builder)
	cmd.Stdout = stdout
	stderr := new(strings.Builder)
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		parts := strings.Split(stderr.String(), "\n\n")
		for _, part := range parts {
			if strings.HasPrefix(part, "Error") {
				return "", errors.New(part)
			}
		}
		return "", err
	}
	return stdout.String(), nil
}

func findBinary(name string) (string, error) {
	dangerBin, err := runCommand("which", name)
	if err != nil {
		return "", fmt.Errorf("could not find `%s` binary: %w", name, err)
	}
	return strings.TrimSpace(dangerBin), nil
}

func GetPR(url string, dangerBin string) (PR, error) {
	var err error
	if dangerBin == "" {
		dangerBin, err = findBinary(dangerJsBinary)
		if err != nil {
			return PR{}, err
		}
	}

	prJSON, err := runCommand(dangerBin, "pr", url, "--json")
	if err != nil {
		return PR{}, fmt.Errorf("could not download PR JSON with danger-js: %w", err)
	}

	var pr PR
	if err = json.Unmarshal([]byte(prJSON), &pr); err != nil {
		return PR{}, err
	}
	return pr, nil
}

func Process(command string, args []string) error {
	dangerBin, err := findBinary(dangerJsBinary)
	if err != nil {
		return err
	}
	dangerGoBin, err := findBinary(dangerGoBinary)
	if err != nil {
		return err
	}
	// The `danger` (javascript) command will call the process specified,
	// ie. `danger-go`, with the first argument of `runner` followed by the
	// arguments it received.
	cmdArgs := append([]string{command, "--process", dangerGoBin, "--passURLForDSL"}, args...)
	cmd := exec.Command(dangerBin, cmdArgs...)
	fmt.Printf("Running: %s\n", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
