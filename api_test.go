package danger_test

import (
	"testing"

	"github.com/moolmanruan/danger-go"
	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {
	d := danger.New()

	d.Message("a message", "", 0)

	require.Equal(t,
		[]danger.Violation{
			{Message: "a message"},
		},
		d.Results.Messages)
}

func TestWarn(t *testing.T) {
	d := danger.New()

	d.Warn("a warning", "", 0)

	require.Equal(t,
		[]danger.Violation{
			{Message: "a warning"},
		},
		d.Results.Warnings)
}

func TestFail(t *testing.T) {
	d := danger.New()
	d.Fail("a failure", "", 0)

	require.Equal(t,
		[]danger.Violation{
			{Message: "a failure"},
		},
		d.Results.Fails)
}

func TestMarkdown(t *testing.T) {
	d := danger.New()

	d.Markdown("some markdown", "", 0)

	require.Equal(t,
		[]danger.Violation{
			{Message: "some markdown"},
		},
		d.Results.Markdowns)
}
