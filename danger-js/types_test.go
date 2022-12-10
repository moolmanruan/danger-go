package danger_js_test

import (
	danger_js "danger-go/danger-js"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

import _ "embed"

//go:embed pr_github.json
var gitHubPRJSON string

//go:embed pr_gitlab.json
var gitLabPRJSON string

func TestUnmarshalGitHub(t *testing.T) {
	var pr danger_js.PR
	err := json.Unmarshal([]byte(gitHubPRJSON), &pr)
	require.Nil(t, err)
}

func TestUnmarshalGitLab(t *testing.T) {
	var pr danger_js.PR
	err := json.Unmarshal([]byte(gitLabPRJSON), &pr)
	require.Nil(t, err)
}
