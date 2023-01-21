package danger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {
	d := New()

	d.Message("a message", "", 0)

	require.Equal(t,
		[]Violation{
			{Message: "a message"},
		},
		d.results.Messages)
}

func TestWarn(t *testing.T) {
	d := New()

	d.Warn("a warning", "", 0)

	require.Equal(t,
		[]Violation{
			{Message: "a warning"},
		},
		d.results.Warnings)
}

func TestFail(t *testing.T) {
	d := New()
	d.Fail("a failure", "", 0)

	require.Equal(t,
		[]Violation{
			{Message: "a failure"},
		},
		d.results.Fails)
}

func TestMarkdown(t *testing.T) {
	d := New()

	d.Markdown("some markdown", "", 0)

	require.Equal(t,
		[]Violation{
			{Message: "some markdown"},
		},
		d.results.Markdowns)
}
